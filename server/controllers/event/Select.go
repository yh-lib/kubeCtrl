package event

import (
	"server/controllers"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	controllers.KubectlFunc(c, "event", "list")
}
