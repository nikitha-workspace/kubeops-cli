package kube

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func DeleteDeployment(name string) {

	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		fmt.Println("Error loading kubeconfig:", err)
		return
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("Error creating client:", err)
		return
	}

	err = clientset.AppsV1().
		Deployments("default").
		Delete(context.Background(), name, metav1.DeleteOptions{})

	if err != nil {
		fmt.Println("Error deleting deployment:", err)
		return
	}

	fmt.Println("✅ Deployment deleted successfully!")
}