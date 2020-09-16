package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
)

// RocketmqExporter defines the specification for the rocketmq exporter
type RocketmqExporter struct {
	Enabled bool `json:"enabled,omitempty"`
	Image string `json:"image,omitempty"`
	ImagePullPolicy corev1.PullPolicy `json:"imagePullPolicy,omitempty"`
}
