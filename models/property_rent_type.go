package models

// PropertyRentType represents the structure of the property_rent_type table
type PropertyRentType struct {
	RentTypeId     int     `gorm:"column:rent_type_id" json:"rent_type_id"`
	RentTypeDetail *string `gorm:"column:rent_type_detail" json:"rent_type_detail"`
}

func (PropertyRentType) TableName() string {
	return "property_rent_type"
}
