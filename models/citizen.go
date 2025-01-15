package models

import (
	"github.com/google/uuid"
	"time"
)

// Citizen represents the structure of the citizen table
type Citizen struct {
	Id uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	PersonTypeId *int `gorm:"column:person_type_id" json:"person_type_id"`
	TaxId *string `gorm:"column:tax_id" json:"tax_id"`
	PrefixNameId *int `gorm:"column:prefix_name_id" json:"prefix_name_id"`
	FirstName *string `gorm:"column:first_name" json:"first_name"`
	LastName *string `gorm:"column:last_name" json:"last_name"`
	PhoneNumber *string `gorm:"column:phone_number" json:"phone_number"`
	LineId *string `gorm:"column:line_id" json:"line_id"`
	Email *string `gorm:"column:email" json:"email"`
	AddressHouseNumber *string `gorm:"column:address_house_number" json:"address_house_number"`
	AddressZone *string `gorm:"column:address_zone" json:"address_zone"`
	AddressStreet *string `gorm:"column:address_street" json:"address_street"`
	AddressAlleyway *string `gorm:"column:address_alleyway" json:"address_alleyway"`
	AddressProvinceCode *string `gorm:"column:address_province_code" json:"address_province_code"`
	AddressSubDistrictCode *string `gorm:"column:address_sub_district_code" json:"address_sub_district_code"`
	AddressDistrictCode *string `gorm:"column:address_district_code" json:"address_district_code"`
	AddressPostcode *string `gorm:"column:address_postcode" json:"address_postcode"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	CorporateName *string `gorm:"column:corporate_name" json:"corporate_name"`
	Codept4 *string `gorm:"column:codept4" json:"codept4"`
	FaxNo *string `gorm:"column:fax_no" json:"fax_no"`
	MuniCode *string `gorm:"column:muni_code" json:"muni_code"`
	DateSource *time.Time `gorm:"column:date_source" json:"date_source"`
	CitizenIdCheck *bool `gorm:"column:citizen_id_check" json:"citizen_id_check"`
	CodeName *string `gorm:"column:code_name" json:"code_name"`
	CurrentAddressHouseNumber *string `gorm:"column:current_address_house_number" json:"current_address_house_number"`
	CurrentAddressZone *string `gorm:"column:current_address_zone" json:"current_address_zone"`
	CurrentAddressStreet *string `gorm:"column:current_address_street" json:"current_address_street"`
	CurrentAddressAlleyway *string `gorm:"column:current_address_alleyway" json:"current_address_alleyway"`
	CurrentAddressProvinceCode *string `gorm:"column:current_address_province_code" json:"current_address_province_code"`
	CurrentAddressSubDistrictCode *string `gorm:"column:current_address_sub_district_code" json:"current_address_sub_district_code"`
	CurrentAddressDistrictCode *string `gorm:"column:current_address_district_code" json:"current_address_district_code"`
	CurrentAddressPostcode *string `gorm:"column:current_address_postcode" json:"current_address_postcode"`
	IsSameAddress *bool `gorm:"column:is_same_address" json:"is_same_address"`
	SyncId *uuid.UUID `gorm:"column:sync_id" json:"sync_id"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	DeletedBy *uuid.UUID `gorm:"column:deleted_by" json:"deleted_by"`
	CreatedBy *uuid.UUID `gorm:"column:created_by" json:"created_by"`
	UpdatedBy *uuid.UUID `gorm:"column:updated_by" json:"updated_by"`
	CitizenId *uuid.UUID `gorm:"column:citizen_id" json:"citizen_id"`
}
