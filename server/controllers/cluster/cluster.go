// 存放一些共用的代码
package cluster

import (
	"context"
	"server/config"
	"server/utils"
	"server/utils/logs"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// 定义一个结构体，用于描述创建集群所用的配置信息
type ClusterInfo struct {
	Id       string `json:"clusterId" binding:"required"`
	Alias    string `json:"clusterAlias" binding:"required"`
	Location string `json:"clusterLocation" binding:"required"`
}

type ClusterConfig struct {
	ClusterInfo
	Kubeconfig string `json:"kubeconfig" binding:"required"`
}

// 定义一个结构体，用于描述集群的状态
type ClusterStatus struct {
	ClusterInfo
	Version string `json:"clusterVersion"`
	Status  string `json:"clusterStatus"`
}

// 定义一个结构体的方法，用于判断集群的状态
func (c *ClusterConfig) getClusterStatus() (ClusterStatus, error) {
	// 判断集群是否是正常
	clusterStatus := ClusterStatus{}
	clusterStatus.ClusterInfo = c.ClusterInfo
	// 基于前端传递的kubeconfig构建新的clientSet
	restConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(c.Kubeconfig))
	if err != nil {
		return clusterStatus, err
	}
	clientSet, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return clusterStatus, err
	}
	// 获取集群版本
	serverVersion, err := clientSet.Discovery().ServerVersion()
	if err != nil {
		return clusterStatus, err
	}
	clusterVersion := serverVersion.String()
	clusterStatus.Version = clusterVersion
	clusterStatus.Status = "Active"

	return clusterStatus, nil
}

// 添加或更新集群信息
func createOrUpdate(c *gin.Context, method string) {
	// 定义存放返回前端数据的变量
	var returnData = config.NewRetrunData()
	var arg string
	if method == "create" {
		arg = "添加"
	} else {
		arg = "更新"
	}
	// 接受前端提交的信息，并存放到一个结构体中
	clusterConfig := ClusterConfig{}
	if err := c.ShouldBindJSON(&clusterConfig); err != nil {
		msg := arg + "集群的配置信息不完整:" + err.Error()
		returnData.Status = 400
		returnData.Message = msg
		c.JSON(200, returnData)
		return
	}
	logs.Info(map[string]any{"集群别名": clusterConfig.Alias, "集群ID": clusterConfig.Id}, "开始校验集群配置信息")
	// 判断集群是否正常
	ClusterStatus, err := clusterConfig.getClusterStatus()
	if err != nil {
		msg := "无法获取集群信息:" + err.Error()
		returnData.Status = 400
		returnData.Message = msg
		c.JSON(200, returnData)
		logs.Error(map[string]any{"error": err.Error()}, arg+"集群失败")
		return
	}
	logs.Info(map[string]any{"集群状态": ClusterStatus}, "集群配置信息校验通过")
	// step2	secret 配置信息
	var clusterConfigSectet corev1.Secret
	clusterConfigSectet.Name = clusterConfig.Id
	clusterConfigSectet.Labels = make(map[string]string)
	clusterConfigSectet.Labels["author"] = "yaohui.li"
	clusterConfigSectet.Labels["type"] = "clusterDetail"
	// 添加注释，保存集群的配置信息
	clusterConfigSectet.Annotations = make(map[string]string)
	clusterConfigSectet.Annotations = utils.Struct2Map(ClusterStatus)
	// 保存kubeconfig
	clusterConfigSectet.StringData = make(map[string]string)
	clusterConfigSectet.StringData["kubeconfig"] = clusterConfig.Kubeconfig
	// step1	核心操作	创建 secret
	var secret *corev1.Secret
	if method == "create" {
		secret, err = config.InClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).Create(context.TODO(), &clusterConfigSectet, metav1.CreateOptions{})
		if err != nil {
			logs.Error(map[string]any{"集群别名": clusterConfig.Alias, "集群ID": clusterConfig.Id, "ERROR": err.Error()}, "集群添加失败")
			msg := "集群添加失败:" + err.Error()
			returnData.Message = msg
			returnData.Status = 400
			c.JSON(200, returnData)
			return
		}
	} else {
		secret, err = config.InClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).Update(context.TODO(), &clusterConfigSectet, metav1.UpdateOptions{})
		if err != nil {
			logs.Error(map[string]any{"集群别名": clusterConfig.Alias, "集群ID": clusterConfig.Id, "ERROR": err.Error()}, "集群更新失败")
			msg := "集群更新失败:" + err.Error()
			returnData.Message = msg
			returnData.Status = 400
			c.JSON(200, returnData)
			return
		}
	}
	// 更新存放kubeconfig的变量
	clusterId := secret.Name
	kubeConfigStr := string(secret.Data["kubeconfig"])
	config.ClusterKubeconfig[clusterId] = kubeConfigStr

	logs.Info(map[string]any{"集群别名": clusterConfig.Alias, "集群ID": clusterConfig.Id}, "集群"+arg+"成功")
	returnData.Status = 200
	returnData.Message = "集群" + arg + "成功"
	c.JSON(200, returnData)
}
