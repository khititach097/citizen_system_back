package models

import (
)

// PropertySignboardDisplayType represents the structure of the property_signboard_display_type table
type PropertySignboardDisplayType struct {
	PropertySignboardDisplayType int `gorm:"column:property_signboard_display_type;primaryKey" json:"property_signboard_display_type"`
	PropertySignboardDisplayName *string `gorm:"column:property_signboard_display_name" json:"property_signboard_display_name"`
	PropertySignboardDisplayDesc *string `gorm:"column:property_signboard_display_desc" json:"property_signboard_display_desc"`
}
