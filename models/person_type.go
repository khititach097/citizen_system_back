package models

// PersonType represents the structure of the person_type table
type PersonType struct {
	PersonTypeId       int     `gorm:"column:person_type_id" json:"person_type_id"`
	PersonTypeName     *string `gorm:"column:person_type_name" json:"person_type_name"`
	PersonTypeDropdown string  `gorm:"column:person_type_dropdown" json:"person_type_dropdown"`
}

func (PersonType) TableName() string {
	return "person_type"
}
