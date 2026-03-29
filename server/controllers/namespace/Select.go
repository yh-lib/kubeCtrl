package namespace

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	controllers.KubectlFunc(c, "namespace", "get")
}

func List(c *gin.Context) {
	controllers.KubectlFunc(c, "namespace", "list")
}
