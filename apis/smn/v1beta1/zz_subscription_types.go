/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SubscriptionObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SubscriptionParameters struct {

	// Message endpoint.
	// For an HTTP subscription, the endpoint starts with http://.
	// For an HTTPS subscription, the endpoint starts with https://.
	// For an email subscription, the endpoint is a mail address.
	// For an SMS message subscription, the endpoint is a phone number.
	// +kubebuilder:validation:Required
	Endpoint *string `json:"endpoint" tf:"endpoint,omitempty"`

	// Project ID of the topic creator.
	// +kubebuilder:validation:Optional
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// Protocol of the message endpoint. Currently, email,
	// sms, http, and https are supported.
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// Remark information. The remarks must be a UTF-8-coded
	// character string containing 128 bytes.
	// +kubebuilder:validation:Optional
	Remark *string `json:"remark,omitempty" tf:"remark,omitempty"`

	// Subscription status.
	// 0 indicates that the subscription is not confirmed.
	// 1 indicates that the subscription is confirmed.
	// 3 indicates that the subscription is canceled.
	// +kubebuilder:validation:Optional
	Status *float64 `json:"status,omitempty" tf:"status,omitempty"`

	// Resource identifier of a subscription, which
	// is unique.
	// +kubebuilder:validation:Optional
	SubscriptionUrn *string `json:"subscriptionUrn,omitempty" tf:"subscription_urn,omitempty"`

	// Resource identifier of a topic, which is unique.
	// +kubebuilder:validation:Required
	TopicUrn *string `json:"topicUrn" tf:"topic_urn,omitempty"`
}

// SubscriptionSpec defines the desired state of Subscription
type SubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubscriptionParameters `json:"forProvider"`
}

// SubscriptionStatus defines the observed state of Subscription.
type SubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Subscription is the Schema for the Subscriptions API. ""page_title: "flexibleengine_smn_subscription_v2"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type Subscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubscriptionSpec   `json:"spec"`
	Status            SubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubscriptionList contains a list of Subscriptions
type SubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subscription `json:"items"`
}

// Repository type metadata.
var (
	Subscription_Kind             = "Subscription"
	Subscription_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Subscription_Kind}.String()
	Subscription_KindAPIVersion   = Subscription_Kind + "." + CRDGroupVersion.String()
	Subscription_GroupVersionKind = CRDGroupVersion.WithKind(Subscription_Kind)
)

func init() {
	SchemeBuilder.Register(&Subscription{}, &SubscriptionList{})
}
