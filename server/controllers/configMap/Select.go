package configMap

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	controllers.KubectlFunc(c, "configMap", "get")
}

func List(c *gin.Context) {
	controllers.KubectlFunc(c, "configMap", "list")
}
