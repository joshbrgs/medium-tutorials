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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TacoOrderSpec defines the desired state of TacoOrder
type TacoOrderSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Type of taco (e.g., al pastor, carne asada, veggie)
	Type string `json:"type,omitempty"`

	// Quantity of tacos to be made
	Quantity int `json:"quantity,omitempty"`

	// Extra toppings like guac, salsa, or cheese
	Extras []string `json:"extras,omitempty"`
}

// TacoOrderStatus defines the observed state of TacoOrder
type TacoOrderStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Served int `json:"served,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// TacoOrder is the Schema for the tacoorders API
type TacoOrder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TacoOrderSpec   `json:"spec,omitempty"`
	Status TacoOrderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TacoOrderList contains a list of TacoOrder
type TacoOrderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TacoOrder `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TacoOrder{}, &TacoOrderList{})
}
