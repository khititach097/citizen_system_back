package models

import (
)

// MasterSubdistrict represents the structure of the master_subdistrict table
type MasterSubdistrict struct {
	AdRegion *int `gorm:"column:ad_region" json:"ad_region"`
	RegionCode *int `gorm:"column:region_code" json:"region_code"`
	RegionNameT *string `gorm:"column:region_name_t" json:"region_name_t"`
	RegionNameE *string `gorm:"column:region_name_e" json:"region_name_e"`
	AdProvince *int `gorm:"column:ad_province" json:"ad_province"`
	ProvinceCode *string `gorm:"column:province_code" json:"province_code"`
	ProvinceNameT *string `gorm:"column:province_name_t" json:"province_name_t"`
	AdDistrict *int `gorm:"column:ad_district" json:"ad_district"`
	DistrictCode *string `gorm:"column:district_code" json:"district_code"`
	DistrictNameT *string `gorm:"column:district_name_t" json:"district_name_t"`
	DistrictNameE *string `gorm:"column:district_name_e" json:"district_name_e"`
	AdSubdistrict *int `gorm:"column:ad_subdistrict" json:"ad_subdistrict"`
	SubdistrictCode *string `gorm:"column:subdistrict_code" json:"subdistrict_code"`
	SubdistrictNameT *string `gorm:"column:subdistrict_name_t" json:"subdistrict_name_t"`
	SubdistrictNameE *string `gorm:"column:subdistrict_name_e" json:"subdistrict_name_e"`
	ShapeLen *string `gorm:"column:shape_len" json:"shape_len"`
	ShapeArea *string `gorm:"column:shape_area" json:"shape_area"`
	Geometry *string `gorm:"column:geometry" json:"geometry"`
	GeomId *string `gorm:"column:geom_id" json:"geom_id"`
	ReportDt *string `gorm:"column:report_dt" json:"report_dt"`
	ProvinceNameE *string `gorm:"column:province_name_e" json:"province_name_e"`
	Postcode *string `gorm:"column:postcode" json:"postcode"`
}
