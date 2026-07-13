package kube

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func CreateDeployment(deployment *appsv1.Deployment) {

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

	namespace := deployment.Namespace

	if namespace == "" {
		namespace = "default"
	}

	_, err = clientset.AppsV1().
		Deployments(namespace).
		Create(context.Background(), deployment, metav1.CreateOptions{})

	if err != nil {
		fmt.Println("Error creating deployment:", err)
		return
	}

	fmt.Println("✅ Deployment created successfully!")
}