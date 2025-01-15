package models

import (
)

// MasterMunicipality represents the structure of the master_municipality table
type MasterMunicipality struct {
	MunicipalityCode *string `gorm:"column:municipality_code" json:"municipality_code"`
	MunicipalityNameT *string `gorm:"column:municipality_name_t" json:"municipality_name_t"`
	DistrictNameT *string `gorm:"column:district_name_t" json:"district_name_t"`
	Geometry *string `gorm:"column:geometry" json:"geometry"`
	GeomId *string `gorm:"column:geom_id" json:"geom_id"`
	ReportDt *string `gorm:"column:report_dt" json:"report_dt"`
	ProvinceNameT *string `gorm:"column:province_name_t" json:"province_name_t"`
	ProvinceCode *string `gorm:"column:province_code" json:"province_code"`
	DistrictCode *string `gorm:"column:district_code" json:"district_code"`
	SubDistrictCode *string `gorm:"column:sub_district_code" json:"sub_district_code"`
	DefaultLocation *string `gorm:"column:default_location" json:"default_location"`
	AccessPlatform *string `gorm:"column:access_platform" json:"access_platform"`
	MuniLocation *string `gorm:"column:muni_location" json:"muni_location"`
}
