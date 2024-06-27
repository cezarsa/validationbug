// +groupName=validationbug
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
	// +kubebuilder:validation:XValidation:rule="size(self.field) > 2",message="validation 1"
	Thing Thing `json:"thing"`
}

// +kubebuilder:validation:XValidation:rule="has(self.field)",message="validation 2"
type Thing struct {
	Field *string `json:"field"`
}
