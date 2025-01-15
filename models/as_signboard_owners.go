package models

import (
	"github.com/google/uuid"
	"time"
)

// AsSignboardOwners represents the structure of the as_signboard_owners table
type AsSignboardOwners struct {
	Id uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	CreatedBy *uuid.UUID `gorm:"column:created_by" json:"created_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	UpdatedBy *uuid.UUID `gorm:"column:updated_by" json:"updated_by"`
	SignboardId *uuid.UUID `gorm:"column:signboard_id" json:"signboard_id"`
	OwnerId *uuid.UUID `gorm:"column:owner_id" json:"owner_id"`
	OwnerLineNo *string `gorm:"column:owner_line_no" json:"owner_line_no"`
	MuniCode *string `gorm:"column:muni_code" json:"muni_code"`
	CutaxOwnerlabelId *string `gorm:"column:cutax_ownerlabel_id" json:"cutax_ownerlabel_id"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	DeletedBy *uuid.UUID `gorm:"column:deleted_by" json:"deleted_by"`
	MainOwner *string `gorm:"column:main_owner" json:"main_owner"`
	SignboardOwnerCode *string `gorm:"column:signboard_owner_code" json:"signboard_owner_code"`
	TaxYear *string `gorm:"column:tax_year" json:"tax_year"`
	SourceDt *string `gorm:"column:source_dt" json:"source_dt"`
}
