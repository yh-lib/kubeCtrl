package configMap

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	controllers.KubectlFunc(c, "configMap", "update")
}
