package models

import (
	"github.com/google/uuid"
)

// PropertyBuildingSubType represents the structure of the property_building_sub_type table
type PropertyBuildingSubType struct {
	BuildingSubId *string `gorm:"column:building_sub_id" json:"building_sub_id"`
	BuildingMainId *string `gorm:"column:building_main_id" json:"building_main_id"`
	BuildingTypeName *string `gorm:"column:building_type_name" json:"building_type_name"`
	Id uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	BuildingTypeYear *string `gorm:"column:building_type_year" json:"building_type_year"`
	MuniCode *string `gorm:"column:muni_code" json:"muni_code"`
	Annual *string `gorm:"column:annual" json:"annual"`
	SourceName *string `gorm:"column:source_name" json:"source_name"`
	SourceDt *string `gorm:"column:source_dt" json:"source_dt"`
}
