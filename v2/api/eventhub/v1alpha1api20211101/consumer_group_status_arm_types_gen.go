// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

// Deprecated version of ConsumerGroup_STATUS. Use v1beta20211101.ConsumerGroup_STATUS instead
type ConsumerGroup_STATUS_ARM struct {
	Id         *string                              `json:"id,omitempty"`
	Location   *string                              `json:"location,omitempty"`
	Name       *string                              `json:"name,omitempty"`
	Properties *ConsumerGroup_Properties_STATUS_ARM `json:"properties,omitempty"`
	SystemData *SystemData_STATUS_ARM               `json:"systemData,omitempty"`
	Type       *string                              `json:"type,omitempty"`
}

// Deprecated version of ConsumerGroup_Properties_STATUS. Use v1beta20211101.ConsumerGroup_Properties_STATUS instead
type ConsumerGroup_Properties_STATUS_ARM struct {
	CreatedAt    *string `json:"createdAt,omitempty"`
	UpdatedAt    *string `json:"updatedAt,omitempty"`
	UserMetadata *string `json:"userMetadata,omitempty"`
}
