package models

import (
	"github.com/google/uuid"
	"time"
)

// AsLand represents the structure of the as_lands table
type AsLand struct {
	ID                     uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	CreatedAt              *time.Time     `json:"created_at"`
	CreatedBy              *uuid.UUID     `gorm:"type:uuid" json:"created_by"`
	UpdatedAt              *time.Time     `json:"updated_at"`
	UpdatedBy              *uuid.UUID     `gorm:"type:uuid" json:"updated_by"`
	DeletedAt              *time.Time     `json:"deleted_at"`
	DeletedBy              *uuid.UUID     `gorm:"type:uuid" json:"deleted_by"`
	ParcelNo              *string         `gorm:"type:varchar(255)" json:"parcel_no"`
	LandNo                *string         `gorm:"type:varchar(255)" json:"land_no"`
	SurveyNo             *string         `gorm:"type:varchar(255)" json:"survey_no"`
	LandSpaceRai         *float64        `gorm:"type:float8" json:"land_space_rai"`
	LandSpaceNgan        *float64        `gorm:"type:float8" json:"land_space_ngan"`
	LandSpaceWa          *float64        `gorm:"type:float8" json:"land_space_wa"`
	LandSpaceSubWa       *float64        `gorm:"type:float8" json:"land_space_sub_wa"`
	MapGeometry          *string         `gorm:"type:geometry" json:"map_geometry"`
	MapLat               *string         `gorm:"type:varchar" json:"map_lat"`
	MapLong              *string         `gorm:"type:varchar" json:"map_long"`
	UtmMap1              *string         `gorm:"type:varchar" json:"utm_map1"`
	UtmMap2              *string         `gorm:"type:varchar" json:"utm_map2"`
	UtmMap3              *string         `gorm:"type:varchar" json:"utm_map3"`
	UtmMap4              *string         `gorm:"type:varchar" json:"utm_map4"`
	UtmScale             *string         `gorm:"type:varchar" json:"utm_scale"`
	TaxYear              *string         `gorm:"type:varchar" json:"tax_year"`
	MuniCode            *string         `gorm:"type:varchar" json:"muni_code"`
	ParcelType          *int            `gorm:"type:int" json:"parcel_type"`
	SourceDt            *time.Time      `json:"source_dt"`
	DeedNo              *string         `gorm:"type:varchar" json:"deed_no"`
	CutaxLandID         *string         `gorm:"type:varchar(20)" json:"cutax_land_id"`
	SyncID              *uuid.UUID      `gorm:"type:uuid" json:"sync_id"`
	EstimatedPrice      *float64        `gorm:"type:float8" json:"estimated_price"`
	IsEstimatedByTreasury *bool          `json:"is_estimated_by_treasury"`
	LandZone            *string         `gorm:"type:varchar" json:"land_zone"`
	LandDistrictID      *int            `gorm:"type:int" json:"land_district_id"`
	LandSubDistrictID   *int            `gorm:"type:int" json:"land_sub_district_id"`
	EstimatedPriceByTreasury *float64    `gorm:"type:float8" json:"estimated_price_by_treasury"`
	Note                *string         `gorm:"type:text" json:"note"`
	SpecialUseTaxType   *string         `gorm:"type:varchar(255)" json:"special_usetax_type"`
	Road                *string         `gorm:"type:text" json:"road"`
}
