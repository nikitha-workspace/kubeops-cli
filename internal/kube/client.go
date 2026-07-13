package kube

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func Connect() {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	
    if err != nil {
		fmt.Println("Error loading kubeconfig:", err)
	    return
    }

    clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		fmt.Println("Error creating Kubernetes client:", err)
		return
	}

	
    fmt.Println("====================================")
	fmt.Println("Connected to Kubernetes Cluster")
	fmt.Println("API Server :", config.Host)
	fmt.Println("====================================")

	namespaces, err := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})

    if err != nil {
	    fmt.Println("Error getting namespaces:", err)
	    return
	}

    fmt.Println()
	fmt.Println("Namespaces:") 

    for _, ns := range namespaces.Items {
	    fmt.Println("-", ns.Name)
    }
	 
}