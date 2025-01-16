package models

import (
	"time"

	"github.com/google/uuid"
)

// AsLands represents the structure of the as_lands table
type AsLands struct {
	Id                       uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	CreatedAt                *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	CreatedBy                *uuid.UUID `gorm:"column:created_by" json:"created_by"`
	UpdatedAt                *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	UpdatedBy                *uuid.UUID `gorm:"column:updated_by" json:"updated_by"`
	DeletedAt                *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	DeletedBy                *uuid.UUID `gorm:"column:deleted_by" json:"deleted_by"`
	ParcelNo                 *string    `gorm:"column:parcel_no" json:"parcel_no"`
	LandNo                   *string    `gorm:"column:land_no" json:"land_no"`
	SurveyNo                 *string    `gorm:"column:survey_no" json:"survey_no"`
	LandSpaceRai             *float64   `gorm:"column:land_space_rai" json:"land_space_rai"`
	LandSpaceNgan            *float64   `gorm:"column:land_space_ngan" json:"land_space_ngan"`
	LandSpaceWa              *float64   `gorm:"column:land_space_wa" json:"land_space_wa"`
	LandSpaceSubWa           *float64   `gorm:"column:land_space_sub_wa" json:"land_space_sub_wa"`
	MapGeometry              *string    `gorm:"column:map_geometry" json:"map_geometry"`
	MapLat                   *string    `gorm:"column:map_lat" json:"map_lat"`
	MapLong                  *string    `gorm:"column:map_long" json:"map_long"`
	UtmMap1                  *string    `gorm:"column:utm_map1" json:"utm_map1"`
	UtmMap2                  *string    `gorm:"column:utm_map2" json:"utm_map2"`
	UtmMap3                  *string    `gorm:"column:utm_map3" json:"utm_map3"`
	UtmMap4                  *string    `gorm:"column:utm_map4" json:"utm_map4"`
	UtmScale                 *string    `gorm:"column:utm_scale" json:"utm_scale"`
	TaxYear                  *string    `gorm:"column:tax_year" json:"tax_year"`
	MuniCode                 *string    `gorm:"column:muni_code" json:"muni_code"`
	ParcelType               *int       `gorm:"column:parcel_type" json:"parcel_type"`
	SourceDt                 *time.Time `gorm:"column:source_dt" json:"source_dt"`
	DeedNo                   *string    `gorm:"column:deed_no" json:"deed_no"`
	CutaxLandId              *string    `gorm:"column:cutax_land_id" json:"cutax_land_id"`
	SyncId                   *uuid.UUID `gorm:"column:sync_id" json:"sync_id"`
	EstimatedPrice           *float64   `gorm:"column:estimated_price" json:"estimated_price"`
	IsEstimatedByTreasury    *bool      `gorm:"column:is_estimated_by_treasury" json:"is_estimated_by_treasury"`
	LandZone                 *string    `gorm:"column:land_zone" json:"land_zone"`
	LandDistrictId           *int       `gorm:"column:land_district_id" json:"land_district_id"`
	LandSubDistrictId        *int       `gorm:"column:land_sub_district_id" json:"land_sub_district_id"`
	EstimatedPriceByTreasury *float64   `gorm:"column:estimated_price_by_treasury" json:"estimated_price_by_treasury"`
	Note                     *string    `gorm:"column:note" json:"note"`
	SpecialUsetaxType        *string    `gorm:"column:special_usetax_type" json:"special_usetax_type"`
	Road                     *string    `gorm:"column:road" json:"road"`
}

func (AsLands) TableName() string {
	return "as_lands"
}
