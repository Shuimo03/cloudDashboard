package services

import (
	"context"
	"dashboard/app/internal/cli/kubernetes"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetPodsNamesService 获取命名空间下所有的pod
func GetPodsNamesList(namespace string) []string {

	if namespace == "" {
		namespace = "default"
	}

	kubeCli := kubernetes.KubeCLI()
	res := make([]string, 0)
	podsList, err := kubeCli.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Pod or NameSpace Error", err)
	}
	for _, v := range podsList.Items {
		res = append(res, v.Name)
	}
	return res
}

// DeletePod删除指定namespace下的pods
// 这里最好是使用map,现在只是为了实现功能,所以使用了切片
func DeletePod(namespace, podName string) string {
	kubeCli := kubernetes.KubeCLI()
	namespaceList := GetNamespace()
	deleteNameSpace := ""
	for i := 0; i < len(namespaceList); i++ {
		if namespace == namespaceList[i] {
			deleteNameSpace = namespace
		}
	}
	err := kubeCli.CoreV1().Pods(deleteNameSpace).Delete(context.TODO(), podName, metav1.DeleteOptions{})
	if err != nil {
		fmt.Println("Pod delete fail")
	}
	return "删除成功"
}
