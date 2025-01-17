// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601

// Deprecated version of Server_STATUS. Use v1beta20210601.Server_STATUS instead
type Server_STATUS_ARM struct {
	Id         *string                      `json:"id,omitempty"`
	Location   *string                      `json:"location,omitempty"`
	Name       *string                      `json:"name,omitempty"`
	Properties *ServerProperties_STATUS_ARM `json:"properties,omitempty"`
	Sku        *Sku_STATUS_ARM              `json:"sku,omitempty"`
	SystemData *SystemData_STATUS_ARM       `json:"systemData,omitempty"`
	Tags       map[string]string            `json:"tags,omitempty"`
	Type       *string                      `json:"type,omitempty"`
}

// Deprecated version of ServerProperties_STATUS. Use v1beta20210601.ServerProperties_STATUS instead
type ServerProperties_STATUS_ARM struct {
	AdministratorLogin       *string                             `json:"administratorLogin,omitempty"`
	AvailabilityZone         *string                             `json:"availabilityZone,omitempty"`
	Backup                   *Backup_STATUS_ARM                  `json:"backup,omitempty"`
	CreateMode               *ServerProperties_CreateMode_STATUS `json:"createMode,omitempty"`
	FullyQualifiedDomainName *string                             `json:"fullyQualifiedDomainName,omitempty"`
	HighAvailability         *HighAvailability_STATUS_ARM        `json:"highAvailability,omitempty"`
	MaintenanceWindow        *MaintenanceWindow_STATUS_ARM       `json:"maintenanceWindow,omitempty"`
	MinorVersion             *string                             `json:"minorVersion,omitempty"`
	Network                  *Network_STATUS_ARM                 `json:"network,omitempty"`
	PointInTimeUTC           *string                             `json:"pointInTimeUTC,omitempty"`
	SourceServerResourceId   *string                             `json:"sourceServerResourceId,omitempty"`
	State                    *ServerProperties_State_STATUS      `json:"state,omitempty"`
	Storage                  *Storage_STATUS_ARM                 `json:"storage,omitempty"`
	Version                  *ServerVersion_STATUS               `json:"version,omitempty"`
}

// Deprecated version of Sku_STATUS. Use v1beta20210601.Sku_STATUS instead
type Sku_STATUS_ARM struct {
	Name *string          `json:"name,omitempty"`
	Tier *Sku_Tier_STATUS `json:"tier,omitempty"`
}

// Deprecated version of Backup_STATUS. Use v1beta20210601.Backup_STATUS instead
type Backup_STATUS_ARM struct {
	BackupRetentionDays *int                              `json:"backupRetentionDays,omitempty"`
	EarliestRestoreDate *string                           `json:"earliestRestoreDate,omitempty"`
	GeoRedundantBackup  *Backup_GeoRedundantBackup_STATUS `json:"geoRedundantBackup,omitempty"`
}

// Deprecated version of HighAvailability_STATUS. Use v1beta20210601.HighAvailability_STATUS instead
type HighAvailability_STATUS_ARM struct {
	Mode                    *HighAvailability_Mode_STATUS  `json:"mode,omitempty"`
	StandbyAvailabilityZone *string                        `json:"standbyAvailabilityZone,omitempty"`
	State                   *HighAvailability_State_STATUS `json:"state,omitempty"`
}

// Deprecated version of MaintenanceWindow_STATUS. Use v1beta20210601.MaintenanceWindow_STATUS instead
type MaintenanceWindow_STATUS_ARM struct {
	CustomWindow *string `json:"customWindow,omitempty"`
	DayOfWeek    *int    `json:"dayOfWeek,omitempty"`
	StartHour    *int    `json:"startHour,omitempty"`
	StartMinute  *int    `json:"startMinute,omitempty"`
}

// Deprecated version of Network_STATUS. Use v1beta20210601.Network_STATUS instead
type Network_STATUS_ARM struct {
	DelegatedSubnetResourceId   *string                             `json:"delegatedSubnetResourceId,omitempty"`
	PrivateDnsZoneArmResourceId *string                             `json:"privateDnsZoneArmResourceId,omitempty"`
	PublicNetworkAccess         *Network_PublicNetworkAccess_STATUS `json:"publicNetworkAccess,omitempty"`
}

// Deprecated version of Sku_Tier_STATUS. Use v1beta20210601.Sku_Tier_STATUS instead
type Sku_Tier_STATUS string

const (
	Sku_Tier_STATUS_Burstable       = Sku_Tier_STATUS("Burstable")
	Sku_Tier_STATUS_GeneralPurpose  = Sku_Tier_STATUS("GeneralPurpose")
	Sku_Tier_STATUS_MemoryOptimized = Sku_Tier_STATUS("MemoryOptimized")
)

// Deprecated version of Storage_STATUS. Use v1beta20210601.Storage_STATUS instead
type Storage_STATUS_ARM struct {
	StorageSizeGB *int `json:"storageSizeGB,omitempty"`
}
