// 项目的总入口
package main

import (
	"context"
	"encoding/json"
	"fmt"
	_ "server/config"
	"server/utils/logs"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// 操作参数
var (
	option         string = "get"       // option: add, delete, update, get, list
	resource       string = "namespace" // resource: namespace, deployment
	resourceName   string = "scaffold"  // resourceName: the name of the namespace or deployment to be added or deleted
	resourceNsName string = "scaffold"  // resourceNsName: the name of the namespace where the deployment is located, only used when resource is deployment

	newNamespace corev1.Namespace  // 定义一个新的nameSspace对象
	newDeopyment appsv1.Deployment // 定义一个新的deployment对象
)

// kubectl api-resources

// 增加一个namespace
func addNamespace(clientset *kubernetes.Clientset, resourceName string) {
	newNamespace.Name = resourceName
	_, err := clientset.CoreV1().Namespaces().Create(context.TODO(), &newNamespace, metav1.CreateOptions{})
	if err != nil {
		logs.Error(map[string]any{"Error:": err.Error()}, "namespace创建失败.")
	}
}

// 增加一个deployment
func addDeployment(clientset *kubernetes.Clientset, resourceName string) {
	newDeopyment.Name = resourceName
	newDeopyment.Spec.Replicas = new(int32)
	*newDeopyment.Spec.Replicas = 1
	newDeopyment.Spec.Template.Spec.Containers = append(newDeopyment.Spec.Template.Spec.Containers, corev1.Container{
		Name:  "nginx",
		Image: "nginx:latest",
	})

	label := make(map[string]string)
	label["app"] = "nginx"
	label["version"] = "v1"
	newDeopyment.Spec.Selector = &metav1.LabelSelector{}
	newDeopyment.Spec.Selector.MatchLabels = label
	newDeopyment.Spec.Template.ObjectMeta.Labels = label

	_, err := clientset.AppsV1().Deployments(resourceNsName).Create(context.TODO(), &newDeopyment, metav1.CreateOptions{})
	if err != nil {
		logs.Error(map[string]any{"Error:": err.Error()}, "deployment创建失败.")
	}
}

func addDeploymentForJson(clientset *kubernetes.Clientset, resourceName string) {
	deployJson := `{
		"kind": "Deployment",
		"apiVersion": "apps/v1",
		"metadata": {
			"name": "redis",
			"creationTimestamp": null,
			"labels": {
				"app": "redis"
			}
		},
		"spec": {
			"replicas": 1,
			"selector": {
				"matchLabels": {
					"app": "redis"
				}
			},
			"template": {
				"metadata": {
					"creationTimestamp": null,
					"labels": {
						"app": "redis"
					}
				},
				"spec": {
					"containers": [
						{
							"name": "redis",
							"image": "redis",
							"resources": {}
						}
					]
				}
			},
			"strategy": {}
		},
		"status": {}
	}`
	err := json.Unmarshal([]byte(deployJson), &newDeopyment)
	if err != nil {
		logs.Error(map[string]any{"Error:": err.Error()}, "deployment json解析失败.")
	}
	_, err = clientset.AppsV1().Deployments(resourceNsName).Create(context.TODO(), &newDeopyment, metav1.CreateOptions{})
	if err != nil {
		logs.Error(map[string]any{"Error:": err.Error()}, "deployment创建失败.")
	}
}

// 删除一个namespace
func deleteNamespace(clientset *kubernetes.Clientset, resourceName string) {
	err := clientset.CoreV1().Namespaces().Delete(context.TODO(), resourceName, metav1.DeleteOptions{})
	if err != nil {
		logs.Error(map[string]any{"Error:": err.Error()}, "namespace删除失败.")
	}
}

// 删除一个deployment
func deleteDeployment(clientset *kubernetes.Clientset, resourceName string) {
	err := clientset.AppsV1().Deployments(resourceNsName).Delete(context.TODO(), resourceName, metav1.DeleteOptions{})
	if err != nil {
		logs.Error(map[string]any{"Error": err.Error()}, "deployment删除失败.")
	}
}

// 更新一个deployment
func updateDeployment(clientset *kubernetes.Clientset, resourceName string) {
	// 获取当前deployment对象
	deploy, err := clientset.AppsV1().Deployments(resourceNsName).Get(context.TODO(), resourceName, metav1.GetOptions{})
	if err != nil {
		logs.Error(map[string]any{"Error:": err.Error()}, "deployment detail获取失败.")
		return
	}
	// 更新deployment对象
	deploy.Spec.Replicas = new(int32)
	*deploy.Spec.Replicas = 1
	_, err = clientset.AppsV1().Deployments(resourceNsName).Update(context.TODO(), deploy, metav1.UpdateOptions{})
	if err != nil {
		logs.Error(map[string]any{"Error:": err.Error()}, "deployment更新失败.")
	}
}

// +++++查
// 获取当前deployments列表
func listDeployments(clientset *kubernetes.Clientset, resourceNsName string) {
	deploys, err := clientset.AppsV1().Deployments(resourceNsName).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logs.Error(map[string]any{"Error:": err.Error()}, "deployment list获取失败.")
		return
	}
	for _, deploy := range deploys.Items {
		logs.Info(nil, deploy.Name)
	}
}

// 获取namespaces列表
func listNamespaces(clientset *kubernetes.Clientset) {
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logs.Error(map[string]any{"Error:": err.Error()}, "namespace list获取失败.")
		return
	}
	for _, ns := range namespaces.Items {
		logs.Info(nil, ns.Name)
	}
}

// 获取pods列表
func listPods(clientset *kubernetes.Clientset, resourceNsName string) {
	pods, err := clientset.CoreV1().Pods(resourceNsName).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logs.Error(map[string]any{"Error:": err.Error()}, "pods list获取失败.")
		return
	}
	for _, pod := range pods.Items {
		logs.Info(nil, pod.Name)
	}
}

// 查看ns详情
func getNsDetail(clientset *kubernetes.Clientset, resourceName string) {
	nsDetail, err := clientset.CoreV1().Namespaces().Get(context.TODO(), resourceName, metav1.GetOptions{})
	if err != nil {
		logs.Error(map[string]any{"Error:": err.Error()}, "namespace detail获取失败.")
		return
	}
	fmt.Println("查询到resourceName的详情:", nsDetail)
}

func selectActions() {
	// step1 初始化 config 实例
	config, err := clientcmd.BuildConfigFromFlags("", "./config/baseKubeConfig")
	if err != nil {
		panic(err.Error())
	}

	// step2 创建客户端工具 clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// step3 通过参数控制执行方法
	switch option {
	case "add":
		switch resource {
		case "namespace":
			addNamespace(clientset, resourceName)
		case "deployment":
			// addDeployment(clientset, resourceName)
			addDeploymentForJson(clientset, resourceName)
		}
	case "delete":
		switch resource {
		case "namespace":
			deleteNamespace(clientset, resourceName)
		case "deployment":
			deleteDeployment(clientset, resourceName)
		}
	case "update":
		switch resource {
		case "deployment":
			updateDeployment(clientset, resourceName)
		}
	case "list":
		switch resource {
		case "namespace":
			listNamespaces(clientset)
		case "deployment":
			listDeployments(clientset, resourceNsName)
		case "pod":
			listPods(clientset, resourceNsName)
		}
	case "get":
		switch resource {
		case "namespace":
			getNsDetail(clientset, resourceName)
			// case "deployment":
			// 	getDeploDetail(clientset, resourceNsName)
			// case "pod":
			// 	getPodDetail(clientset, resourceNsName)
		}
	default:
		fmt.Println("invalid option")
	}
}

func mainBak() {
	selectActions()
}
