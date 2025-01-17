// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

// Deprecated version of MongoDBCollectionGetResults_STATUS. Use v1beta20210515.MongoDBCollectionGetResults_STATUS instead
type MongoDBCollectionGetResults_STATUS_ARM struct {
	Id         *string                                    `json:"id,omitempty"`
	Location   *string                                    `json:"location,omitempty"`
	Name       *string                                    `json:"name,omitempty"`
	Properties *MongoDBCollectionGetProperties_STATUS_ARM `json:"properties,omitempty"`
	Tags       map[string]string                          `json:"tags,omitempty"`
	Type       *string                                    `json:"type,omitempty"`
}

// Deprecated version of MongoDBCollectionGetProperties_STATUS. Use v1beta20210515.MongoDBCollectionGetProperties_STATUS instead
type MongoDBCollectionGetProperties_STATUS_ARM struct {
	Options  *OptionsResource_STATUS_ARM                         `json:"options,omitempty"`
	Resource *MongoDBCollectionGetProperties_Resource_STATUS_ARM `json:"resource,omitempty"`
}

// Deprecated version of MongoDBCollectionGetProperties_Resource_STATUS. Use v1beta20210515.MongoDBCollectionGetProperties_Resource_STATUS instead
type MongoDBCollectionGetProperties_Resource_STATUS_ARM struct {
	AnalyticalStorageTtl *int                    `json:"analyticalStorageTtl,omitempty"`
	Etag                 *string                 `json:"_etag,omitempty"`
	Id                   *string                 `json:"id,omitempty"`
	Indexes              []MongoIndex_STATUS_ARM `json:"indexes,omitempty"`
	Rid                  *string                 `json:"_rid,omitempty"`
	ShardKey             map[string]string       `json:"shardKey,omitempty"`
	Ts                   *float64                `json:"_ts,omitempty"`
}

// Deprecated version of OptionsResource_STATUS. Use v1beta20210515.OptionsResource_STATUS instead
type OptionsResource_STATUS_ARM struct {
	AutoscaleSettings *AutoscaleSettings_STATUS_ARM `json:"autoscaleSettings,omitempty"`
	Throughput        *int                          `json:"throughput,omitempty"`
}

// Deprecated version of AutoscaleSettings_STATUS. Use v1beta20210515.AutoscaleSettings_STATUS instead
type AutoscaleSettings_STATUS_ARM struct {
	MaxThroughput *int `json:"maxThroughput,omitempty"`
}

// Deprecated version of MongoIndex_STATUS. Use v1beta20210515.MongoIndex_STATUS instead
type MongoIndex_STATUS_ARM struct {
	Key     *MongoIndexKeys_STATUS_ARM    `json:"key,omitempty"`
	Options *MongoIndexOptions_STATUS_ARM `json:"options,omitempty"`
}

// Deprecated version of MongoIndexKeys_STATUS. Use v1beta20210515.MongoIndexKeys_STATUS instead
type MongoIndexKeys_STATUS_ARM struct {
	Keys []string `json:"keys,omitempty"`
}

// Deprecated version of MongoIndexOptions_STATUS. Use v1beta20210515.MongoIndexOptions_STATUS instead
type MongoIndexOptions_STATUS_ARM struct {
	ExpireAfterSeconds *int  `json:"expireAfterSeconds,omitempty"`
	Unique             *bool `json:"unique,omitempty"`
}
