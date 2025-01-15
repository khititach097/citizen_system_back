package models

import (
	"time"

	"github.com/google/uuid"
)

// MasterTaxRateSignboard represents the structure of the master_tax_rate_signboard table
type MasterTaxRateSignboard struct {
	Id                     uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	TaxYear                *string    `gorm:"column:tax_year" json:"tax_year"`
	SignboardTypeId        *int       `gorm:"column:signboard_type_id" json:"signboard_type_id"`
	SignboardDisplayTypeId *int       `gorm:"column:signboard_display_type_id" json:"signboard_display_type_id"`
	TaxRate                *string    `gorm:"column:tax_rate" json:"tax_rate"`
	TaxRateValue           *float64   `gorm:"column:tax_rate_value" json:"tax_rate_value"`
	Activate               *bool      `gorm:"column:activate" json:"activate"`
	CreatedAt              *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	CreatedBy              *uuid.UUID `gorm:"column:created_by" json:"created_by"`
	UpdatedAt              *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	UpdatedBy              *string    `gorm:"column:updated_by" json:"updated_by"`
	SignboardDesc          *string    `gorm:"column:signboard_desc" json:"signboard_desc"`
	TaxBase                *float64   `gorm:"column:tax_base" json:"tax_base"`
	TaxAreaUnit            *string    `gorm:"column:tax_area_unit" json:"tax_area_unit"`
}

func (MasterTaxRateSignboard) TableName() string {
	return "master_tax_rate_signboard"
}
