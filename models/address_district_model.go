package models

import (
	"github.com/google/uuid"
)

// AddressDistrict represents the structure of the address_district table
type AddressDistrict struct {
	Id             uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	ProvinceId     *int      `gorm:"column:province_id" json:"province_id"`
	PrefectureId   *int      `gorm:"column:prefecture_id" json:"prefecture_id"`
	PrefectureName *string   `gorm:"column:prefecture_name" json:"prefecture_name"`
	Active         *int      `gorm:"column:active" json:"active"`
}

func (AddressDistrict) TableName() string {
	return "address_district"
}
