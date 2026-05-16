package cluster

import (
	"context"
	"server/config"
	"server/utils/logs"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Info struct {
	GetStatus bool `json:"getStatus" form:"getStatus"`
}

// 更新集群状态
func updateClusterStatus(clusterInfo corev1.Secret) {
	// clientSet 实例化
	restConfig, err := clientcmd.RESTConfigFromKubeConfig(clusterInfo.Data["kubeconfig"])
	if err != nil {
		logs.Error(map[string]any{"Error": err.Error()}, "kubeconfig 初始化失败")
		clusterInfo.Annotations["clusterStatus"] = "Inactive"
		return
	}
	restConfig.Timeout = 1 * time.Second
	clientSet, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		logs.Error(map[string]any{"Error": err.Error()}, "clientSet 初始化失败")
		clusterInfo.Annotations["clusterStatus"] = "Inactive"
		return
	}
	// 获取集群版本及状态
	serverVersion, err := clientSet.Discovery().ServerVersion()
	if err != nil {
		logs.Error(map[string]any{"Error": err.Error()}, "集群版本获取失败")
		clusterInfo.Annotations["clusterStatus"] = "Inactive"
		return
	}
	clusterVersion := serverVersion.String()
	clusterInfo.Annotations["clusterVersion"] = clusterVersion
	clusterInfo.Annotations["clusterStatus"] = "Active"
	// Node 数量
	nodes, err := clientSet.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logs.Error(map[string]any{"Error": err.Error()}, "Node 数量获取失败")
		return
	}
	clusterInfo.Annotations["clusterNodeNum"] = strconv.Itoa(len(nodes.Items))
	// Pod 数量
	pods, err := clientSet.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logs.Error(map[string]any{"Error": err.Error()}, "Pod 数量获取失败")
		return
	}
	clusterInfo.Annotations["clusterPodNum"] = strconv.Itoa(len(pods.Items))
	// Ns 数量
	ns, err := clientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logs.Error(map[string]any{"Error": err.Error()}, "Ns 数量获取失败")
		return
	}
	clusterInfo.Annotations["clusterNsNum"] = strconv.Itoa(len(ns.Items))
}

func List(c *gin.Context) {
	var (
		err  error
		info Info
	)
	err = c.ShouldBindQuery(&info)
	logs.Info(map[string]any{"Info": info}, "/api/cluster/list 前后端数据绑定结果")
	// 定义存放返回前端数据的变量
	returnData := config.NewRetrunData()
	// 更新集群状态
	clusterList, err := config.InClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).List(context.TODO(), metav1.ListOptions{LabelSelector: "type=clusterDetail"})
	if err != nil {
		logs.Error(map[string]any{"Error": err.Error()}, "获取集群列表失败")
		returnData.Status = 400
		returnData.Message = "获取集群列表失败"
		c.JSON(400, returnData)
	} else {
		logs.Debug(nil, "获取集群列表成功")
		returnData.Status = 200
		returnData.Message = "获取集群列表成功"
		var clusterItems []map[string]string
		for _, v := range clusterList.Items {
			// 更新集群状态
			if info.GetStatus == true {
				updateClusterStatus(v)
			}
			clusterItems = append(clusterItems, v.Annotations)
		}
		returnData.Data["items"] = clusterItems
		c.JSON(200, returnData)
	}
}
