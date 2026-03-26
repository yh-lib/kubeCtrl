package pod

import (
	"context"
	"server/config"
	"server/controllers"
	"server/utils/logs"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Get(c *gin.Context) {
	logs.Info(nil, "详情逻辑")
	returnData := config.ReturnData{}
	returnData.Data = map[string]any{}
	clientSet, basicInfo, err := controllers.BasicInit(c, nil)
	if err != nil {
		logs.Error(map[string]any{"Error": err}, "clientSet 初始化失败")
		returnData.Message = err.Error()
		returnData.Status = 400
		c.JSON(200, returnData)
		return
	}
	podDetail, err := clientSet.CoreV1().Pods(basicInfo.NameSpace).Get(context.TODO(), basicInfo.Name, metav1.GetOptions{})
	if err != nil {
		logs.Error(map[string]any{"Error": err}, "获取 pod: "+basicInfo.Name+" 详情失败")
		returnData.Status = 400
		returnData.Message = "获取 pod: " + basicInfo.Name + " 详情失败: " + err.Error()
	} else {
		logs.Error(map[string]any{"Error": err}, "获取 pod: "+basicInfo.Name+" 详情成功")
		returnData.Status = 200
		returnData.Message = "获取 pod: " + basicInfo.Name + " 详情成功"
		returnData.Data["item"] = podDetail
	}
	c.JSON(200, returnData)
}

func List(c *gin.Context) {
	logs.Info(nil, "列表逻辑")
	returnData := config.ReturnData{}
	returnData.Data = map[string]any{}
	clientSet, basicInfo, err := controllers.BasicInit(c, nil)
	if err != nil {
		logs.Error(map[string]any{"Error": err}, "clientSet 初始化失败")
		returnData.Message = err.Error()
		returnData.Status = 400
		c.JSON(200, returnData)
		return
	}
	podList, err := clientSet.CoreV1().Pods(basicInfo.NameSpace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logs.Error(map[string]any{"Error": err}, "查询 pod 列表失败")
		returnData.Status = 400
		returnData.Message = "查询 pod 列表失败: " + err.Error()
	} else {
		returnData.Status = 200
		returnData.Message = "查询 pod 列表成功"
		returnData.Data["items"] = podList.Items
	}
	c.JSON(200, returnData)
}
