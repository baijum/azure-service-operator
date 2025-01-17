// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210501

import "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

// Deprecated version of Server_STATUS. Use v1beta20210501.Server_STATUS instead
type Server_STATUS_ARM struct {
	Id         *string                      `json:"id,omitempty"`
	Identity   *Identity_STATUS_ARM         `json:"identity,omitempty"`
	Location   *string                      `json:"location,omitempty"`
	Name       *string                      `json:"name,omitempty"`
	Properties *ServerProperties_STATUS_ARM `json:"properties,omitempty"`
	Sku        *Sku_STATUS_ARM              `json:"sku,omitempty"`
	SystemData *SystemData_STATUS_ARM       `json:"systemData,omitempty"`
	Tags       map[string]string            `json:"tags,omitempty"`
	Type       *string                      `json:"type,omitempty"`
}

// Deprecated version of Identity_STATUS. Use v1beta20210501.Identity_STATUS instead
type Identity_STATUS_ARM struct {
	PrincipalId            *string               `json:"principalId,omitempty"`
	TenantId               *string               `json:"tenantId,omitempty"`
	Type                   *Identity_Type_STATUS `json:"type,omitempty"`
	UserAssignedIdentities map[string]v1.JSON    `json:"userAssignedIdentities,omitempty"`
}

// Deprecated version of ServerProperties_STATUS. Use v1beta20210501.ServerProperties_STATUS instead
type ServerProperties_STATUS_ARM struct {
	AdministratorLogin       *string                             `json:"administratorLogin,omitempty"`
	AvailabilityZone         *string                             `json:"availabilityZone,omitempty"`
	Backup                   *Backup_STATUS_ARM                  `json:"backup,omitempty"`
	CreateMode               *ServerProperties_CreateMode_STATUS `json:"createMode,omitempty"`
	DataEncryption           *DataEncryption_STATUS_ARM          `json:"dataEncryption,omitempty"`
	FullyQualifiedDomainName *string                             `json:"fullyQualifiedDomainName,omitempty"`
	HighAvailability         *HighAvailability_STATUS_ARM        `json:"highAvailability,omitempty"`
	MaintenanceWindow        *MaintenanceWindow_STATUS_ARM       `json:"maintenanceWindow,omitempty"`
	Network                  *Network_STATUS_ARM                 `json:"network,omitempty"`
	ReplicaCapacity          *int                                `json:"replicaCapacity,omitempty"`
	ReplicationRole          *ReplicationRole_STATUS             `json:"replicationRole,omitempty"`
	RestorePointInTime       *string                             `json:"restorePointInTime,omitempty"`
	SourceServerResourceId   *string                             `json:"sourceServerResourceId,omitempty"`
	State                    *ServerProperties_State_STATUS      `json:"state,omitempty"`
	Storage                  *Storage_STATUS_ARM                 `json:"storage,omitempty"`
	Version                  *ServerVersion_STATUS               `json:"version,omitempty"`
}

// Deprecated version of Sku_STATUS. Use v1beta20210501.Sku_STATUS instead
type Sku_STATUS_ARM struct {
	Name *string          `json:"name,omitempty"`
	Tier *Sku_Tier_STATUS `json:"tier,omitempty"`
}

// Deprecated version of Backup_STATUS. Use v1beta20210501.Backup_STATUS instead
type Backup_STATUS_ARM struct {
	BackupRetentionDays *int                     `json:"backupRetentionDays,omitempty"`
	EarliestRestoreDate *string                  `json:"earliestRestoreDate,omitempty"`
	GeoRedundantBackup  *EnableStatusEnum_STATUS `json:"geoRedundantBackup,omitempty"`
}

// Deprecated version of DataEncryption_STATUS. Use v1beta20210501.DataEncryption_STATUS instead
type DataEncryption_STATUS_ARM struct {
	GeoBackupKeyURI                 *string                     `json:"geoBackupKeyURI,omitempty"`
	GeoBackupUserAssignedIdentityId *string                     `json:"geoBackupUserAssignedIdentityId,omitempty"`
	PrimaryKeyURI                   *string                     `json:"primaryKeyURI,omitempty"`
	PrimaryUserAssignedIdentityId   *string                     `json:"primaryUserAssignedIdentityId,omitempty"`
	Type                            *DataEncryption_Type_STATUS `json:"type,omitempty"`
}

// Deprecated version of HighAvailability_STATUS. Use v1beta20210501.HighAvailability_STATUS instead
type HighAvailability_STATUS_ARM struct {
	Mode                    *HighAvailability_Mode_STATUS  `json:"mode,omitempty"`
	StandbyAvailabilityZone *string                        `json:"standbyAvailabilityZone,omitempty"`
	State                   *HighAvailability_State_STATUS `json:"state,omitempty"`
}

// Deprecated version of Identity_Type_STATUS. Use v1beta20210501.Identity_Type_STATUS instead
type Identity_Type_STATUS string

const Identity_Type_STATUS_UserAssigned = Identity_Type_STATUS("UserAssigned")

// Deprecated version of MaintenanceWindow_STATUS. Use v1beta20210501.MaintenanceWindow_STATUS instead
type MaintenanceWindow_STATUS_ARM struct {
	CustomWindow *string `json:"customWindow,omitempty"`
	DayOfWeek    *int    `json:"dayOfWeek,omitempty"`
	StartHour    *int    `json:"startHour,omitempty"`
	StartMinute  *int    `json:"startMinute,omitempty"`
}

// Deprecated version of Network_STATUS. Use v1beta20210501.Network_STATUS instead
type Network_STATUS_ARM struct {
	DelegatedSubnetResourceId *string                  `json:"delegatedSubnetResourceId,omitempty"`
	PrivateDnsZoneResourceId  *string                  `json:"privateDnsZoneResourceId,omitempty"`
	PublicNetworkAccess       *EnableStatusEnum_STATUS `json:"publicNetworkAccess,omitempty"`
}

// Deprecated version of Sku_Tier_STATUS. Use v1beta20210501.Sku_Tier_STATUS instead
type Sku_Tier_STATUS string

const (
	Sku_Tier_STATUS_Burstable       = Sku_Tier_STATUS("Burstable")
	Sku_Tier_STATUS_GeneralPurpose  = Sku_Tier_STATUS("GeneralPurpose")
	Sku_Tier_STATUS_MemoryOptimized = Sku_Tier_STATUS("MemoryOptimized")
)

// Deprecated version of Storage_STATUS. Use v1beta20210501.Storage_STATUS instead
type Storage_STATUS_ARM struct {
	AutoGrow      *EnableStatusEnum_STATUS `json:"autoGrow,omitempty"`
	Iops          *int                     `json:"iops,omitempty"`
	StorageSizeGB *int                     `json:"storageSizeGB,omitempty"`
	StorageSku    *string                  `json:"storageSku,omitempty"`
}
