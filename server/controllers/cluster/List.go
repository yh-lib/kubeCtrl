package cluster

import (
	"server/config"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	// 定义存放返回前端数据的变量
	var returnData = config.NewRetrunData()
	returnData.Message = "获取list成功"
	c.JSON(200, returnData)
}
