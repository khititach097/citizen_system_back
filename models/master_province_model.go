package models

// MasterProvince represents the structure of the master_province table
type MasterProvince struct {
	AdRegion      *int    `gorm:"column:ad_region" json:"ad_region"`
	RegionCode    *string `gorm:"column:region_code" json:"region_code"`
	RegionNameT   *string `gorm:"column:region_name_t" json:"region_name_t"`
	RegionNameE   *string `gorm:"column:region_name_e" json:"region_name_e"`
	AdProvince    *int    `gorm:"column:ad_province" json:"ad_province"`
	ProvinceCode  *string `gorm:"column:province_code" json:"province_code"`
	ProvinceNameT *string `gorm:"column:province_name_t" json:"province_name_t"`
	ShapeLen      *string `gorm:"column:shape_len" json:"shape_len"`
	ShapeArea     *string `gorm:"column:shape_area" json:"shape_area"`
	Geometry      *string `gorm:"column:geometry" json:"geometry"`
	GeomId        *string `gorm:"column:geom_id" json:"geom_id"`
	ReportDt      *string `gorm:"column:report_dt" json:"report_dt"`
	ProvinceNameE *string `gorm:"column:province_name_e" json:"province_name_e"`
}

func (MasterProvince) TableName() string {
	return "master_province"
}
