package models

// MasterPlatform represents the structure of the master_platform table
type MasterPlatform struct {
	PlatformId   int     `gorm:"column:platform_id;primaryKey" json:"platform_id"`
	PlatformName *string `gorm:"column:platform_name" json:"platform_name"`
}

func (MasterPlatform) TableName() string {
	return "master_platform"
}
