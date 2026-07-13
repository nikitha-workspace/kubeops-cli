package cmd

import (
	"fmt"

	"github.com/nikitha/kubeops-cli/internal/kube"
	
	"github.com/nikitha/kubeops-cli/internal/validator"
	"github.com/nikitha/kubeops-cli/internal/parser"
	
)

func Deploy(fileName string){
	fmt.Println("====================================")
	fmt.Println("Deployment Started")
	fmt.Println("Deployment File :", fileName)
	fmt.Println("====================================")

	deployment, err := parser.ParseDeployment(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println()
	fmt.Println("===== Parsed Deployment =====")
	fmt.Println("API Version :", deployment.APIVersion)
	fmt.Println("Kind        :", deployment.Kind)

	fmt.Println("Name        :", deployment.Metadata.Name)
	fmt.Println("Namespace   :", deployment.Metadata.Namespace)

	fmt.Println("Replicas    :", deployment.Spec.Replicas)

	
	fmt.Println()
	fmt.Println("===== Container Details =====")

	if len(deployment.Spec.Template.Spec.Containers) == 0 {
		fmt.Println("No containers found in deployment.")
		return
	}

	container := deployment.Spec.Template.Spec.Containers[0]

	fmt.Println("Container Name :", container.Name)
	fmt.Println("Image          :", container.Image)

	if len(container.Ports) > 0 {
		fmt.Println("Port           :", container.Ports[0].ContainerPort)
	} else {
		fmt.Println("Port           :, No Ports defined")
	}
	

	validator.ValidateDeployment(deployment)

	k8sDeployment := kube.ConvertDeployment(deployment)
	
	kube.CreateDeployment(k8sDeployment)

 
}