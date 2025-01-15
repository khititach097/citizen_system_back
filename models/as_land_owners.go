package models

import (
	"time"

	"github.com/google/uuid"
)

// AsLandOwners represents the structure of the as_land_owners table
type AsLandOwners struct {
	Id               uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	CreatedAt        *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	CreatedBy        *uuid.UUID `gorm:"column:created_by" json:"created_by"`
	UpdatedAt        *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	UpdatedBy        *uuid.UUID `gorm:"column:updated_by" json:"updated_by"`
	DeletedAt        *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	DeletedBy        *uuid.UUID `gorm:"column:deleted_by" json:"deleted_by"`
	LandId           *uuid.UUID `gorm:"column:land_id" json:"land_id"`
	OwnerId          *uuid.UUID `gorm:"column:owner_id" json:"owner_id"`
	OwnerLineNo      *string    `gorm:"column:owner_line_no" json:"owner_line_no"`
	MuniCode         *string    `gorm:"column:muni_code" json:"muni_code"`
	DateSource       *time.Time `gorm:"column:date_source" json:"date_source"`
	CutaxOwnerlandId *string    `gorm:"column:cutax_ownerland_id" json:"cutax_ownerland_id"`
}

func (AsLandOwners) TableName() string {
	return "as_land_owners"
}
