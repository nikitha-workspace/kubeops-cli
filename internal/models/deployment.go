package models

type Metadata struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
}

type Spec struct {
	Replicas int          `yaml:"replicas"`
	Template TemplateSpec `yaml:"template"`
}

type Deployment struct {
	APIVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}

type Port struct {
	ContainerPort int `yaml:"containerPort"`
}

type Container struct {
	Name  string `yaml:"name"`
	Image string `yaml:"image"`
	Ports []Port `yaml:"ports"`
}

type PodSpec struct {
	Containers []Container `yaml:"containers"`
}

type TemplateSpec struct {
	Spec PodSpec `yaml:"spec"`
}