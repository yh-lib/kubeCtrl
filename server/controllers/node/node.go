// 存放一些共用的代码
package node

import (
	"server/config"
	"server/controllers"
	"strconv"

	"server/utils/kubeutils"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
)

// 获取节点列表
func NodesNum(c *gin.Context) string {
	var (
		returndata config.ReturnData
		info       controllers.Info
		itemsNum   string
	)
	kubeconfig := controllers.NewInfo(c, &info)
	// 通过构造函数 构建一个实例
	instance := kubeutils.NewNode(kubeconfig, nil)
	items, _ := instance.List("", info.LabelSelector, info.FieldSelector)
	if itemValue, ok := items.([]corev1.Node); ok {
		itemsNum = strconv.Itoa(len(itemValue))
	}
	returndata.Status = 200
	returndata.Message = "节点数据量: " + itemsNum
	c.JSON(200, returndata)
	return itemsNum
}
