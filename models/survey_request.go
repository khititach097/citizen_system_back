package models

import (
	"github.com/google/uuid"
	"time"
)

// SurveyRequest represents the structure of the survey_request table
type SurveyRequest struct {
	Id uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	TaxYear *string `gorm:"column:tax_year" json:"tax_year"`
	PriorityStatusId *int `gorm:"column:priority_status_id" json:"priority_status_id"`
	LandId *uuid.UUID `gorm:"column:land_id" json:"land_id"`
	SurveyInquirerId *uuid.UUID `gorm:"column:survey_inquirer_id" json:"survey_inquirer_id"`
	ReceiveRequestById *uuid.UUID `gorm:"column:receive_request_by_id" json:"receive_request_by_id"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	Status *int `gorm:"column:status" json:"status"`
	PersonTypeId *int `gorm:"column:person_type_id" json:"person_type_id"`
	IdCardNum *string `gorm:"column:id_card_num" json:"id_card_num"`
	PhoneNumber *string `gorm:"column:phone_number" json:"phone_number"`
	Note *string `gorm:"column:note" json:"note"`
	SurveyCode *string `gorm:"column:survey_code" json:"survey_code"`
	FirstName *string `gorm:"column:first_name" json:"first_name"`
	LastName *string `gorm:"column:last_name" json:"last_name"`
	CorporateName *string `gorm:"column:corporate_name" json:"corporate_name"`
	CreatedBy *uuid.UUID `gorm:"column:created_by" json:"created_by"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	DeletedBy *uuid.UUID `gorm:"column:deleted_by" json:"deleted_by"`
	MuniCode *string `gorm:"column:muni_code" json:"muni_code"`
	ProvinceCode *string `gorm:"column:province_code" json:"province_code"`
	SurveyTransaction *string `gorm:"column:survey_transaction" json:"survey_transaction"`
	ApprovedBy *uuid.UUID `gorm:"column:approved_by" json:"approved_by"`
	ApprovedAt *time.Time `gorm:"column:approved_at" json:"approved_at"`
	NoteLand *string `gorm:"column:note_land" json:"note_land"`
	NoteUsedLandBuilding *string `gorm:"column:note_used_land_building" json:"note_used_land_building"`
	NoteSignboard *string `gorm:"column:note_signboard" json:"note_signboard"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	UpdatedBy *uuid.UUID `gorm:"column:updated_by" json:"updated_by"`
	PropertyProcessDetailId *int `gorm:"column:property_process_detail_id" json:"property_process_detail_id"`
	NoteCondo *string `gorm:"column:note_condo" json:"note_condo"`
	SurveyTypeId *int `gorm:"column:survey_type_id" json:"survey_type_id"`
}
