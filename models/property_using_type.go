package models

import (
)

// PropertyUsingType represents the structure of the property_using_type table
type PropertyUsingType struct {
	UsingTypeId int `gorm:"column:using_type_id;primaryKey" json:"using_type_id"`
	UsingTypeDetail *string `gorm:"column:using_type_detail" json:"using_type_detail"`
	AllowSelect *int `gorm:"column:allow_select" json:"allow_select"`
}
