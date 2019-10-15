package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type UnoperatedInstance struct {
	ManagementSecretName string `json:"managementSecretName"`
	TargetName           string `json:"targetName"`
	TargetPort           int32  `json:"targetPort"`
}

// DatabaseInstanceSpec defines the desired state of DatabaseInstance
// +k8s:openapi-gen=true
type DatabaseInstanceSpec struct {
	UnoperatedInstance UnoperatedInstance `json:"unoperatedInstance"`
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// DatabaseInstanceStatus defines the observed state of DatabaseInstance
// +k8s:openapi-gen=true
type DatabaseInstanceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DatabaseInstance is the Schema for the databaseinstances API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=databaseinstances,scope=Namespaced
type DatabaseInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatabaseInstanceSpec   `json:"spec,omitempty"`
	Status DatabaseInstanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DatabaseInstanceList contains a list of DatabaseInstance
type DatabaseInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DatabaseInstance{}, &DatabaseInstanceList{})
}
