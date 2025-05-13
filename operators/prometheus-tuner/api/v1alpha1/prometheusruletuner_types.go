/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PrometheusRuleTunerSpec defines the desired state of PrometheusRuleTuner
type PrometheusRuleTunerSpec struct {
	// DeploymentName is the name of the deployment to monitor
	DeploymentName string `json:"deploymentName"`

	// Namespace is the namespace of the deployment
	Namespace string `json:"namespace"`
}

// PrometheusRuleTunerStatus defines the observed state of PrometheusRuleTuner
type PrometheusRuleTunerStatus struct {
	// CPUUsage in millicores (e.g., "250m")
	CPUUsage string `json:"cpuUsage,omitempty"`

	// MemoryUsage in MiB (e.g., "128Mi")
	MemoryUsage string `json:"memoryUsage,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// PrometheusRuleTuner is the Schema for the prometheusruletuners API
type PrometheusRuleTuner struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PrometheusRuleTunerSpec   `json:"spec,omitempty"`
	Status PrometheusRuleTunerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrometheusRuleTunerList contains a list of PrometheusRuleTuner
type PrometheusRuleTunerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrometheusRuleTuner `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PrometheusRuleTuner{}, &PrometheusRuleTunerList{})
}
