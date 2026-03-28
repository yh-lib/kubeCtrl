package cluster

import (
	"context"
	"server/config"
	"server/controllers/node"
	"server/utils/logs"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func List(c *gin.Context) {
	// 定义存放返回前端数据的变量
	var returnData = config.NewRetrunData()
	clusterList, err := config.InClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logs.Error(map[string]any{"Error": err.Error()}, "获取集群列表失败")
		returnData.Status = 400
		returnData.Message = "获取集群列表失败"
		c.JSON(400, returnData)
	} else {
		logs.Info(nil, "获取集群列表成功")
		returnData.Status = 200
		returnData.Message = "获取集群列表成功"
		var clusterItem []map[string]string
		for _, v := range clusterList.Items {
			// 获取节点数量
			v.Annotations["clusterSize"] = node.NodesNum(c)
			clusterItem = append(clusterItem, v.Annotations)
		}
		returnData.Data["items"] = clusterItem
		c.JSON(200, returnData)
	}
}
