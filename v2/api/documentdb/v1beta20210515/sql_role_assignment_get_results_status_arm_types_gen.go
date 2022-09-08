// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515

type SqlRoleAssignmentGetResults_STATUSARM struct {
	// Id: The unique resource identifier of the database account.
	Id *string `json:"id,omitempty"`

	// Name: The name of the database account.
	Name *string `json:"name,omitempty"`

	// Properties: Properties related to the Role Assignment.
	Properties *SqlRoleAssignmentResource_STATUSARM `json:"properties,omitempty"`

	// Type: The type of Azure resource.
	Type *string `json:"type,omitempty"`
}

type SqlRoleAssignmentResource_STATUSARM struct {
	// PrincipalId: The unique identifier for the associated AAD principal in the AAD graph to which access is being granted
	// through this Role Assignment. Tenant ID for the principal is inferred using the tenant associated with the subscription.
	PrincipalId *string `json:"principalId,omitempty"`

	// RoleDefinitionId: The unique identifier for the associated Role Definition.
	RoleDefinitionId *string `json:"roleDefinitionId,omitempty"`

	// Scope: The data plane resource path for which access is being granted through this Role Assignment.
	Scope *string `json:"scope,omitempty"`
}
