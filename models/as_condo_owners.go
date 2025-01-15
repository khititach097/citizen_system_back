package models

import (
	"github.com/google/uuid"
	"time"
)

// AsCondoOwners represents the structure of the as_condo_owners table
type AsCondoOwners struct {
	Id uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	CondoRoomId *uuid.UUID `gorm:"column:condo_room_id" json:"condo_room_id"`
	CreatedBy *uuid.UUID `gorm:"column:created_by" json:"created_by"`
	OwnerId *uuid.UUID `gorm:"column:owner_id" json:"owner_id"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	MainOwner *string `gorm:"column:main_owner" json:"main_owner"`
	UpdatedBy *uuid.UUID `gorm:"column:updated_by" json:"updated_by"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	DeletedBy *uuid.UUID `gorm:"column:deleted_by" json:"deleted_by"`
	CondoId *uuid.UUID `gorm:"column:condo_id" json:"condo_id"`
	OwnerLineNo *string `gorm:"column:owner_line_no" json:"owner_line_no"`
	MuniCode *string `gorm:"column:muni_code" json:"muni_code"`
	TaxYear *string `gorm:"column:tax_year" json:"tax_year"`
	CondoOwnerCode *string `gorm:"column:condo_owner_code" json:"condo_owner_code"`
	SourceDt *string `gorm:"column:source_dt" json:"source_dt"`
}
