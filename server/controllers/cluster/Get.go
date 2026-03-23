package cluster

import (
	"context"
	"server/config"
	"server/utils/logs"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Get(c *gin.Context) {
	// 定义存放返回前端数据的变量
	var returnData = config.NewRetrunData()
	clusterId := c.Query("clusterid")
	clusterDetail, err := config.InClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).Get(context.TODO(), clusterId, metav1.GetOptions{})
	if err != nil {
		logs.Error(map[string]any{"clusterId": clusterId, "Error": err.Error()}, "获取集群详情失败")
		returnData.Status = 400
		returnData.Message = "获取集群详情失败: " + err.Error()
		c.JSON(400, returnData)
	} else {
		logs.Info(map[string]any{"clusterId": clusterId}, "获取集群详情成功")
		returnData.Status = 200
		returnData.Message = "获取集群详情成功"
		clusterConfigMap := clusterDetail.Annotations
		clusterConfigMap["kubeconfig"] = string(clusterDetail.Data["kubeconfig"])
		returnData.Data["item"] = clusterConfigMap
		c.JSON(200, returnData)
	}
}
