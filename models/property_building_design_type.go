package models

import (
)

// PropertyBuildingDesignType represents the structure of the property_building_design_type table
type PropertyBuildingDesignType struct {
	BuildingDesignTypeId int `gorm:"column:building_design_type_id" json:"building_design_type_id"`
	BuildingDesignTypeName *string `gorm:"column:building_design_type_name" json:"building_design_type_name"`
}
