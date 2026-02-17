/*
Copyright 2026.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GreetingSpec defines the desired state of Greeting
type GreetingSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// The following markers will use OpenAPI v3 schema to validate the value
	// More info: https://book.kubebuilder.io/reference/markers/crd-validation.html

	// Name to greet (required)
	// +kubebuilder:validation:Required
	Name string `json:"name"`

	// Greeting word (optional, defaults to "Hello")
	// +kubebuilder:default="Hello"
	Greeting string `json:"greeting,omitempty"`
}

// GreetingStatus defines the observed state of Greeting // OR WHAT THE CONTROLLER WRITES BACK
type GreetingStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// The composed message
	Message string `json:"message,omitempty"`

	// Whether reconciliation succeeded
	Ready bool `json:"ready"`
}

// Greeting
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
// +kubebuilder:printcolumn:name="Ready",type="boolean",JSONPath=".status.ready"
// Greeting is the Schema for the greetings API
type Greeting struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`

	// spec defines the desired state of Greeting
	// +required
	Spec GreetingSpec `json:"spec"`

	// status defines the observed state of Greeting
	// +optional
	Status GreetingStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// GreetingList contains a list of Greeting
type GreetingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []Greeting `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Greeting{}, &GreetingList{})
}
