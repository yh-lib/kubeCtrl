package kubeutils

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"server/utils/logs"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// 定义结构体
type Tools struct {
	ClusterId     string
	DynamicClient *dynamic.DynamicClient
}

func NewClientSet(kubeconfig string, timeout int) (clientset *kubernetes.Clientset, err error) {
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(kubeconfig))
	if err != nil {
		msg := "解析kubeconfig错误:" + err.Error()
		return nil, errors.New(msg)
	}
	// 设置超时时间
	config.Timeout = time.Duration(timeout) * time.Second
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		msg := "创建clientset失败:" + err.Error()
		return nil, errors.New(msg)
	}
	return clientset, nil
}

// New函数可以用于配置一些默认值
func NewTools(kubeconfig string) (tools *Tools, err error) {
	// 加载kubeconfig文件，并创建restConfig
	config, _ := clientcmd.RESTConfigFromKubeConfig([]byte(kubeconfig))
	// 创建一个discovery客户端
	dynamicClient, err := dynamic.NewForConfig(config)
	tools = &Tools{}
	tools.DynamicClient = dynamicClient
	return tools, err
}

func createOrUpdate(dynamicClient *dynamic.DynamicClient, yamlContent, method string) (string, error) {
	// 拆分yaml
	var errMsgList []string
	errCount := 0
	methodMsg := "创建"
	yamlList := strings.Split(yamlContent, "---")
	// 循环yaml列表
	for k, v := range yamlList {
		// 去除空行和\n
		if v == "" || v == "\n" {
			continue
		}
		index := k + 1
		logs.Debug(map[string]interface{}{"yaml": v, "index": index}, "基于yaml创建或更新")
		// 创建一个非结构化数据对象，用于接受解析的yaml文件内容
		obj := &unstructured.Unstructured{}
		// GVK：Group Version Kind
		_, gvk, err := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme).Decode([]byte(v), nil, obj)
		if err != nil {
			errCount += 1
			// 如果此处失败，说明这一列的yaml内容有问题，直接下一个即可
			msg := fmt.Sprintf("第%d项yaml数据序列化: %s \n", index, err.Error())
			errMsgList = append(errMsgList, msg)
			continue
		}
		// yaml序列化成功继续往下执行
		// 获取非结构化数据的namespace
		namespace := obj.GetNamespace()
		if namespace == "" {
			namespace = "default"
		}
		// 创建一个gvr：Group Version Resource  Resource=Kind + s example deployments
		resource := gvk.Kind + "s"
		gvr := schema.GroupVersionResource{
			Group:    gvk.Group,
			Version:  gvk.Version,
			Resource: strings.ToLower(resource),
		}
		// 创建dynamic资源接口
		dynamicResourceInterface := dynamicClient.Resource(gvr).Namespace(namespace)
		switch method {
		case "Create", "Apply":
			_, err = dynamicResourceInterface.Create(context.TODO(), obj, metav1.CreateOptions{})
			if err != nil && method == "Apply" {
				// 尝试更新
				_, err = dynamicResourceInterface.Update(context.TODO(), obj, metav1.UpdateOptions{})
			}
		case "Update":
			methodMsg = "更新"
			_, err = dynamicResourceInterface.Update(context.TODO(), obj, metav1.UpdateOptions{})
		// case "Apply":
		// 	methodMsg = "应用"
		// 	name := obj.GetName()
		// 	_, err = dynamicResourceInterface.Apply(context.TODO(), name, obj, metav1.ApplyOptions{})
		case "Delete":
			methodMsg = "删除"
			name := obj.GetName()
			err = dynamicResourceInterface.Delete(context.TODO(), name, metav1.DeleteOptions{})
		}
		if err != nil {
			errCount += 1
			msg := fmt.Sprintf("第%d项yaml数据%s失败: %s", index, methodMsg, err.Error())
			errMsgList = append(errMsgList, msg)
		} else {
			msg := fmt.Sprintf("第%d项yaml数据%s成功", index, methodMsg)
			errMsgList = append(errMsgList, msg)
		}
	}
	if errCount == 0 {
		return "", nil
	} else {
		errMsg := fmt.Sprintf("%s失败", methodMsg)
		return strings.Join(errMsgList, "\n"), errors.New(errMsg)
	}

}

// 创建资源
func (c *Tools) Create(yamlContent string) (msg string, err error) {
	msg, err = createOrUpdate(c.DynamicClient, yamlContent, "Create")
	return
}

func (c *Tools) Update(yamlContent string) (msg string, err error) {
	msg, err = createOrUpdate(c.DynamicClient, yamlContent, "Update")
	return
}

func (c *Tools) Apply(yamlContent string) (msg string, err error) {
	msg, err = createOrUpdate(c.DynamicClient, yamlContent, "Apply")
	return
}

func (c *Tools) Delete(yamlContent string) (msg string, err error) {
	msg, err = createOrUpdate(c.DynamicClient, yamlContent, "Delete")
	return
}
