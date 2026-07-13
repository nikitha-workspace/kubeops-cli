package cmd

import "github.com/nikitha/kubeops-cli/internal/kube"

func Scale(name string, replicas int32) {
	kube.ScaleDeployment(name, replicas)
}