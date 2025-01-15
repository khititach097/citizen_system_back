package models

import (
	"time"

	"github.com/google/uuid"
)

// AsCondoRoomType represents the structure of the as_condo_room_type table
type AsCondoRoomType struct {
	Id                 uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	CondoId            uuid.UUID  `gorm:"column:condo_id" json:"condo_id"`
	UsageType          *string    `gorm:"column:usage_type" json:"usage_type"`
	UsageAreaType      *string    `gorm:"column:usage_area_type" json:"usage_area_type"`
	UsagePriceEstimate *string    `gorm:"column:usage_price_estimate" json:"usage_price_estimate"`
	Area               *string    `gorm:"column:area" json:"area"`
	TotalArea          *string    `gorm:"column:total_area" json:"total_area"`
	Notes              *string    `gorm:"column:notes" json:"notes"`
	CreatedAt          *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	CreatedBy          *uuid.UUID `gorm:"column:created_by" json:"created_by"`
	UpdatedAt          *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	UpdatedBy          *uuid.UUID `gorm:"column:updated_by" json:"updated_by"`
	DeletedAt          *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	DeletedBy          *uuid.UUID `gorm:"column:deleted_by" json:"deleted_by"`
	SourceDt           *string    `gorm:"column:source_dt" json:"source_dt"`
	MuniCode           *string    `gorm:"column:muni_code" json:"muni_code"`
	RoomTypeId         *uuid.UUID `gorm:"column:room_type_id" json:"room_type_id"`
	RoomDetail         *string    `gorm:"column:room_detail" json:"room_detail"`
}

func (AsCondoRoomType) TableName() string {
	return "as_condo_room_type"
}
