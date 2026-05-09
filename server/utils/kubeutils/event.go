package kubeutils

import (
	"context"

	"server/utils/logs"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	typedv1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

// 定义结构体
type Event struct {
	InstanceInterface typedv1.CoreV1Interface
	Item              *corev1.Event
}

// New函数可以用于配置一些默认值
func NewEvent(kubeconfig string, item *corev1.Event) *Event {
	// 首先调用instance的init函数，生成一个ResourceInstance的实例，并配置默认值和生成clientset
	instance := ResourceInstance{}
	instance.Init(kubeconfig)
	// 定义一个Deployment实例
	resource := Event{}
	resource.InstanceInterface = instance.Clientset.CoreV1()
	resource.Item = item
	// 将item转换为具体的类型
	return &resource
}

// 获取资源列表
func (c *Event) List(namespace, labelSelector, fieldSelector string) (items interface{}, err error) {
	logs.Info(map[string]interface{}{"命名空间": namespace}, "查询Deployment列表")
	// 有可能是根据查询条件进行查询
	listOptions := v1.ListOptions{
		FieldSelector: fieldSelector,
		LabelSelector: labelSelector,
	}
	list, err := c.InstanceInterface.Events(namespace).List(context.TODO(), listOptions)
	items = list.Items
	return items, err
}

// 创建资源
func (c *Event) Create(namespace string) error {
	logs.Info(map[string]any{"namespace": namespace}, "Event 暂不支持 Create 操作")
	return nil
}

// 删除资源
func (c *Event) Delete(namespace, name string, gracePeriodSeconds *int64) error {
	logs.Info(map[string]any{"namespace": namespace, "name": name, "gracePeriodSeconds": gracePeriodSeconds}, "Event 暂不支持 Delete 操作")
	return nil
}

// 删除多个
func (c *Event) DeleteList(namespace string, nameList []string, gracePeriodSeconds *int64) error {
	logs.Info(map[string]any{"namespace": namespace, "nameList": nameList, "gracePeriodSeconds": gracePeriodSeconds}, "Event 暂不支持 DeleteList 操作")
	return nil
}

// 更新资源
func (c *Event) Update(namespace string) error {
	logs.Info(map[string]any{"namespace": namespace}, "Event 暂不支持 Update 操作")
	return nil
}

// 获取资源配置
func (c *Event) Get(namespace, name string) (interface{}, error) {
	logs.Info(map[string]any{"namespace": namespace, "name": name}, "Event 暂不支持 Get 操作")
	return nil, nil
}
