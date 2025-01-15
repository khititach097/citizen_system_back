package models

import (
)

// AddressProvince represents the structure of the address_province table
type AddressProvince struct {
	ProvinceId int `gorm:"column:province_id" json:"province_id"`
	ProvinceName *string `gorm:"column:province_name" json:"province_name"`
	Active *int `gorm:"column:active" json:"active"`
}
