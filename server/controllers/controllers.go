// 控制器层  实现路由的处理逻辑
package controllers

import (
	"errors"
	"server/config"
	"server/utils/logs"

	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type BasicInfo struct {
	ClusterId  string   `json:"clusterId" form:"clusterId"`
	NameSpace  string   `json:"nameSpace" form:"nameSpace"`
	Name       string   `json:"name" form:"name"`
	Item       any      `json:"item"`
	DeleteList []string `json:"deleteList"`
}

type Info struct {
	BasicInfo
	ReturnData    config.ReturnData
	LabelSelector string `json:"labelSelector" form:"labelSelector"`
	FieldSelector string `json:"fieldSelector" form:"fieldSelector"`
	// 判断是否是强制删除
	Force bool `json:"force" form:"force"`
}

func NewInfo(c *gin.Context, info *Info) (kubeconfig string) {
	var (
		err        error
		returnData config.ReturnData
		// 赋值变量部分
		requestMethod = c.Request.Method
	)
	// 前后端数据绑定
	switch requestMethod {
	case "GET":
		err = c.ShouldBindQuery(&info)
	case "POST":
		err = c.ShouldBindJSON(&info)
	default:
		err = errors.New("不支持的请求类型")
	}
	logs.Debug(map[string]any{"Info": info}, "数据绑定结果")
	if err != nil {
		msg := "请求数据绑定后端失败: " + err.Error()
		returnData.Message = msg
		returnData.Status = 400
		c.JSON(200, returnData)
		logs.Error(nil, msg)
		return
	}
	// 获取kubeconfig
	return config.ClusterKubeconfig[info.ClusterId]
}

func BasicInit(c *gin.Context, item any) (clientSet *kubernetes.Clientset, basicInfo BasicInfo, err error) {
	// 绑定前端传递的数据到后端变量中
	basicInfo = BasicInfo{}
	basicInfo.Item = item
	requestMethod := c.Request.Method
	switch requestMethod {
	case "GET":
		err = c.ShouldBindQuery(&basicInfo)
	case "POST":
		err = c.ShouldBindJSON(&basicInfo)
	default:
		err = errors.New("不支持的请求类型")
	}
	if err != nil {
		msg := "请求数据绑定后端失败: " + err.Error()
		return nil, basicInfo, errors.New(msg)
	}
	if basicInfo.NameSpace == "" {
		basicInfo.NameSpace = "default"
	}
	// 实例化clientSet
	kubeConfig := config.ClusterKubeconfig[basicInfo.ClusterId]
	kubeConfigByte := []byte(kubeConfig)
	// kubeConfig
	restConfig, err := clientcmd.RESTConfigFromKubeConfig(kubeConfigByte)
	if err != nil {
		msg := "实例化 kubeconfig 失败:" + err.Error()
		return nil, basicInfo, errors.New(msg)
	}
	// clientSet
	clientSet, err = kubernetes.NewForConfig(restConfig)
	if err != nil {
		msg := "实例化clientSet失败:" + err.Error()
		return nil, basicInfo, errors.New(msg)
	}
	// 校验 kubeconfig
	serverVersion, err := clientSet.Discovery().ServerVersion()
	if err != nil {
		msg := "kubeconfig 校验失败: " + err.Error()
		return nil, basicInfo, errors.New(msg)
	}
	logs.Debug(nil, "集群版本："+serverVersion.String())
	return clientSet, basicInfo, nil
}
