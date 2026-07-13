package kube

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func DeploymentStatus(name string) {

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
		fmt.Println("Error getting deployment:", err)
		return
	}

	fmt.Println("===== Deployment Status =====")
	fmt.Println("Name                :", deployment.Name)
	fmt.Println("Namespace           :", deployment.Namespace)

	if deployment.Spec.Replicas != nil {
		fmt.Println("Desired Replicas    :", *deployment.Spec.Replicas)
	}

	fmt.Println("Ready Replicas      :", deployment.Status.ReadyReplicas)
	fmt.Println("Available Replicas  :", deployment.Status.AvailableReplicas)

	if len(deployment.Spec.Template.Spec.Containers) > 0 {
		fmt.Println("Image               :", deployment.Spec.Template.Spec.Containers[0].Image)
	}
}