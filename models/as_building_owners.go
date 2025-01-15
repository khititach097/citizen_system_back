package models

import (
	"time"

	"github.com/google/uuid"
)

// AsBuildingOwners represents the structure of the as_building_owners table
type AsBuildingOwners struct {
	Id                   uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	CreatedAt            *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	CreatedBy            *uuid.UUID `gorm:"column:created_by" json:"created_by"`
	UpdatedAt            *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	UpdatedBy            *uuid.UUID `gorm:"column:updated_by" json:"updated_by"`
	DeletedAt            *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	DeletedBy            *uuid.UUID `gorm:"column:deleted_by" json:"deleted_by"`
	BuildingId           *uuid.UUID `gorm:"column:building_id" json:"building_id"`
	OwnerId              *uuid.UUID `gorm:"column:owner_id" json:"owner_id"`
	OwnerLineNo          *string    `gorm:"column:owner_line_no" json:"owner_line_no"`
	BuildingOwnerId      *uuid.UUID `gorm:"column:building_owner_id" json:"building_owner_id"`
	MuniCode             *string    `gorm:"column:muni_code" json:"muni_code"`
	CutaxOwnerbuildingId *string    `gorm:"column:cutax_ownerbuilding_id" json:"cutax_ownerbuilding_id"`
}

func (AsBuildingOwners) TableName() string {
	return "as_building_owners"
}
