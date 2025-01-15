package models

import (
)

// PropertyUsingDetailType represents the structure of the property_using_detail_type table
type PropertyUsingDetailType struct {
	UsingDetailTypeId int `gorm:"column:using_detail_type_id" json:"using_detail_type_id"`
	UsingDetailTypeName *string `gorm:"column:using_detail_type_name" json:"using_detail_type_name"`
}
