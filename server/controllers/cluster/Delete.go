package cluster

import (
	"context"
	"server/config"
	"server/utils/logs"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type RequestBody struct {
	ClusterId string `json:"clusterId"`
}

func Delete(c *gin.Context) {
	// 定义存放返回前端数据的变量
	var returnData = config.NewRetrunData()

	// 接受前端提交的信息，并存放到一个结构体中
	clusterConfig := RequestBody{}
	if err := c.ShouldBindJSON(&clusterConfig); err != nil {
		msg := err.Error()
		returnData.Status = 400
		returnData.Message = msg
		c.JSON(200, returnData)
		return
	}
	clusterId := clusterConfig.ClusterId
	logs.Info(nil, "开始运行集群 "+clusterId+" 删除逻辑")
	err := config.InClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).Delete(context.TODO(), clusterId, metav1.DeleteOptions{})
	if err != nil {
		logs.Error(nil, "删除集群 "+clusterId+" 失败")
		returnData.Status = 400
		returnData.Message = "删除集群 " + clusterId + " 失败: " + err.Error()
		c.JSON(400, returnData)
	} else {
		logs.Info(nil, "删除集群 "+clusterId+" 成功")
		// 更新存放kubeconfig的变量
		delete(config.ClusterKubeconfig, clusterId)
		returnData.Status = 200
		returnData.Message = "删除集群 " + clusterId + " 成功"
		c.JSON(200, returnData)
	}
}
