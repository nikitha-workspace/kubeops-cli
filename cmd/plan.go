package cmd

import (
	"fmt"

	"github.com/nikitha/kubeops-cli/internal/parser"
	"github.com/nikitha/kubeops-cli/internal/validator"
)

func Plan(fileName string) {

	fmt.Println("====================================")
	fmt.Println("Deployment Plan")
	fmt.Println("Deployment File :", fileName)
	fmt.Println("====================================")

	deployment, err := parser.ParseDeployment(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	validator.ValidateDeployment(deployment)

	fmt.Println()
	fmt.Println("========== PLAN ==========")
	fmt.Println("Deployment :", deployment.Metadata.Name)
	fmt.Println("Namespace  :", deployment.Metadata.Namespace)
	fmt.Println("Replicas   :", deployment.Spec.Replicas)

	container := deployment.Spec.Template.Spec.Containers[0]

	fmt.Println("Container  :", container.Name)
	fmt.Println("Image      :", container.Image)

	if len(container.Ports) > 0 {
		fmt.Println("Port       :", container.Ports[0].ContainerPort)

	fmt.Println()
	fmt.Println("===== Deployment Summary =====")

    fmt.Println("Pods to Create :", deployment.Spec.Replicas)
    fmt.Println("Containers     :", len(deployment.Spec.Template.Spec.Containers))
    fmt.Println("Images Used    :")
    for _, c := range deployment.Spec.Template.Spec.Containers {
	fmt.Println(" -", c.Image)
	}
	}
}