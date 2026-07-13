package cmd

import "github.com/nikitha/kubeops-cli/internal/kube"

func Deployments() {
	kube.ListDeployments()
}