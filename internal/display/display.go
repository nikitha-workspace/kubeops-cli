package display

import (
	"fmt"

	"github.com/nikitha/kubeops-cli/internal/models"
)

func PrintDeployment(deployment models.Deployment) {

	fmt.Println()
	fmt.Println("===== Parsed Deployment =====")

	fmt.Println("API Version :", deployment.APIVersion)
	fmt.Println("Kind        :", deployment.Kind)
	fmt.Println("Name        :", deployment.Metadata.Name)
	fmt.Println("Namespace   :", deployment.Metadata.Namespace)
	fmt.Println("Replicas    :", deployment.Spec.Replicas)
}