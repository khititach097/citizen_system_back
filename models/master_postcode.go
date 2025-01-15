package models

// MasterPostcode represents the structure of the master_postcode table
type MasterPostcode struct {
	Tambonthai       *string `gorm:"column:tambonthai" json:"tambonthai"`
	Tamboneng        *string `gorm:"column:tamboneng" json:"tamboneng"`
	Districtid       *string `gorm:"column:districtid" json:"districtid"`
	Districtthai     *string `gorm:"column:districtthai" json:"districtthai"`
	Districteng      *string `gorm:"column:districteng" json:"districteng"`
	Provinceid       *string `gorm:"column:provinceid" json:"provinceid"`
	Provincethai     *string `gorm:"column:provincethai" json:"provincethai"`
	Provinceeng      *string `gorm:"column:provinceeng" json:"provinceeng"`
	Postalcoderemark *string `gorm:"column:postalcoderemark" json:"postalcoderemark"`
	Postcodemain     *string `gorm:"column:postcodemain" json:"postcodemain"`
	Postcodeall      *string `gorm:"column:postcodeall" json:"postcodeall"`
	ReportDt         *string `gorm:"column:report_dt" json:"report_dt"`
	SubdistrictCode  *string `gorm:"column:subdistrict_code" json:"subdistrict_code"`
}

func (MasterPostcode) TableName() string {
	return "master_postcode"
}
