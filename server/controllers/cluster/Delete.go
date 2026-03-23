package cluster

import (
	"context"
	"server/config"
	"server/utils/logs"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Delete(c *gin.Context) {
	// 定义存放返回前端数据的变量
	var returnData = config.NewRetrunData()
	clusterID := c.Query("clusterid")
	logs.Info(nil, "开始运行集群 "+clusterID+" 删除逻辑")
	err := config.InClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).Delete(context.TODO(), clusterID, metav1.DeleteOptions{})
	if err != nil {
		logs.Error(nil, "删除集群 "+clusterID+" 失败")
		returnData.Status = 400
		returnData.Message = "删除集群 " + clusterID + " 失败: " + err.Error()
		c.JSON(400, returnData)
	} else {
		logs.Info(nil, "删除集群 "+clusterID+" 成功")
		returnData.Status = 200
		returnData.Message = "删除集群 " + clusterID + " 成功"
		c.JSON(200, returnData)
	}
}
