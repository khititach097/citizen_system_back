package models

import (
	"github.com/google/uuid"
	"time"
)

// MasterUserStatus represents the structure of the master_user_status table
type MasterUserStatus struct {
	Id uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	ModuleCode string `gorm:"column:module_code" json:"module_code"`
	ModuleDesT *string `gorm:"column:module_des_t" json:"module_des_t"`
	ModuleDesE *string `gorm:"column:module_des_e" json:"module_des_e"`
	StatusDesc *string `gorm:"column:status_desc" json:"status_desc"`
	ActionCode *string `gorm:"column:action_code" json:"action_code"`
	MasterStatus *int `gorm:"column:master_status" json:"master_status"`
	ModuleValue *int `gorm:"column:module_value" json:"module_value"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	CreatedBy *uuid.UUID `gorm:"column:created_by" json:"created_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	UpdatedBy *uuid.UUID `gorm:"column:updated_by" json:"updated_by"`
}
