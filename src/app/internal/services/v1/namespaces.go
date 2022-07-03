package services

import (
	"context"
	"dashboard/src/app/internal/cli/kubernetes"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//getNamespace获取 namespace
func GetNamespace() []string {
	kubeCli := kubernetes.KubeCLI()
	namespaceLists, err := kubeCli.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("namespace err:", err)
	}

	res := make([]string, 0)

	for _, v := range namespaceLists.Items {
		res = append(res, v.Name)
	}
	return res
}
