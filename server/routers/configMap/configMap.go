package configMap

import (
	"server/controllers/configMap"

	"github.com/gin-gonic/gin"
)

func create(configMapGroup *gin.RouterGroup) {
	configMapGroup.POST("/create", configMap.Create)
}
func delete(configMapGroup *gin.RouterGroup) {
	configMapGroup.POST("/delete", configMap.Delete)
}
func deleteList(configMapGroup *gin.RouterGroup) {
	configMapGroup.POST("/deleteList", configMap.DeleteList)
}
func update(configMapGroup *gin.RouterGroup) {
	configMapGroup.POST("/update", configMap.Update)
}
func get(configMapGroup *gin.RouterGroup) {
	configMapGroup.GET("/get", configMap.Get)
}
func list(configMapGroup *gin.RouterGroup) {
	configMapGroup.GET("/list", configMap.List)
}

func RegisterSubRouter(g *gin.RouterGroup) {
	// 配置登录功能的路由策略
	configMapGroup := g.Group("configMap")
	create(configMapGroup)
	delete(configMapGroup)
	deleteList(configMapGroup)
	update(configMapGroup)
	get(configMapGroup)
	list(configMapGroup)
}
