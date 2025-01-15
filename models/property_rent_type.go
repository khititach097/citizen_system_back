package models

import (
)

// PropertyRentType represents the structure of the property_rent_type table
type PropertyRentType struct {
	RentTypeId int `gorm:"column:rent_type_id" json:"rent_type_id"`
	RentTypeDetail *string `gorm:"column:rent_type_detail" json:"rent_type_detail"`
}
