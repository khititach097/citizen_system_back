package models

import (
)

// PropertySignboardType represents the structure of the property_signboard_type table
type PropertySignboardType struct {
	SignboardTypeId int `gorm:"column:signboard_type_id" json:"signboard_type_id"`
	SignboardTypeName *string `gorm:"column:signboard_type_name" json:"signboard_type_name"`
}
