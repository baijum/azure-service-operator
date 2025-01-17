// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

// Deprecated version of Subnet_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded. Use v1beta20201101.Subnet_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded instead
type Subnet_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded_ARM struct {
	Etag       *string                            `json:"etag,omitempty"`
	Id         *string                            `json:"id,omitempty"`
	Name       *string                            `json:"name,omitempty"`
	Properties *SubnetPropertiesFormat_STATUS_ARM `json:"properties,omitempty"`
	Type       *string                            `json:"type,omitempty"`
}

// Deprecated version of SubnetPropertiesFormat_STATUS. Use v1beta20201101.SubnetPropertiesFormat_STATUS instead
type SubnetPropertiesFormat_STATUS_ARM struct {
	AddressPrefix                      *string                                                                        `json:"addressPrefix,omitempty"`
	AddressPrefixes                    []string                                                                       `json:"addressPrefixes,omitempty"`
	ApplicationGatewayIpConfigurations []ApplicationGatewayIPConfiguration_STATUS_ARM                                 `json:"applicationGatewayIpConfigurations,omitempty"`
	Delegations                        []Delegation_STATUS_ARM                                                        `json:"delegations,omitempty"`
	IpAllocations                      []SubResource_STATUS_ARM                                                       `json:"ipAllocations,omitempty"`
	IpConfigurationProfiles            []IPConfigurationProfile_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded_ARM `json:"ipConfigurationProfiles,omitempty"`
	IpConfigurations                   []IPConfiguration_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded_ARM        `json:"ipConfigurations,omitempty"`
	NatGateway                         *SubResource_STATUS_ARM                                                        `json:"natGateway,omitempty"`
	NetworkSecurityGroup               *NetworkSecurityGroup_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded_ARM    `json:"networkSecurityGroup,omitempty"`
	PrivateEndpointNetworkPolicies     *SubnetPropertiesFormat_PrivateEndpointNetworkPolicies_STATUS                  `json:"privateEndpointNetworkPolicies,omitempty"`
	PrivateEndpoints                   []PrivateEndpoint_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded_ARM        `json:"privateEndpoints,omitempty"`
	PrivateLinkServiceNetworkPolicies  *SubnetPropertiesFormat_PrivateLinkServiceNetworkPolicies_STATUS               `json:"privateLinkServiceNetworkPolicies,omitempty"`
	ProvisioningState                  *ProvisioningState_STATUS                                                      `json:"provisioningState,omitempty"`
	Purpose                            *string                                                                        `json:"purpose,omitempty"`
	ResourceNavigationLinks            []ResourceNavigationLink_STATUS_ARM                                            `json:"resourceNavigationLinks,omitempty"`
	RouteTable                         *RouteTable_STATUS_SubResourceEmbedded_ARM                                     `json:"routeTable,omitempty"`
	ServiceAssociationLinks            []ServiceAssociationLink_STATUS_ARM                                            `json:"serviceAssociationLinks,omitempty"`
	ServiceEndpointPolicies            []ServiceEndpointPolicy_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded_ARM  `json:"serviceEndpointPolicies,omitempty"`
	ServiceEndpoints                   []ServiceEndpointPropertiesFormat_STATUS_ARM                                   `json:"serviceEndpoints,omitempty"`
}

// Deprecated version of ApplicationGatewayIPConfiguration_STATUS. Use v1beta20201101.ApplicationGatewayIPConfiguration_STATUS instead
type ApplicationGatewayIPConfiguration_STATUS_ARM struct {
	Etag       *string                                                       `json:"etag,omitempty"`
	Id         *string                                                       `json:"id,omitempty"`
	Name       *string                                                       `json:"name,omitempty"`
	Properties *ApplicationGatewayIPConfigurationPropertiesFormat_STATUS_ARM `json:"properties,omitempty"`
	Type       *string                                                       `json:"type,omitempty"`
}

// Deprecated version of Delegation_STATUS. Use v1beta20201101.Delegation_STATUS instead
type Delegation_STATUS_ARM struct {
	Etag       *string                                       `json:"etag,omitempty"`
	Id         *string                                       `json:"id,omitempty"`
	Name       *string                                       `json:"name,omitempty"`
	Properties *ServiceDelegationPropertiesFormat_STATUS_ARM `json:"properties,omitempty"`
	Type       *string                                       `json:"type,omitempty"`
}

// Deprecated version of IPConfiguration_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded. Use v1beta20201101.IPConfiguration_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded instead
type IPConfiguration_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded_ARM struct {
	Etag       *string                                                                                `json:"etag,omitempty"`
	Id         *string                                                                                `json:"id,omitempty"`
	Name       *string                                                                                `json:"name,omitempty"`
	Properties *IPConfigurationPropertiesFormat_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded_ARM `json:"properties,omitempty"`
}

// Deprecated version of IPConfigurationProfile_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded. Use v1beta20201101.IPConfigurationProfile_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded instead
type IPConfigurationProfile_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded_ARM struct {
	Etag       *string                                                                                       `json:"etag,omitempty"`
	Id         *string                                                                                       `json:"id,omitempty"`
	Name       *string                                                                                       `json:"name,omitempty"`
	Properties *IPConfigurationProfilePropertiesFormat_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded_ARM `json:"properties,omitempty"`
	Type       *string                                                                                       `json:"type,omitempty"`
}

// Deprecated version of NetworkSecurityGroup_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded. Use v1beta20201101.NetworkSecurityGroup_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded instead
type NetworkSecurityGroup_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded_ARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of PrivateEndpoint_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded. Use v1beta20201101.PrivateEndpoint_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded instead
type PrivateEndpoint_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded_ARM struct {
	ExtendedLocation *ExtendedLocation_STATUS_ARM `json:"extendedLocation,omitempty"`
	Id               *string                      `json:"id,omitempty"`
}

// Deprecated version of ResourceNavigationLink_STATUS. Use v1beta20201101.ResourceNavigationLink_STATUS instead
type ResourceNavigationLink_STATUS_ARM struct {
	Etag       *string                                  `json:"etag,omitempty"`
	Id         *string                                  `json:"id,omitempty"`
	Name       *string                                  `json:"name,omitempty"`
	Properties *ResourceNavigationLinkFormat_STATUS_ARM `json:"properties,omitempty"`
	Type       *string                                  `json:"type,omitempty"`
}

// Deprecated version of RouteTable_STATUS_SubResourceEmbedded. Use v1beta20201101.RouteTable_STATUS_SubResourceEmbedded instead
type RouteTable_STATUS_SubResourceEmbedded_ARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of ServiceAssociationLink_STATUS. Use v1beta20201101.ServiceAssociationLink_STATUS instead
type ServiceAssociationLink_STATUS_ARM struct {
	Etag       *string                                            `json:"etag,omitempty"`
	Id         *string                                            `json:"id,omitempty"`
	Name       *string                                            `json:"name,omitempty"`
	Properties *ServiceAssociationLinkPropertiesFormat_STATUS_ARM `json:"properties,omitempty"`
	Type       *string                                            `json:"type,omitempty"`
}

// Deprecated version of ServiceEndpointPolicy_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded. Use v1beta20201101.ServiceEndpointPolicy_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded instead
type ServiceEndpointPolicy_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded_ARM struct {
	Id   *string `json:"id,omitempty"`
	Kind *string `json:"kind,omitempty"`
}

// Deprecated version of ServiceEndpointPropertiesFormat_STATUS. Use v1beta20201101.ServiceEndpointPropertiesFormat_STATUS instead
type ServiceEndpointPropertiesFormat_STATUS_ARM struct {
	Locations         []string                  `json:"locations,omitempty"`
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`
	Service           *string                   `json:"service,omitempty"`
}

// Deprecated version of SubnetPropertiesFormat_PrivateEndpointNetworkPolicies_STATUS. Use
// v1beta20201101.SubnetPropertiesFormat_PrivateEndpointNetworkPolicies_STATUS instead
type SubnetPropertiesFormat_PrivateEndpointNetworkPolicies_STATUS string

const (
	SubnetPropertiesFormat_PrivateEndpointNetworkPolicies_STATUS_Disabled = SubnetPropertiesFormat_PrivateEndpointNetworkPolicies_STATUS("Disabled")
	SubnetPropertiesFormat_PrivateEndpointNetworkPolicies_STATUS_Enabled  = SubnetPropertiesFormat_PrivateEndpointNetworkPolicies_STATUS("Enabled")
)

// Deprecated version of SubnetPropertiesFormat_PrivateLinkServiceNetworkPolicies_STATUS. Use
// v1beta20201101.SubnetPropertiesFormat_PrivateLinkServiceNetworkPolicies_STATUS instead
type SubnetPropertiesFormat_PrivateLinkServiceNetworkPolicies_STATUS string

const (
	SubnetPropertiesFormat_PrivateLinkServiceNetworkPolicies_STATUS_Disabled = SubnetPropertiesFormat_PrivateLinkServiceNetworkPolicies_STATUS("Disabled")
	SubnetPropertiesFormat_PrivateLinkServiceNetworkPolicies_STATUS_Enabled  = SubnetPropertiesFormat_PrivateLinkServiceNetworkPolicies_STATUS("Enabled")
)

// Deprecated version of ApplicationGatewayIPConfigurationPropertiesFormat_STATUS. Use v1beta20201101.ApplicationGatewayIPConfigurationPropertiesFormat_STATUS instead
type ApplicationGatewayIPConfigurationPropertiesFormat_STATUS_ARM struct {
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`
	Subnet            *SubResource_STATUS_ARM   `json:"subnet,omitempty"`
}

// Deprecated version of IPConfigurationProfilePropertiesFormat_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded. Use v1beta20201101.IPConfigurationProfilePropertiesFormat_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded instead
type IPConfigurationProfilePropertiesFormat_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded_ARM struct {
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`
}

// Deprecated version of IPConfigurationPropertiesFormat_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded. Use v1beta20201101.IPConfigurationPropertiesFormat_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded instead
type IPConfigurationPropertiesFormat_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded_ARM struct {
	PrivateIPAddress          *string                                                                `json:"privateIPAddress,omitempty"`
	PrivateIPAllocationMethod *IPAllocationMethod_STATUS                                             `json:"privateIPAllocationMethod,omitempty"`
	ProvisioningState         *ProvisioningState_STATUS                                              `json:"provisioningState,omitempty"`
	PublicIPAddress           *PublicIPAddress_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded_ARM `json:"publicIPAddress,omitempty"`
}

// Deprecated version of ResourceNavigationLinkFormat_STATUS. Use v1beta20201101.ResourceNavigationLinkFormat_STATUS instead
type ResourceNavigationLinkFormat_STATUS_ARM struct {
	Link               *string                   `json:"link,omitempty"`
	LinkedResourceType *string                   `json:"linkedResourceType,omitempty"`
	ProvisioningState  *ProvisioningState_STATUS `json:"provisioningState,omitempty"`
}

// Deprecated version of ServiceAssociationLinkPropertiesFormat_STATUS. Use v1beta20201101.ServiceAssociationLinkPropertiesFormat_STATUS instead
type ServiceAssociationLinkPropertiesFormat_STATUS_ARM struct {
	AllowDelete        *bool                     `json:"allowDelete,omitempty"`
	Link               *string                   `json:"link,omitempty"`
	LinkedResourceType *string                   `json:"linkedResourceType,omitempty"`
	Locations          []string                  `json:"locations,omitempty"`
	ProvisioningState  *ProvisioningState_STATUS `json:"provisioningState,omitempty"`
}

// Deprecated version of ServiceDelegationPropertiesFormat_STATUS. Use v1beta20201101.ServiceDelegationPropertiesFormat_STATUS instead
type ServiceDelegationPropertiesFormat_STATUS_ARM struct {
	Actions           []string                  `json:"actions,omitempty"`
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`
	ServiceName       *string                   `json:"serviceName,omitempty"`
}

// Deprecated version of PublicIPAddress_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded. Use v1beta20201101.PublicIPAddress_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded instead
type PublicIPAddress_STATUS_VirtualNetworks_Subnet_SubResourceEmbedded_ARM struct {
	ExtendedLocation *ExtendedLocation_STATUS_ARM   `json:"extendedLocation,omitempty"`
	Id               *string                        `json:"id,omitempty"`
	Sku              *PublicIPAddressSku_STATUS_ARM `json:"sku,omitempty"`
	Zones            []string                       `json:"zones,omitempty"`
}
