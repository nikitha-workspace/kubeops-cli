package cmd

import "github.com/nikitha/kubeops-cli/internal/kube"

func Status(name string) {
	kube.DeploymentStatus(name)
}