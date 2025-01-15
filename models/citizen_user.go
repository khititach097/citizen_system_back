package models

import (
	"github.com/google/uuid"
	"time"
)

// CitizenUser represents the structure of the citizen_user table
type CitizenUser struct {
	Id uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserEmail string `gorm:"column:user_email" json:"user_email"`
	Password *string `gorm:"column:password" json:"password"`
	UserTel *string `gorm:"column:user_tel" json:"user_tel"`
	Active *bool `gorm:"column:active" json:"active"`
	UserName *string `gorm:"column:user_name" json:"user_name"`
	UserLastName *string `gorm:"column:user_last_name" json:"user_last_name"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	CreatedBy *string `gorm:"column:created_by" json:"created_by"`
	UpdatedBy *string `gorm:"column:updated_by" json:"updated_by"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	UserStatus *int `gorm:"column:user_status" json:"user_status"`
	UserId *string `gorm:"column:user_id" json:"user_id"`
	ForceLogout *bool `gorm:"column:force_logout" json:"force_logout"`
	LastAccess *time.Time `gorm:"column:last_access" json:"last_access"`
}
