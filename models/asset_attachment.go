package models

import (
	"github.com/google/uuid"
	"time"
)

// AssetAttachment represents the structure of the asset_attachment table
type AssetAttachment struct {
	AssetId *uuid.UUID `gorm:"column:asset_id" json:"asset_id"`
	ImgId uuid.UUID `gorm:"column:img_id" json:"img_id"`
	Key *string `gorm:"column:key" json:"key"`
	OriginalFileName *string `gorm:"column:original_file_name" json:"original_file_name"`
	FileType *string `gorm:"column:file_type" json:"file_type"`
	FileSize *string `gorm:"column:file_size" json:"file_size"`
	CreatedDate *time.Time `gorm:"column:created_date" json:"created_date"`
	UpdatedDate *time.Time `gorm:"column:updated_date" json:"updated_date"`
	CreatedBy *string `gorm:"column:created_by" json:"created_by"`
	UpdatedBy *string `gorm:"column:updated_by" json:"updated_by"`
	AssetSelected *bool `gorm:"column:asset_selected" json:"asset_selected"`
	ImageFrom *int `gorm:"column:image_from" json:"image_from"`
	PropertyTypeId *int `gorm:"column:property_type_id" json:"property_type_id"`
	SurveyRequestId *uuid.UUID `gorm:"column:survey_request_id" json:"survey_request_id"`
	TaxYear *string `gorm:"column:tax_year" json:"tax_year"`
	Source *string `gorm:"column:source" json:"source"`
	ImportFrom *string `gorm:"column:import_from" json:"import_from"`
	MuniCode *string `gorm:"column:muni_code" json:"muni_code"`
}
