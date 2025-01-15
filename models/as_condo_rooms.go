package models

import (
	"time"

	"github.com/google/uuid"
)

// AsCondoRooms represents the structure of the as_condo_rooms table
type AsCondoRooms struct {
	Id              uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	CondoId         *uuid.UUID `gorm:"column:condo_id" json:"condo_id"`
	Floor           *string    `gorm:"column:floor" json:"floor"`
	BuiltYear       *string    `gorm:"column:built_year" json:"built_year"`
	HouseholdTypeId *string    `gorm:"column:household_type_id" json:"household_type_id"`
	PricePerMeter   *float64   `gorm:"column:price_per_meter" json:"price_per_meter"`
	Width           *float64   `gorm:"column:width" json:"width"`
	Length          *float64   `gorm:"column:length" json:"length"`
	Area            *float64   `gorm:"column:area" json:"area"`
	RoomNo          *string    `gorm:"column:room_no" json:"room_no"`
	TaxDeductionId  *string    `gorm:"column:tax_deduction_id" json:"tax_deduction_id"`
	CreatedAt       *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	CreatedBy       *uuid.UUID `gorm:"column:created_by" json:"created_by"`
	UpdatedAt       *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	UpdatedBy       *uuid.UUID `gorm:"column:updated_by" json:"updated_by"`
	DeletedAt       *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	DeletedBy       *uuid.UUID `gorm:"column:deleted_by" json:"deleted_by"`
	MuniCode        *string    `gorm:"column:muni_code" json:"muni_code"`
	TaxYear         *string    `gorm:"column:tax_year" json:"tax_year"`
	RoomCode        *string    `gorm:"column:room_code" json:"room_code"`
	SyncId          *uuid.UUID `gorm:"column:sync_id" json:"sync_id"`
	SourceDt        *string    `gorm:"column:source_dt" json:"source_dt"`
	Note            *string    `gorm:"column:note" json:"note"`
	RoomTypeId      *uuid.UUID `gorm:"column:room_type_id" json:"room_type_id"`
	Rental          *bool      `gorm:"column:rental" json:"rental"`
	RentYear        *string    `gorm:"column:rent_year" json:"rent_year"`
	TotalRent       *string    `gorm:"column:total_rent" json:"total_rent"`
	CondoRoomCode   *string    `gorm:"column:condo_room_code" json:"condo_room_code"`
	RemarkYearFrom  *string    `gorm:"column:remark_year_from" json:"remark_year_from"`
	RemarkYearTo    *string    `gorm:"column:remark_year_to" json:"remark_year_to"`
}

func (AsCondoRooms) TableName() string {
	return "as_condo_rooms"
}
