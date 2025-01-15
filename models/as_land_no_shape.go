package models

import (
	"github.com/google/uuid"
	"time"
)

// AsLandNoShape represents the structure of the as_land_no_shape table
type AsLandNoShape struct {
	LandId uuid.UUID `gorm:"column:land_id;primaryKey" json:"land_id"`
	Annual *string `gorm:"column:annual" json:"annual"`
	LandType *int `gorm:"column:land_type" json:"land_type"`
	LandNo *string `gorm:"column:land_no" json:"land_no"`
	SurveyNo *string `gorm:"column:survey_no" json:"survey_no"`
	Utm1 *string `gorm:"column:utm1" json:"utm1"`
	Utm2 *string `gorm:"column:utm2" json:"utm2"`
	Utm3 *string `gorm:"column:utm3" json:"utm3"`
	Utm4 *string `gorm:"column:utm4" json:"utm4"`
	UtmScale *string `gorm:"column:utm_scale" json:"utm_scale"`
	Moo *string `gorm:"column:moo" json:"moo"`
	Village *string `gorm:"column:village" json:"village"`
	Soi *string `gorm:"column:soi" json:"soi"`
	RoadId *int `gorm:"column:road_id" json:"road_id"`
	RoadName *string `gorm:"column:road_name" json:"road_name"`
	SubdistrictId *int `gorm:"column:subdistrict_id" json:"subdistrict_id"`
	DistrictId *int `gorm:"column:district_id" json:"district_id"`
	ProvinceId *int `gorm:"column:province_id" json:"province_id"`
	RegionId *int `gorm:"column:region_id" json:"region_id"`
	Rai *float64 `gorm:"column:rai" json:"rai"`
	Ngan *float64 `gorm:"column:ngan" json:"ngan"`
	Wa *float64 `gorm:"column:wa" json:"wa"`
	PricePerWa *float64 `gorm:"column:price_per_wa" json:"price_per_wa"`
	PriceEstimatePerWa *float64 `gorm:"column:price_estimate_per_wa" json:"price_estimate_per_wa"`
	Notes *string `gorm:"column:notes" json:"notes"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	CreatedBy *uuid.UUID `gorm:"column:created_by" json:"created_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	UpdatedBy *uuid.UUID `gorm:"column:updated_by" json:"updated_by"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	DeletedBy *uuid.UUID `gorm:"column:deleted_by" json:"deleted_by"`
	TotalPrice *float64 `gorm:"column:total_price" json:"total_price"`
	ParcelNo *string `gorm:"column:parcel_no" json:"parcel_no"`
	DeedNo *string `gorm:"column:deed_no" json:"deed_no"`
	MapsheetNo *string `gorm:"column:mapsheet_no" json:"mapsheet_no"`
	LandCode *string `gorm:"column:land_code" json:"land_code"`
	SourceName *string `gorm:"column:source_name" json:"source_name"`
	SourceDt *string `gorm:"column:source_dt" json:"source_dt"`
	SyncId *uuid.UUID `gorm:"column:sync_id" json:"sync_id"`
	MuniCode *string `gorm:"column:muni_code" json:"muni_code"`
	IsEstimatedByTreasury *bool `gorm:"column:is_estimated_by_treasury" json:"is_estimated_by_treasury"`
	MapLandId *uuid.UUID `gorm:"column:map_land_id" json:"map_land_id"`
	LandZone *string `gorm:"column:land_zone" json:"land_zone"`
	LandDistrictId *int `gorm:"column:land_district_id" json:"land_district_id"`
	EstimatedPriceByTreasury *float64 `gorm:"column:estimated_price_by_treasury" json:"estimated_price_by_treasury"`
	Note *string `gorm:"column:note" json:"note"`
	SpecialUsetaxType *string `gorm:"column:special_usetax_type" json:"special_usetax_type"`
}
