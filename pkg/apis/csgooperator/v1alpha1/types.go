package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CSGOServer is a specification for a CSGOServer resource
type CSGOServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CSGOServerSpec   `json:"spec"`
	Status CSGOServerStatus `json:"status"`
}

// CSGOServerSpec is the spec for a CSGOServer resource
type CSGOServerSpec struct {
	DeploymentName string `json:"deploymentName"`
	Replicas       *int32 `json:"replicas"`
}

// CSGOServerStatus is the status for a CSGOServer resource
type CSGOServerStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CSGOServerList is a list of CSGOServer resources
type CSGOServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []CSGOServer `json:"items"`
}
