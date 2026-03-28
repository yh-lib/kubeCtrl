// 存放一些共用的代码
package node

import (
	"server/config"
	"server/controllers"
	"strconv"

	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
)

func NodesNum(c *gin.Context) string {
	// 获取节点列表
	// 定义变量
	var (
		kubeUtilser kubeutils.KubeUtilser
		returndata  config.ReturnData
		info        controllers.Info
		itemsNum    string
	)
	kubeconfig := controllers.NewInfo(c, &info)
	// 通过构造函数 构建一个实例
	instance := kubeutils.NewNode(kubeconfig, nil)
	kubeUtilser = instance
	items, _ := kubeUtilser.List("", info.LabelSelector, info.FieldSelector)
	if value, ok := items.([]any); ok {
		itemsNum = strconv.Itoa(len(value))
	}
	returndata.Status = 200
	returndata.Message = "节点数据量: " + itemsNum
	c.JSON(200, returndata)
	return itemsNum
}
