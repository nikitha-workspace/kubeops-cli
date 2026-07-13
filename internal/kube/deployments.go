package kube

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func ListDeployments() {

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

	deployments, err := clientset.AppsV1().
		Deployments("default").
		List(context.Background(), metav1.ListOptions{})

	if err != nil {
		fmt.Println("Error getting deployments:", err)
		return
	}

	fmt.Println("===== Deployments =====")

	if len(deployments.Items) == 0 {
		fmt.Println("No deployments found.")
		return
	} // <-- THIS BRACE WAS MISSING

	for _, deployment := range deployments.Items {

		fmt.Println()
		fmt.Println("Deployment Name :", deployment.Name)

		if deployment.Spec.Replicas != nil {
			fmt.Println("Desired Replicas :", *deployment.Spec.Replicas)
		}

		fmt.Println("Available Replicas :", deployment.Status.AvailableReplicas)
	}
}