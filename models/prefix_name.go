package models

import (
	"github.com/google/uuid"
)

// PrefixName represents the structure of the prefix_name table
type PrefixName struct {
	Id uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Name *string `gorm:"column:name" json:"name"`
	PersonType *int `gorm:"column:person_type" json:"person_type"`
	Active *string `gorm:"column:active" json:"active"`
}
