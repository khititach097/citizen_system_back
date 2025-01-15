package models

import (
	"github.com/google/uuid"
	"time"
)

// AsAdjoiningLands represents the structure of the as_adjoining_lands table
type AsAdjoiningLands struct {
	Id uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	LandId *uuid.UUID `gorm:"column:land_id" json:"land_id"`
	AdjoiningLandId *uuid.UUID `gorm:"column:adjoining_land_id" json:"adjoining_land_id"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	CreatedBy *uuid.UUID `gorm:"column:created_by" json:"created_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	UpdatedBy *uuid.UUID `gorm:"column:updated_by" json:"updated_by"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	DeletedBy *uuid.UUID `gorm:"column:deleted_by" json:"deleted_by"`
	MuniCode *string `gorm:"column:muni_code" json:"muni_code"`
	TaxYear *string `gorm:"column:tax_year" json:"tax_year"`
	SourceDt *time.Time `gorm:"column:source_dt" json:"source_dt"`
	Status *string `gorm:"column:status" json:"status"`
}
