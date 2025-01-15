package models

import (
)

// ProgressStatus represents the structure of the progress_status table
type ProgressStatus struct {
	ProgressId int `gorm:"column:progress_id" json:"progress_id"`
	ProgressName *string `gorm:"column:progress_name" json:"progress_name"`
	ProgressReal *string `gorm:"column:progress_real" json:"progress_real"`
	Active *int `gorm:"column:active" json:"active"`
}
