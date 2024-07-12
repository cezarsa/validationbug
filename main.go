// +groupName=validation.bug
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//go:generate controller-gen crd paths="." output:crd:dir=crd/

// +kubebuilder:object:root=true
type Foo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FooSpec `json:"spec,omitempty"`
}

type FooSpec struct {
	// +kubebuilder:validation:XValidation:rule="true",message="validation 2"
	Thing Thing `json:"thing"`
}

// +kubebuilder:validation:XValidation:rule="true",message="validation 1"
type Thing struct {
	Field *string `json:"field"`
}
