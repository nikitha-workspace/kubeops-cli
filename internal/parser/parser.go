package parser

import (
	"os"

	"github.com/nikitha/kubeops-cli/internal/models"
	"gopkg.in/yaml.v3"
)

func ParseDeployment(fileName string) (models.Deployment, error) {

	var deployment models.Deployment

	data, err := os.ReadFile(fileName)

	if err != nil {
		return deployment, err
	}

	err = yaml.Unmarshal(data, &deployment)

	if err != nil {
		return deployment, err
	}

	return deployment, nil
}