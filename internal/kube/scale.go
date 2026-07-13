package kube

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func ScaleDeployment(name string, replicas int32) {

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

	deployment, err := clientset.AppsV1().
		Deployments("default").
		Get(context.Background(), name, metav1.GetOptions{})

	if err != nil {
		fmt.Println("Error finding deployment:", err)
		return
	}

	deployment.Spec.Replicas = &replicas

	_, err = clientset.AppsV1().
		Deployments("default").
		Update(context.Background(), deployment, metav1.UpdateOptions{})

	if err != nil {
		fmt.Println("Error scaling deployment:", err)
		return
	}

	fmt.Println("✅ Deployment scaled successfully!")
}