package models

import (
	"github.com/google/uuid"
)

// AddressSubDistrict represents the structure of the address_sub_district table
type AddressSubDistrict struct {
	Id           uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	ProvinceId   *int      `gorm:"column:province_id" json:"province_id"`
	PrefectureId *int      `gorm:"column:prefecture_id" json:"prefecture_id"`
	DistrictId   *int      `gorm:"column:district_id" json:"district_id"`
	DistrictName *string   `gorm:"column:district_name" json:"district_name"`
	Active       *int      `gorm:"column:active" json:"active"`
}

func (AddressSubDistrict) TableName() string {
	return "address_sub_district"
}
