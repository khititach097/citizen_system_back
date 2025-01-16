package models

import (
	"time"

	"github.com/google/uuid"
)

// PatchLog represents the structure of the patch_log table
type PatchLog struct {
	Id         uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	Version    *string    `gorm:"column:version" json:"version"`
	DetailLog  *string    `gorm:"column:detail_log" json:"detail_log"`
	UpdateDate *time.Time `gorm:"column:update_date" json:"update_date"`
}

func (PatchLog) TableName() string {
	return "patch_log"
}
