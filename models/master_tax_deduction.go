package models

import (
	"github.com/google/uuid"
	"time"
)

// MasterTaxDeduction represents the structure of the master_tax_deduction table
type MasterTaxDeduction struct {
	Id uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Description *string `gorm:"column:description" json:"description"`
	Depreciate *float64 `gorm:"column:depreciate" json:"depreciate"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	CreatedBy *uuid.UUID `gorm:"column:created_by" json:"created_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	UpdatedBy *uuid.UUID `gorm:"column:updated_by" json:"updated_by"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	DeletedBy *uuid.UUID `gorm:"column:deleted_by" json:"deleted_by"`
	SourceDt *string `gorm:"column:source_dt" json:"source_dt"`
	Note *string `gorm:"column:note" json:"note"`
}
