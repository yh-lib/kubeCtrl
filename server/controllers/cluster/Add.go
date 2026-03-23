package cluster

import (
	"context"
	"server/config"
	"server/utils"
	"server/utils/logs"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var returnData = config.NewRetrunData()

func Add(c *gin.Context) {
	logs.Info(nil, "开始运行集群添加逻辑")
	// 接受前端提交的信息，并存放到一个结构体中
	clusterConfig := ClusterConfig{}
	if err := c.ShouldBindJSON(&clusterConfig); err != nil {
		msg := "添加集群的配置信息不完整:" + err.Error()
		returnData.Status = 400
		returnData.Message = msg
		c.JSON(200, returnData)
		return
	}
	logs.Info(map[string]any{"集群别名": clusterConfig.Alias, "集群ID": clusterConfig.Id}, "开始添加集群")
	// 判断集群是否正常
	ClusterStatus, err := clusterConfig.getClusterStatus()
	if err != nil {
		msg := "无法获取集群信息:" + err.Error()
		returnData.Status = 400
		returnData.Message = msg
		c.JSON(200, returnData)
		logs.Error(map[string]any{"error": err.Error()}, "添加集群失败")
		return
	}
	logs.Info(map[string]any{"集群状态": ClusterStatus}, "集群配置信息校验通过")
	// step2	secret 配置信息
	var clusterConfigSectet corev1.Secret
	clusterConfigSectet.Name = clusterConfig.Id
	clusterConfigSectet.Labels = make(map[string]string)
	clusterConfigSectet.Labels["kubeeasy.com/cluster.metadata"] = "true"
	// 添加注释，保存集群的配置信息
	clusterConfigSectet.Annotations = make(map[string]string)
	clusterConfigSectet.Annotations = utils.Struct2Map(ClusterStatus)
	// 保存kubeconfig
	clusterConfigSectet.StringData = make(map[string]string)
	clusterConfigSectet.StringData["kubeconfig"] = clusterConfig.Kubeconfig
	// step1	核心操作	创建 secret
	_, err = config.InClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).Create(context.TODO(), &clusterConfigSectet, metav1.CreateOptions{})
	if err != nil {
		logs.Error(map[string]any{"集群别名": clusterConfig.Alias, "集群ID": clusterConfig.Id, "ERROR": err.Error()}, "集群添加失败")
		msg := "集群添加失败:" + err.Error()
		returnData.Message = msg
		returnData.Status = 400
		c.JSON(200, returnData)
		return
	}
	logs.Info(map[string]any{"集群别名": clusterConfig.Alias, "集群ID": clusterConfig.Id}, "集群添加成功")
	returnData.Status = 200
	returnData.Message = "集群添加成功"
	c.JSON(200, returnData)
}
