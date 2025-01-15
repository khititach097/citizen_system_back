package models

import (
	"github.com/google/uuid"
)

// AddressPostcode represents the structure of the address_postcode table
type AddressPostcode struct {
	Id           uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	PrefectureId *int      `gorm:"column:prefecture_id" json:"prefecture_id"`
	Postcode     *int      `gorm:"column:postcode" json:"postcode"`
	Active       *int      `gorm:"column:active" json:"active"`
}

func (AddressPostcode) TableName() string {
	return "address_postcode"
}
