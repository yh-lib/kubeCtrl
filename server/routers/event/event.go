package event

import (
	"server/controllers/event"

	"github.com/gin-gonic/gin"
)

func list(eventGroup *gin.RouterGroup) {
	eventGroup.GET("/list", event.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	// 配置登录功能的路由策略
	eventGroup := g.Group("event")
	list(eventGroup)
}
