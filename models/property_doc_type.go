package models

import (
)

// PropertyDocType represents the structure of the property_doc_type table
type PropertyDocType struct {
	DocTypeId int `gorm:"column:doc_type_id" json:"doc_type_id"`
	DocTypeName *string `gorm:"column:doc_type_name" json:"doc_type_name"`
}
