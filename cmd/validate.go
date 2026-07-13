package cmd

import (
	"fmt"

	"github.com/nikitha/kubeops-cli/internal/parser"
	"github.com/nikitha/kubeops-cli/internal/validator"
)

func Validate(fileName string) {

	fmt.Println("====================================")
	fmt.Println("Validate Command")
	fmt.Println("Deployment File :",fileName)
	fmt.Println("====================================")

	deployment, err := parser.ParseDeployment(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	validator.ValidateDeployment(deployment)
}