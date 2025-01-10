package models

import (
	// "gorm.io/gorm"
)

// Asset represents the asset table in the database
type Asset struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (Asset) TableName() string {
	return "assets"
}
