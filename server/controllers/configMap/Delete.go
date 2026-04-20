package configMap

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	controllers.KubectlFunc(c, "configMap", "delete")
}

func DeleteList(c *gin.Context) {
	controllers.KubectlFunc(c, "configMap", "deleteList")
}
