package kubeutils

import (
	"context"

	"server/utils/logs"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

// 定义结构体
type ReplicaSet struct {
	InstanceInterface typedv1.AppsV1Interface
	Item              *appsv1.ReplicaSet
}

// New函数可以用于配置一些默认值
func NewReplicaSet(kubeconfig string, item *appsv1.ReplicaSet) *ReplicaSet {
	// 首先调用instance的init函数，生成一个ResourceInstance的实例，并配置默认值和生成clientset
	instance := ResourceInstance{}
	instance.Init(kubeconfig)
	// 定义一个ReplicaSet实例
	resource := ReplicaSet{}
	resource.InstanceInterface = instance.Clientset.AppsV1()
	resource.Item = item
	return &resource
}

// 创建资源
func (c *ReplicaSet) Create(namespace string) error {
	logs.Info(map[string]interface{}{"名称": c.Item.Name, "命名空间": namespace}, "创建ReplicaSet")
	c.Item.ResourceVersion = ""
	_, err := c.InstanceInterface.ReplicaSets(namespace).Create(context.TODO(), c.Item, v1.CreateOptions{})
	return err
}

// 删除资源
func (c *ReplicaSet) Delete(namespace, name string, gracePeriodSeconds *int64) error {
	logs.Warning(map[string]interface{}{"名称": name, "命名空间": namespace}, "删除ReplicaSet")
	deleteOption := v1.DeleteOptions{}
	// gracePeriodSeconds可配置，如果为0代表是强制删除
	if gracePeriodSeconds != nil {
		// 说明传递了gracePeriodSeconds
		deleteOption.GracePeriodSeconds = gracePeriodSeconds
	}
	err := c.InstanceInterface.ReplicaSets(namespace).Delete(context.TODO(), name, deleteOption)
	return err
}

// 删除多个
func (c *ReplicaSet) DeleteList(namespace string, nameList []string, gracePeriodSeconds *int64) error {
	// 删除多个时，结构体会接收一个nameList的切片，循环该切片，然后调用Delete函数即可
	for _, name := range nameList {
		// 调用删除函数
		c.Delete(namespace, name, gracePeriodSeconds)
	}
	// 忽略错误
	return nil
}

// 更新资源
func (c *ReplicaSet) Update(namespace string) error {
	logs.Warning(map[string]interface{}{"名称": c.Item.Name, "命名空间": namespace}, "更新ReplicaSet")
	c.Item.ObjectMeta.ResourceVersion = ""
	_, err := c.InstanceInterface.ReplicaSets(namespace).Update(context.TODO(), c.Item, v1.UpdateOptions{})
	return err
}

// 获取资源列表
func (c *ReplicaSet) List(namespace, labelSelector, fieldSelector string) (items interface{}, err error) {
	logs.Info(map[string]interface{}{"命名空间": namespace}, "查询ReplicaSet列表")
	// 有可能是根据查询条件进行查询
	listOptions := v1.ListOptions{
		FieldSelector: fieldSelector,
		LabelSelector: labelSelector,
	}
	list, err := c.InstanceInterface.ReplicaSets(namespace).List(context.TODO(), listOptions)
	items = list.Items
	return items, err
}

// 获取资源配置
func (c *ReplicaSet) Get(namespace, name string) (item interface{}, err error) {
	logs.Info(map[string]interface{}{"名称": name, "命名空间": namespace}, "查询ReplicaSet详情")
	i, err := c.InstanceInterface.ReplicaSets(namespace).Get(context.TODO(), name, v1.GetOptions{})
	c.Item = i
	i.APIVersion = "v1"
	i.Kind = "ReplicaSet"
	item = i
	return item, err
}
