package models

// PropertyHouseholdType represents the structure of the property_household_type table
type PropertyHouseholdType struct {
	HouseholdTypeId   int     `gorm:"column:household_type_id;primaryKey" json:"household_type_id"`
	HouseholdTypeName *string `gorm:"column:household_type_name" json:"household_type_name"`
}

func (PropertyHouseholdType) TableName() string {
	return "property_household_type"
}
