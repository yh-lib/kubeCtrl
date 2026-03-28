package node

import (
	"fmt"
	"server/config"
	"server/controllers"
	"server/utils/logs"

	"github.com/dotbalo/kubeutils/kubeutils"
	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
)

func Update(c *gin.Context) {
	logs.Info(nil, "更新逻辑")
	// 定义变量
	var (
		kubeUtilser kubeutils.KubeUtilser
		returndata  config.ReturnData
		info        controllers.Info
		item        corev1.Node
	)
	info.Item = &item
	kubeconfig := controllers.NewInfo(c, &info)
	// 通过构造函数 构建一个实例
	instance := kubeutils.NewNode(kubeconfig, &item)
	kubeUtilser = instance
	err := kubeUtilser.Update(info.NameSpace)
	if err != nil {
		fmt.Println(err)
	}
	returndata.Status = 200
	returndata.Message = "更新成功"
	c.JSON(200, returndata)
}
