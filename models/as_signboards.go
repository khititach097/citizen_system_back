package models

import (
	"github.com/google/uuid"
	"time"
)

// AsSignboards represents the structure of the as_signboards table
type AsSignboards struct {
	Id uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	CreatedBy *uuid.UUID `gorm:"column:created_by" json:"created_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	UpdatedBy *uuid.UUID `gorm:"column:updated_by" json:"updated_by"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	DeletedBy *uuid.UUID `gorm:"column:deleted_by" json:"deleted_by"`
	SignboardWidth *float64 `gorm:"column:signboard_width" json:"signboard_width"`
	SignboardHeight *float64 `gorm:"column:signboard_height" json:"signboard_height"`
	SignboardSide *int `gorm:"column:signboard_side" json:"signboard_side"`
	SignboardUnitTotal *int `gorm:"column:signboard_unit_total" json:"signboard_unit_total"`
	PropertySignboardTypeId *int `gorm:"column:property_signboard_type_id" json:"property_signboard_type_id"`
	SignboardText *string `gorm:"column:signboard_text" json:"signboard_text"`
	Note *string `gorm:"column:note" json:"note"`
	LandId *uuid.UUID `gorm:"column:land_id" json:"land_id"`
	MapGeometry *string `gorm:"column:map_geometry" json:"map_geometry"`
	MapLat *string `gorm:"column:map_lat" json:"map_lat"`
	MapLong *string `gorm:"column:map_long" json:"map_long"`
	UtmMap1 *string `gorm:"column:utm_map1" json:"utm_map1"`
	UtmMap2 *string `gorm:"column:utm_map2" json:"utm_map2"`
	UtmMap3 *string `gorm:"column:utm_map3" json:"utm_map3"`
	UtmMap4 *string `gorm:"column:utm_map4" json:"utm_map4"`
	UtmScale *string `gorm:"column:utm_scale" json:"utm_scale"`
	TaxYear *string `gorm:"column:tax_year" json:"tax_year"`
	SignboardSetupDate *time.Time `gorm:"column:signboard_setup_date" json:"signboard_setup_date"`
	ParcelCode *string `gorm:"column:parcel_code" json:"parcel_code"`
	AddressPlaceNumber *string `gorm:"column:address_place_number" json:"address_place_number"`
	AddressVillageName *string `gorm:"column:address_village_name" json:"address_village_name"`
	AddressZone *string `gorm:"column:address_zone" json:"address_zone"`
	AddressAlleyWay *string `gorm:"column:address_alley_way" json:"address_alley_way"`
	AddressStreet *string `gorm:"column:address_street" json:"address_street"`
	AddressSubDistrictId *string `gorm:"column:address_sub_district_id" json:"address_sub_district_id"`
	SignboardTotalArea *string `gorm:"column:signboard_total_area" json:"signboard_total_area"`
	Stu *string `gorm:"column:stu" json:"stu"`
	MuniCode *string `gorm:"column:muni_code" json:"muni_code"`
	SourceDt *time.Time `gorm:"column:source_dt" json:"source_dt"`
	CutaxLabelId *string `gorm:"column:cutax_label_id" json:"cutax_label_id"`
	SyncId *uuid.UUID `gorm:"column:sync_id" json:"sync_id"`
	PropertySignboardDisplayTypeId *int `gorm:"column:property_signboard_display_type_id" json:"property_signboard_display_type_id"`
	AddressProvinceId *int `gorm:"column:address_province_id" json:"address_province_id"`
	AddressDistrictId *int `gorm:"column:address_district_id" json:"address_district_id"`
	AddressPostcode *string `gorm:"column:address_postcode" json:"address_postcode"`
	SignboardCode *string `gorm:"column:signboard_code" json:"signboard_code"`
	SignboardName *string `gorm:"column:signboard_name" json:"signboard_name"`
	CancelStatus *int `gorm:"column:cancel_status" json:"cancel_status"`
	CancelDate *time.Time `gorm:"column:cancel_date" json:"cancel_date"`
	CancelBy *uuid.UUID `gorm:"column:cancel_by" json:"cancel_by"`
	CancelAt *time.Time `gorm:"column:cancel_at" json:"cancel_at"`
}
