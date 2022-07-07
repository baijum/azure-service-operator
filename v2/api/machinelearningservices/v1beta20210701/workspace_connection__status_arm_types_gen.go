// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210701

type WorkspaceConnection_StatusARM struct {
	// Id: ResourceId of the workspace connection.
	Id *string `json:"id,omitempty"`

	// Name: Friendly name of the workspace connection.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of workspace connection.
	Properties *WorkspaceConnectionProps_StatusARM `json:"properties,omitempty"`

	// Type: Resource type of workspace connection.
	Type *string `json:"type,omitempty"`
}

type WorkspaceConnectionProps_StatusARM struct {
	// AuthType: Authorization type of the workspace connection.
	AuthType *string `json:"authType,omitempty"`

	// Category: Category of the workspace connection.
	Category *string `json:"category,omitempty"`

	// Target: Target of the workspace connection.
	Target *string `json:"target,omitempty"`

	// Value: Value details of the workspace connection.
	Value *string `json:"value,omitempty"`

	// ValueFormat: format for the workspace connection value
	ValueFormat *WorkspaceConnectionPropsStatusValueFormat `json:"valueFormat,omitempty"`
}

type WorkspaceConnectionPropsStatusValueFormat string

const WorkspaceConnectionPropsStatusValueFormatJSON = WorkspaceConnectionPropsStatusValueFormat("JSON")
