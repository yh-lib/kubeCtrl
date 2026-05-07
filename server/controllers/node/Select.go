package node

import (
	"server/controllers"
	"server/utils/logs"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	controllers.KubectlFunc(c, "node", "get")
}

func List(c *gin.Context) {
	controllers.KubectlFunc(c, "node", "list")
}

func Metrics(c *gin.Context) {
	logs.Info(nil, "metrics路由实现")
	// controllers.KubectlFunc(c, "node", "list")
}
