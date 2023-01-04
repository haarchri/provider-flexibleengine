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

type EndpointObservation struct {

	// The unique ID of the VPC endpoint.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The packet ID of the VPC endpoint.
	PacketID *float64 `json:"packetId,omitempty" tf:"packet_id,omitempty"`

	// The domain name for accessing the associated VPC endpoint service.
	// This parameter is only available when enable_dns is set to true.
	PrivateDomainName *string `json:"privateDomainName,omitempty" tf:"private_domain_name,omitempty"`

	// The name of the VPC endpoint service.
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name,omitempty"`

	// The type of the VPC endpoint service.
	ServiceType *string `json:"serviceType,omitempty" tf:"service_type,omitempty"`

	// The status of the VPC endpoint. The value can be accepted, pendingAcceptance or rejected.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type EndpointParameters struct {

	// Specifies whether to create a private domain name. The default value is true.
	// Changing this creates a new VPC endpoint.
	// +kubebuilder:validation:Optional
	EnableDNS *bool `json:"enableDns,omitempty" tf:"enable_dns,omitempty"`

	// Specifies whether to enable access control. The default value is false.
	// Changing this creates a new VPC endpoint.
	// +kubebuilder:validation:Optional
	EnableWhitelist *bool `json:"enableWhitelist,omitempty" tf:"enable_whitelist,omitempty"`

	// Specifies the IP address for accessing the associated VPC endpoint service.
	// Only IPv4 addresses are supported. Changing this creates a new VPC endpoint.
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// Specifies the network ID of the subnet in the VPC specified by vpc_id.
	// Changing this creates a new VPC endpoint.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1.VPCSubnet
	// +crossplane:generate:reference:extractor=github.com/FrangipaneTeam/provider-flexibleengine/config/common.IDExtractor()
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a VPCSubnet in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a VPCSubnet in vpc to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// The region in which to create the VPC endpoint.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the ID of the VPC endpoint service.
	// Changing this creates a new VPC endpoint.
	// +crossplane:generate:reference:type=VPCEPService
	// +kubebuilder:validation:Optional
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`

	// Reference to a VPCEPService to populate serviceId.
	// +kubebuilder:validation:Optional
	ServiceIDRef *v1.Reference `json:"serviceIdRef,omitempty" tf:"-"`

	// Selector for a VPCEPService to populate serviceId.
	// +kubebuilder:validation:Optional
	ServiceIDSelector *v1.Selector `json:"serviceIdSelector,omitempty" tf:"-"`

	// The key/value pairs to associate with the VPC endpoint.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the ID of the VPC where the VPC endpoint is to be created.
	// Changing this creates a new VPC endpoint.
	// +crossplane:generate:reference:type=github.com/FrangipaneTeam/provider-flexibleengine/apis/vpc/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in vpc to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`

	// Specifies the list of IP address or CIDR block which can be accessed to the VPC endpoint.
	// Changing this creates a new VPC endpoint.
	// +kubebuilder:validation:Optional
	Whitelist []*string `json:"whitelist,omitempty" tf:"whitelist,omitempty"`
}

// EndpointSpec defines the desired state of Endpoint
type EndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EndpointParameters `json:"forProvider"`
}

// EndpointStatus defines the observed state of Endpoint.
type EndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Endpoint is the Schema for the Endpoints API. ""page_title: "flexibleengine_vpcep_endpoint"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type Endpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EndpointSpec   `json:"spec"`
	Status            EndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EndpointList contains a list of Endpoints
type EndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Endpoint `json:"items"`
}

// Repository type metadata.
var (
	Endpoint_Kind             = "Endpoint"
	Endpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Endpoint_Kind}.String()
	Endpoint_KindAPIVersion   = Endpoint_Kind + "." + CRDGroupVersion.String()
	Endpoint_GroupVersionKind = CRDGroupVersion.WithKind(Endpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&Endpoint{}, &EndpointList{})
}
