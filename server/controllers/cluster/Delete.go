package cluster

import (
	"context"
	"server/config"
	"server/utils/logs"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Delete(c *gin.Context) {
	clusterID := c.Query("clusterid")
	logs.Info(nil, "开始删除集群: "+clusterID)
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
