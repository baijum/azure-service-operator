// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	alpha20200601 "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1alpha1api20200601"
	alpha20200601s "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1alpha1api20200601storage"
	v20200601 "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1beta20200601"
	v20200601s "github.com/Azure/azure-service-operator/v2/api/eventgrid/v1beta20200601storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type DomainsTopicExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *DomainsTopicExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&alpha20200601.DomainsTopic{},
		&alpha20200601s.DomainsTopic{},
		&v20200601.DomainsTopic{},
		&v20200601s.DomainsTopic{}}
}
