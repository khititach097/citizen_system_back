package models

import (
	"time"

	"github.com/google/uuid"
)

// AsBuildings represents the structure of the as_buildings table
type AsBuildings struct {
	Id                         uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	CreatedAt                  *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	CreatedBy                  *uuid.UUID `gorm:"column:created_by" json:"created_by"`
	UpdatedAt                  *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	UpdatedBy                  *uuid.UUID `gorm:"column:updated_by" json:"updated_by"`
	DeletedAt                  *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	DeletedBy                  *uuid.UUID `gorm:"column:deleted_by" json:"deleted_by"`
	PropertyBuildingMainTypeId *string    `gorm:"column:property_building_main_type_id" json:"property_building_main_type_id"`
	PropertyBuildingSubTypeId  *string    `gorm:"column:property_building_sub_type_id" json:"property_building_sub_type_id"`
	BuildingNo                 *string    `gorm:"column:building_no" json:"building_no"`
	BuildingDesignTypeId       *int       `gorm:"column:building_design_type_id" json:"building_design_type_id"`
	BuildingTotalRoom          *int       `gorm:"column:building_total_room" json:"building_total_room"`
	BuildingTotalFloor         *string    `gorm:"column:building_total_floor" json:"building_total_floor"`
	BuildingWidthMeter         *float64   `gorm:"column:building_width_meter" json:"building_width_meter"`
	BuildingLengthMeter        *float64   `gorm:"column:building_length_meter" json:"building_length_meter"`
	Note                       *string    `gorm:"column:note" json:"note"`
	BuildingYearTotal          *string    `gorm:"column:building_year_total" json:"building_year_total"`
	LandId                     *uuid.UUID `gorm:"column:land_id" json:"land_id"`
	MapGeometry                *string    `gorm:"column:map_geometry" json:"map_geometry"`
	MapLat                     *string    `gorm:"column:map_lat" json:"map_lat"`
	MapLong                    *string    `gorm:"column:map_long" json:"map_long"`
	UtmMap1                    *string    `gorm:"column:utm_map1" json:"utm_map1"`
	UtmMap2                    *string    `gorm:"column:utm_map2" json:"utm_map2"`
	UtmMap3                    *string    `gorm:"column:utm_map3" json:"utm_map3"`
	UtmMap4                    *string    `gorm:"column:utm_map4" json:"utm_map4"`
	UtmScale                   *string    `gorm:"column:utm_scale" json:"utm_scale"`
	TaxYear                    *string    `gorm:"column:tax_year" json:"tax_year"`
	LandUsedId                 *uuid.UUID `gorm:"column:land_used_id" json:"land_used_id"`
	BuildingCode               *string    `gorm:"column:building_code" json:"building_code"`
	ParcelNo                   *string    `gorm:"column:parcel_no" json:"parcel_no"`
	BuildYear                  *string    `gorm:"column:build_year" json:"build_year"`
	BuildingAllArea            *string    `gorm:"column:building_all_area" json:"building_all_area"`
	BuildingBuildArea          *string    `gorm:"column:building_build_area" json:"building_build_area"`
	MuniCode                   *string    `gorm:"column:muni_code" json:"muni_code"`
	SourceDt                   *string    `gorm:"column:source_dt" json:"source_dt"`
	CutaxBuildingId            *string    `gorm:"column:cutax_building_id" json:"cutax_building_id"`
	SyncId                     *uuid.UUID `gorm:"column:sync_id" json:"sync_id"`
	BuildingTypeYear           *string    `gorm:"column:building_type_year" json:"building_type_year"`
	EmptyYear                  *string    `gorm:"column:empty_year" json:"empty_year"`
	RentYear                   *string    `gorm:"column:rent_year" json:"rent_year"`
	ParcelCode                 *string    `gorm:"column:parcel_code" json:"parcel_code"`
	BuildingHouseCode          *string    `gorm:"column:building_house_code" json:"building_house_code"`
	AddressZone                *string    `gorm:"column:address_zone" json:"address_zone"`
	Road                       *string    `gorm:"column:road" json:"road"`
}

func (AsBuildings) TableName() string {
	return "as_buildings"
}
