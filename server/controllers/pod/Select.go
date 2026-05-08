package pod

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	controllers.KubectlFunc(c, "pod", "get")
}

func List(c *gin.Context) {
	controllers.KubectlFunc(c, "pod", "list")
}

func Status(c *gin.Context) {
	controllers.KubectlFunc(c, "pod", "status")
}
