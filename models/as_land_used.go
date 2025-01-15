package models

import (
	"github.com/google/uuid"
	"time"
)

// AsLandUsed represents the structure of the as_land_used table
type AsLandUsed struct {
	Id uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	CreatedBy *uuid.UUID `gorm:"column:created_by" json:"created_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	UpdatedBy *uuid.UUID `gorm:"column:updated_by" json:"updated_by"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	DeletedBy *uuid.UUID `gorm:"column:deleted_by" json:"deleted_by"`
	LandId *uuid.UUID `gorm:"column:land_id" json:"land_id"`
	HouseholdTypeId *int `gorm:"column:household_type_id" json:"household_type_id"`
	UsingTypeId *int `gorm:"column:using_type_id" json:"using_type_id"`
	UsingDetailId *string `gorm:"column:using_detail_id" json:"using_detail_id"`
	TotalSpaceSquareWa *float64 `gorm:"column:total_space_square_wa" json:"total_space_square_wa"`
	CutaxLandusedId *string `gorm:"column:cutax_landused_id" json:"cutax_landused_id"`
	SyncId *uuid.UUID `gorm:"column:sync_id" json:"sync_id"`
	TotalSpaceRai *int `gorm:"column:total_space_rai" json:"total_space_rai"`
	TotalSpaceNgan *int `gorm:"column:total_space_ngan" json:"total_space_ngan"`
	LandUsedCode *string `gorm:"column:land_used_code" json:"land_used_code"`
	Note *string `gorm:"column:note" json:"note"`
	MuniCode *string `gorm:"column:muni_code" json:"muni_code"`
	UsingRentId *int `gorm:"column:using_rent_id" json:"using_rent_id"`
	Rai *float64 `gorm:"column:rai" json:"rai"`
	Ngan *float64 `gorm:"column:ngan" json:"ngan"`
	Wa *string `gorm:"column:wa" json:"wa"`
	EmptyYear *string `gorm:"column:empty_year" json:"empty_year"`
	TaxDeductionId *string `gorm:"column:tax_deduction_id" json:"tax_deduction_id"`
	ImpactEndYear *string `gorm:"column:impact_end_year" json:"impact_end_year"`
	ImpactStartYear *string `gorm:"column:impact_start_year" json:"impact_start_year"`
	RemarkYearFrom *string `gorm:"column:remark_year_from" json:"remark_year_from"`
	RemarkYearTo *string `gorm:"column:remark_year_to" json:"remark_year_to"`
}
