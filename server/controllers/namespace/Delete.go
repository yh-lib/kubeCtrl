package namespace

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	controllers.KubectlFunc(c, "namespace", "delete")
}

func DeleteList(c *gin.Context) {
	controllers.KubectlFunc(c, "namespace", "deleteList")
}
