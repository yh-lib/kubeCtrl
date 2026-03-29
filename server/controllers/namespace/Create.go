package namespace

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	controllers.KubectlFunc(c, "namespace", "create")
}
