package models

import (
	"github.com/google/uuid"
)

// PropertyBuildingMainType represents the structure of the property_building_main_type table
type PropertyBuildingMainType struct {
	BuildingTypeId   *string   `gorm:"column:building_type_id" json:"building_type_id"`
	BuildingTypeName *string   `gorm:"column:building_type_name" json:"building_type_name"`
	BuildingRateM2   *string   `gorm:"column:building_rate_m2" json:"building_rate_m2"`
	RefSource        *string   `gorm:"column:ref_source" json:"ref_source"`
	RefCode          *string   `gorm:"column:ref_code" json:"ref_code"`
	Id               uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	BuildingTypeYear *string   `gorm:"column:building_type_year" json:"building_type_year"`
	MuniCode         *string   `gorm:"column:muni_code" json:"muni_code"`
	Annual           *string   `gorm:"column:annual" json:"annual"`
	SourceName       *string   `gorm:"column:source_name" json:"source_name"`
	SourceDt         *string   `gorm:"column:source_dt" json:"source_dt"`
}

func (PropertyBuildingMainType) TableName() string {
	return "property_building_main_type"
}
