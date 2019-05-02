package v1beta1

import (
	"k8s.io/apimachinery/pkg/api/resource"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TrafficSplitBackend defines a backend
// +k8s:openapi-gen=true
type TrafficSplitBackend struct {
	Service string             `json:"service,omitempty"`
	Weight  *resource.Quantity `json:"weight,omitempty"`
}

// TrafficSplitSpec defines the desired state of TrafficSplit
// +k8s:openapi-gen=true
type TrafficSplitSpec struct {
	Service  string                `json:"service,omitempty"`
	Backends []TrafficSplitBackend `json:"backends,omitempty"`
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// TrafficSplitStatus defines the observed state of TrafficSplit
// +k8s:openapi-gen=true
type TrafficSplitStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TrafficSplit is the Schema for the trafficsplits API
// +k8s:openapi-gen=true
type TrafficSplit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TrafficSplitSpec   `json:"spec,omitempty"`
	Status TrafficSplitStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TrafficSplitList contains a list of TrafficSplit
type TrafficSplitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrafficSplit `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TrafficSplit{}, &TrafficSplitList{})
}
