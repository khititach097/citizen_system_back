package models

import (
	"time"

	"github.com/google/uuid"
)

// AsCondos represents the structure of the as_condos table
type AsCondos struct {
	Id             uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	CreatedAt      *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	CondoName      *string    `gorm:"column:condo_name" json:"condo_name"`
	BuildingName   *string    `gorm:"column:building_name" json:"building_name"`
	CreatedBy      *uuid.UUID `gorm:"column:created_by" json:"created_by"`
	BuildingNo     *string    `gorm:"column:building_no" json:"building_no"`
	UpdatedAt      *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	UpdatedBy      *uuid.UUID `gorm:"column:updated_by" json:"updated_by"`
	RegistrationNo *string    `gorm:"column:registration_no" json:"registration_no"`
	Floor          *int       `gorm:"column:floor" json:"floor"`
	DeletedAt      *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	Room           *int       `gorm:"column:room" json:"room"`
	DeletedBy      *uuid.UUID `gorm:"column:deleted_by" json:"deleted_by"`
	LandId         *uuid.UUID `gorm:"column:land_id" json:"land_id"`
	BuiltYear      *string    `gorm:"column:built_year" json:"built_year"`
	PricePerMeter  *float64   `gorm:"column:price_per_meter" json:"price_per_meter"`
	MapGeometry    *string    `gorm:"column:map_geometry" json:"map_geometry"`
	Address        *string    `gorm:"column:address" json:"address"`
	MapLat         *string    `gorm:"column:map_lat" json:"map_lat"`
	Moo            *string    `gorm:"column:moo" json:"moo"`
	MapLong        *string    `gorm:"column:map_long" json:"map_long"`
	Village        *string    `gorm:"column:village" json:"village"`
	UtmMap1        *string    `gorm:"column:utm_map1" json:"utm_map1"`
	Soi            *string    `gorm:"column:soi" json:"soi"`
	UtmMap2        *string    `gorm:"column:utm_map2" json:"utm_map2"`
	Road           *string    `gorm:"column:road" json:"road"`
	UtmMap3        *string    `gorm:"column:utm_map3" json:"utm_map3"`
	SubdistrictId  *string    `gorm:"column:subdistrict_id" json:"subdistrict_id"`
	UtmMap4        *string    `gorm:"column:utm_map4" json:"utm_map4"`
	DistrictId     *string    `gorm:"column:district_id" json:"district_id"`
	UtmScale       *string    `gorm:"column:utm_scale" json:"utm_scale"`
	ProvinceId     *string    `gorm:"column:province_id" json:"province_id"`
	MuniCode       *string    `gorm:"column:muni_code" json:"muni_code"`
	TaxYear        *string    `gorm:"column:tax_year" json:"tax_year"`
	PostalCode     *string    `gorm:"column:postal_code" json:"postal_code"`
	SignboardId    *uuid.UUID `gorm:"column:signboard_id" json:"signboard_id"`
	RequestId      *uuid.UUID `gorm:"column:request_id" json:"request_id"`
	CondoCode      *string    `gorm:"column:condo_code" json:"condo_code"`
	SyncId         *uuid.UUID `gorm:"column:sync_id" json:"sync_id"`
	SourceDt       *string    `gorm:"column:source_dt" json:"source_dt"`
	Note           *string    `gorm:"column:note" json:"note"`
	ParcelNo       *string    `gorm:"column:parcel_no" json:"parcel_no"`
}

func (AsCondos) TableName() string {
	return "as_condos"
}
