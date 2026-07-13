package kube

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/nikitha/kubeops-cli/internal/models"
)

func ConvertDeployment(d models.Deployment) *appsv1.Deployment {

	replicas := int32(d.Spec.Replicas)

	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      d.Metadata.Name,
			Namespace: d.Metadata.Namespace,
		},

		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,

			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": d.Metadata.Name,
				},
			},

			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": d.Metadata.Name,
					},
				},

				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  d.Spec.Template.Spec.Containers[0].Name,
							Image: d.Spec.Template.Spec.Containers[0].Image,
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: int32(
										d.Spec.Template.Spec.Containers[0].Ports[0].ContainerPort,
									),
								},
							},
						},
					},
				},
			},
		},
	}
}