// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=documentdb.azure.com,resources=sqlroleassignments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=documentdb.azure.com,resources={sqlroleassignments/status,sqlroleassignments/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20210515.SqlRoleAssignment
// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlRoleAssignments
type SqlRoleAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccounts_SqlRoleAssignment_Spec `json:"spec,omitempty"`
	Status            SqlRoleAssignmentGetResults_STATUS      `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlRoleAssignment{}

// GetConditions returns the conditions of the resource
func (assignment *SqlRoleAssignment) GetConditions() conditions.Conditions {
	return assignment.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (assignment *SqlRoleAssignment) SetConditions(conditions conditions.Conditions) {
	assignment.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &SqlRoleAssignment{}

// AzureName returns the Azure name of the resource
func (assignment *SqlRoleAssignment) AzureName() string {
	return assignment.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (assignment SqlRoleAssignment) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (assignment *SqlRoleAssignment) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (assignment *SqlRoleAssignment) GetSpec() genruntime.ConvertibleSpec {
	return &assignment.Spec
}

// GetStatus returns the status of this resource
func (assignment *SqlRoleAssignment) GetStatus() genruntime.ConvertibleStatus {
	return &assignment.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlRoleAssignments"
func (assignment *SqlRoleAssignment) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlRoleAssignments"
}

// NewEmptyStatus returns a new empty (blank) status
func (assignment *SqlRoleAssignment) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SqlRoleAssignmentGetResults_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (assignment *SqlRoleAssignment) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(assignment.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  assignment.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (assignment *SqlRoleAssignment) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SqlRoleAssignmentGetResults_STATUS); ok {
		assignment.Status = *st
		return nil
	}

	// Convert status to required version
	var st SqlRoleAssignmentGetResults_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	assignment.Status = st
	return nil
}

// Hub marks that this SqlRoleAssignment is the hub type for conversion
func (assignment *SqlRoleAssignment) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (assignment *SqlRoleAssignment) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: assignment.Spec.OriginalVersion,
		Kind:    "SqlRoleAssignment",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20210515.SqlRoleAssignment
// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlRoleAssignments
type SqlRoleAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlRoleAssignment `json:"items"`
}

// Storage version of v1beta20210515.DatabaseAccounts_SqlRoleAssignment_Spec
type DatabaseAccounts_SqlRoleAssignment_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string `json:"azureName,omitempty"`
	OriginalVersion string `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/DatabaseAccount resource
	Owner            *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"DatabaseAccount"`
	PrincipalId      *string                            `json:"principalId,omitempty"`
	PropertyBag      genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	RoleDefinitionId *string                            `json:"roleDefinitionId,omitempty"`
	Scope            *string                            `json:"scope,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccounts_SqlRoleAssignment_Spec{}

// ConvertSpecFrom populates our DatabaseAccounts_SqlRoleAssignment_Spec from the provided source
func (assignment *DatabaseAccounts_SqlRoleAssignment_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == assignment {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(assignment)
}

// ConvertSpecTo populates the provided destination from our DatabaseAccounts_SqlRoleAssignment_Spec
func (assignment *DatabaseAccounts_SqlRoleAssignment_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == assignment {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(assignment)
}

// Storage version of v1beta20210515.SqlRoleAssignmentGetResults_STATUS
type SqlRoleAssignmentGetResults_STATUS struct {
	Conditions       []conditions.Condition `json:"conditions,omitempty"`
	Id               *string                `json:"id,omitempty"`
	Name             *string                `json:"name,omitempty"`
	PrincipalId      *string                `json:"principalId,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RoleDefinitionId *string                `json:"roleDefinitionId,omitempty"`
	Scope            *string                `json:"scope,omitempty"`
	Type             *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SqlRoleAssignmentGetResults_STATUS{}

// ConvertStatusFrom populates our SqlRoleAssignmentGetResults_STATUS from the provided source
func (results *SqlRoleAssignmentGetResults_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == results {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(results)
}

// ConvertStatusTo populates the provided destination from our SqlRoleAssignmentGetResults_STATUS
func (results *SqlRoleAssignmentGetResults_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == results {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(results)
}

func init() {
	SchemeBuilder.Register(&SqlRoleAssignment{}, &SqlRoleAssignmentList{})
}
