/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ServiceAccountIAMPolicyObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ServiceAccountIAMPolicyParameters struct {

	// +kubebuilder:validation:Required
	PolicyData *string `json:"policyData" tf:"policy_data,omitempty"`

	// +crossplane:generate:reference:type=ServiceAccount
	// +crossplane:generate:reference:refFieldName=ServiceAccountRef
	// +crossplane:generate:reference:selectorFieldName=ServiceAccountSelector
	// +kubebuilder:validation:Optional
	ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceAccountRef *v1.Reference `json:"serviceAccountRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ServiceAccountSelector *v1.Selector `json:"serviceAccountSelector,omitempty" tf:"-"`
}

// ServiceAccountIAMPolicySpec defines the desired state of ServiceAccountIAMPolicy
type ServiceAccountIAMPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceAccountIAMPolicyParameters `json:"forProvider"`
}

// ServiceAccountIAMPolicyStatus defines the observed state of ServiceAccountIAMPolicy.
type ServiceAccountIAMPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceAccountIAMPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceAccountIAMPolicy is the Schema for the ServiceAccountIAMPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ServiceAccountIAMPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceAccountIAMPolicySpec   `json:"spec"`
	Status            ServiceAccountIAMPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceAccountIAMPolicyList contains a list of ServiceAccountIAMPolicys
type ServiceAccountIAMPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceAccountIAMPolicy `json:"items"`
}

// Repository type metadata.
var (
	ServiceAccountIAMPolicy_Kind             = "ServiceAccountIAMPolicy"
	ServiceAccountIAMPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceAccountIAMPolicy_Kind}.String()
	ServiceAccountIAMPolicy_KindAPIVersion   = ServiceAccountIAMPolicy_Kind + "." + CRDGroupVersion.String()
	ServiceAccountIAMPolicy_GroupVersionKind = CRDGroupVersion.WithKind(ServiceAccountIAMPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceAccountIAMPolicy{}, &ServiceAccountIAMPolicyList{})
}
