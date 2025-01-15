package models

import (
	"github.com/google/uuid"
	"time"
)

// AsBuildingUsed represents the structure of the as_building_used table
type AsBuildingUsed struct {
	Id uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	CreatedBy *uuid.UUID `gorm:"column:created_by" json:"created_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	UpdatedBy *uuid.UUID `gorm:"column:updated_by" json:"updated_by"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	DeletedBy *uuid.UUID `gorm:"column:deleted_by" json:"deleted_by"`
	BuildingId *uuid.UUID `gorm:"column:building_id" json:"building_id"`
	PropertyRentTypeId *int `gorm:"column:property_rent_type_id" json:"property_rent_type_id"`
	BuildingHouseholdTypeId *int `gorm:"column:building_household_type_id" json:"building_household_type_id"`
	BuildingCultivationSpaceSquareMeter *float64 `gorm:"column:building_cultivation_space_square_meter" json:"building_cultivation_space_square_meter"`
	BuildingSelfUsingSquareMeter *float64 `gorm:"column:building_self_using_square_meter" json:"building_self_using_square_meter"`
	BuildingForRentSquareMeter *float64 `gorm:"column:building_for_rent_square_meter" json:"building_for_rent_square_meter"`
	BuildingEmptySpaceSquareMeter *float64 `gorm:"column:building_empty_space_square_meter" json:"building_empty_space_square_meter"`
	BuildingEtcSquareMeter *float64 `gorm:"column:building_etc_square_meter" json:"building_etc_square_meter"`
	BuildingUsingTypeId *int `gorm:"column:building_using_type_id" json:"building_using_type_id"`
	UsedType *string `gorm:"column:used_type" json:"used_type"`
	FullAreaCheck *bool `gorm:"column:full_area_check" json:"full_area_check"`
	FloorNo *string `gorm:"column:floor_no" json:"floor_no"`
	AreaWidth *string `gorm:"column:area_width" json:"area_width"`
	AreaLength *string `gorm:"column:area_length" json:"area_length"`
	TaxYear *string `gorm:"column:tax_year" json:"tax_year"`
	Notes *string `gorm:"column:notes" json:"notes"`
	MuniCode *string `gorm:"column:muni_code" json:"muni_code"`
	SourceDt *time.Time `gorm:"column:source_dt" json:"source_dt"`
	NonSpecific *float64 `gorm:"column:non_specific" json:"non_specific"`
}
