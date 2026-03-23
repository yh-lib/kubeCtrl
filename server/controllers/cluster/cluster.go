// 存放一些共用的代码
package cluster

import (
	"context"
	"server/config"
	"server/utils"
	"server/utils/logs"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// 定义一个结构体，用于描述创建集群所用的配置信息
type ClusterInfo struct {
	Id       string `json:"id" binding:"required"`
	Alias    string `json:"alias" binding:"required"`
	Location string `json:"location" binding:"required"`
}

type ClusterConfig struct {
	ClusterInfo
	Kubeconfig string `json:"kubeconfig" binding:"required"`
}

// 定义一个结构体，用于描述集群的状态
type ClusterStatus struct {
	ClusterInfo
	Version string `json:"version"`
	Status  string `json:"status"`
}

// 定义一个结构体的方法，用于判断集群的状态
func (c *ClusterConfig) getClusterStatus() (ClusterStatus, error) {
	// 判断集群是否是正常
	ClusterStatus := ClusterStatus{}
	ClusterStatus.ClusterInfo = c.ClusterInfo
	serverVersion, err := config.InClusterClientSet.Discovery().ServerVersion()
	if err != nil {
		return ClusterStatus, err
	}
	clusterVersion := serverVersion.String()
	ClusterStatus.Version = clusterVersion
	ClusterStatus.Status = "Active"
	return ClusterStatus, nil
}

// 添加或更新集群信息
func addOrUpdate(c *gin.Context, method string) {
	// 定义存放返回前端数据的变量
	var returnData = config.NewRetrunData()
	var arg string
	if method == "create" {
		arg = "添加"
	} else {
		arg = "更新"
	}
	// 接受前端提交的信息，并存放到一个结构体中
	clusterConfig := ClusterConfig{}
	if err := c.ShouldBindJSON(&clusterConfig); err != nil {
		msg := arg + "集群的配置信息不完整:" + err.Error()
		returnData.Status = 400
		returnData.Message = msg
		c.JSON(200, returnData)
		return
	}
	logs.Info(map[string]any{"集群别名": clusterConfig.Alias, "集群ID": clusterConfig.Id}, "开始"+arg+"集群")
	// 判断集群是否正常
	ClusterStatus, err := clusterConfig.getClusterStatus()
	if err != nil {
		msg := "无法获取集群信息:" + err.Error()
		returnData.Status = 400
		returnData.Message = msg
		c.JSON(200, returnData)
		logs.Error(map[string]any{"error": err.Error()}, arg+"集群失败")
		return
	}
	logs.Info(map[string]any{"集群状态": ClusterStatus}, "集群配置信息校验通过")
	// step2	secret 配置信息
	var clusterConfigSectet corev1.Secret
	clusterConfigSectet.Name = clusterConfig.Id
	clusterConfigSectet.Labels = make(map[string]string)
	clusterConfigSectet.Labels["kubeeasy.com/cluster.metadata"] = "true"
	// 添加注释，保存集群的配置信息
	clusterConfigSectet.Annotations = make(map[string]string)
	clusterConfigSectet.Annotations = utils.Struct2Map(ClusterStatus)
	// 保存kubeconfig
	clusterConfigSectet.StringData = make(map[string]string)
	clusterConfigSectet.StringData["kubeconfig"] = clusterConfig.Kubeconfig
	// step1	核心操作	创建 secret
	if method == "create" {
		_, err = config.InClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).Create(context.TODO(), &clusterConfigSectet, metav1.CreateOptions{})
		if err != nil {
			logs.Error(map[string]any{"集群别名": clusterConfig.Alias, "集群ID": clusterConfig.Id, "ERROR": err.Error()}, "集群添加失败")
			msg := "集群添加失败:" + err.Error()
			returnData.Message = msg
			returnData.Status = 400
			c.JSON(200, returnData)
			return
		}
	} else {
		_, err = config.InClusterClientSet.CoreV1().Secrets(config.MetadataNamespace).Update(context.TODO(), &clusterConfigSectet, metav1.UpdateOptions{})
		if err != nil {
			logs.Error(map[string]any{"集群别名": clusterConfig.Alias, "集群ID": clusterConfig.Id, "ERROR": err.Error()}, "集群更新失败")
			msg := "集群更新失败:" + err.Error()
			returnData.Message = msg
			returnData.Status = 400
			c.JSON(200, returnData)
			return
		}
	}
	logs.Info(map[string]any{"集群别名": clusterConfig.Alias, "集群ID": clusterConfig.Id}, "集群"+arg+"成功")
	returnData.Status = 200
	returnData.Message = "集群" + arg + "成功"
	c.JSON(200, returnData)
}
