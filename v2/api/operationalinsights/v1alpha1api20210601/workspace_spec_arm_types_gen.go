// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// Deprecated version of Workspace_Spec. Use v1beta20210601.Workspace_Spec instead
type Workspace_Spec_ARM struct {
	Etag       *string                  `json:"etag,omitempty"`
	Location   *string                  `json:"location,omitempty"`
	Name       string                   `json:"name,omitempty"`
	Properties *WorkspaceProperties_ARM `json:"properties,omitempty"`
	Tags       map[string]string        `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Workspace_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (workspace Workspace_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (workspace *Workspace_Spec_ARM) GetName() string {
	return workspace.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.OperationalInsights/workspaces"
func (workspace *Workspace_Spec_ARM) GetType() string {
	return "Microsoft.OperationalInsights/workspaces"
}

// Deprecated version of WorkspaceProperties. Use v1beta20210601.WorkspaceProperties instead
type WorkspaceProperties_ARM struct {
	Features                        *WorkspaceFeatures_ARM                               `json:"features,omitempty"`
	ForceCmkForQuery                *bool                                                `json:"forceCmkForQuery,omitempty"`
	ProvisioningState               *WorkspaceProperties_ProvisioningState               `json:"provisioningState,omitempty"`
	PublicNetworkAccessForIngestion *WorkspaceProperties_PublicNetworkAccessForIngestion `json:"publicNetworkAccessForIngestion,omitempty"`
	PublicNetworkAccessForQuery     *WorkspaceProperties_PublicNetworkAccessForQuery     `json:"publicNetworkAccessForQuery,omitempty"`
	RetentionInDays                 *int                                                 `json:"retentionInDays,omitempty"`
	Sku                             *WorkspaceSku_ARM                                    `json:"sku,omitempty"`
	WorkspaceCapping                *WorkspaceCapping_ARM                                `json:"workspaceCapping,omitempty"`
}

// Deprecated version of WorkspaceCapping. Use v1beta20210601.WorkspaceCapping instead
type WorkspaceCapping_ARM struct {
	DailyQuotaGb *float64 `json:"dailyQuotaGb,omitempty"`
}

// Deprecated version of WorkspaceFeatures. Use v1beta20210601.WorkspaceFeatures instead
type WorkspaceFeatures_ARM struct {
	AdditionalProperties                        map[string]v1.JSON `json:"additionalProperties,omitempty"`
	ClusterResourceId                           *string            `json:"clusterResourceId,omitempty"`
	DisableLocalAuth                            *bool              `json:"disableLocalAuth,omitempty"`
	EnableDataExport                            *bool              `json:"enableDataExport,omitempty"`
	EnableLogAccessUsingOnlyResourcePermissions *bool              `json:"enableLogAccessUsingOnlyResourcePermissions,omitempty"`
	ImmediatePurgeDataOn30Days                  *bool              `json:"immediatePurgeDataOn30Days,omitempty"`
}

// Deprecated version of WorkspaceSku. Use v1beta20210601.WorkspaceSku instead
type WorkspaceSku_ARM struct {
	CapacityReservationLevel *int               `json:"capacityReservationLevel,omitempty"`
	Name                     *WorkspaceSku_Name `json:"name,omitempty"`
}
