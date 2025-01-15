package models

// PropertyType represents the structure of the property_type table
type PropertyType struct {
	PropertyTypeId   int     `gorm:"column:property_type_id" json:"property_type_id"`
	PropertyTypeName *string `gorm:"column:property_type_name" json:"property_type_name"`
}

func (PropertyType) TableName() string {
	return "property_type"
}
