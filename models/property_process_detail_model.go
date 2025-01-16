package models

// PropertyProcessDetail represents the structure of the property_process_detail table
type PropertyProcessDetail struct {
	PropertyProcessDetailId   int64   `gorm:"column:property_process_detail_id" json:"property_process_detail_id"`
	PropertyProcessDetailName *string `gorm:"column:property_process_detail_name" json:"property_process_detail_name"`
}

func (PropertyProcessDetail) TableName() string {
	return "property_process_detail"
}
