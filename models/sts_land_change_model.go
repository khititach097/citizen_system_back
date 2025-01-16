package models

import (
	"time"

	"github.com/google/uuid"
)

// StsLandChange represents the structure of the sts_land_change table
type StsLandChange struct {
	ChangeId       uuid.UUID  `gorm:"column:change_id;primaryKey" json:"change_id"`
	LandId         *uuid.UUID `gorm:"column:land_id" json:"land_id"`
	Geometry       *string    `gorm:"column:geometry" json:"geometry"`
	ParcelNo       *string    `gorm:"column:parcel_no" json:"parcel_no"`
	ClassName      *string    `gorm:"column:class_name" json:"class_name"`
	LandUsedStatus *bool      `gorm:"column:land_used_status" json:"land_used_status"`
	ParcelCode     *string    `gorm:"column:parcel_code" json:"parcel_code"`
	ParcelType     *int       `gorm:"column:parcel_type" json:"parcel_type"`
	LandNo         *string    `gorm:"column:land_no" json:"land_no"`
	SurveyNo       *string    `gorm:"column:survey_no" json:"survey_no"`
	Mapsheet       *string    `gorm:"column:mapsheet" json:"mapsheet"`
	Annual         *string    `gorm:"column:annual" json:"annual"`
	MuniCode       *string    `gorm:"column:muni_code" json:"muni_code"`
	SourceName     *string    `gorm:"column:source_name" json:"source_name"`
	SourceDt       *string    `gorm:"column:source_dt" json:"source_dt"`
	SourceSrid     *int       `gorm:"column:source_srid" json:"source_srid"`
	CreatedDate    *time.Time `gorm:"column:created_date" json:"created_date"`
	CreatedBy      *string    `gorm:"column:created_by" json:"created_by"`
	UpdatedDate    *string    `gorm:"column:updated_date" json:"updated_date"`
	UpdatedBy      *string    `gorm:"column:updated_by" json:"updated_by"`
	Remarks        *string    `gorm:"column:remarks" json:"remarks"`
	ImportDate     *time.Time `gorm:"column:import_date" json:"import_date"`
}

func (StsLandChange) TableName() string {
	return "sts_land_change"
}
