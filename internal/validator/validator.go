package validator

import (
	"fmt"

	"github.com/nikitha/kubeops-cli/internal/models"
)

func ValidateDeployment(deployment models.Deployment) {

	fmt.Println()
	fmt.Println("===== Validating Deployment =====")

	if deployment.APIVersion == "" {
		fmt.Println("❌ API Version is missing")
		return
	}
	fmt.Println("✅ API Version is valid")

	if deployment.Kind != "Deployment" {
		fmt.Println("❌ kind must be 'Deployment'")
		return
	}
	fmt.Println("✅ kind is valid")

	if deployment.Metadata.Name == "" {
		fmt.Println("❌ Deployment Name is missing")
		return
	}
	fmt.Println("✅ Deployment Name is valid")

	if deployment.Spec.Replicas <= 0 {
		fmt.Println("❌ Replicas must be greater than 0")
		return
	}
	fmt.Println("✅ Replicas are valid")


}