package models

import (
	"github.com/google/uuid"
	"time"
)

// AsBuildingLand represents the structure of the as_building_land table
type AsBuildingLand struct {
	LandId *uuid.UUID `gorm:"column:land_id" json:"land_id"`
	BuildingId *uuid.UUID `gorm:"column:building_id" json:"building_id"`
	LandUsedId *uuid.UUID `gorm:"column:land_used_id" json:"land_used_id"`
	BuildingLandId uuid.UUID `gorm:"column:building_land_id;primaryKey" json:"building_land_id"`
	MuniCode *string `gorm:"column:muni_code" json:"muni_code"`
	SourceDt *time.Time `gorm:"column:source_dt" json:"source_dt"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	CreatedBy *uuid.UUID `gorm:"column:created_by" json:"created_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	UpdatedBy *uuid.UUID `gorm:"column:updated_by" json:"updated_by"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	DeletedBy *uuid.UUID `gorm:"column:deleted_by" json:"deleted_by"`
}
