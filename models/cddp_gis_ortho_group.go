package models

import (
	"github.com/google/uuid"
	"time"
)

// CddpGisOrthoGroup represents the structure of the cddp_gis_ortho_group table
type CddpGisOrthoGroup struct {
	OrthoGroupId            uuid.UUID   `gorm:"column:ortho_group_id;primaryKey" json:"ortho_group_id"`
	OrthoGroupBound         *string     `gorm:"column:ortho_group_bound" json:"ortho_group_bound"`
	OrthoGroupType          *string     `gorm:"column:ortho_group_type" json:"ortho_group_type"`
	OrthoGroupFileName      *string     `gorm:"column:ortho_group_file_name" json:"ortho_group_file_name"`
	OrthoGroupFileExtension *string     `gorm:"column:ortho_group_file_extension" json:"ortho_group_file_extension"`
	OrthoGroupFileSize      *float64    `gorm:"column:ortho_group_file_size" json:"ortho_group_file_size"`
	OrthoCompressionRemark  *string     `gorm:"column:ortho_compression_remark" json:"ortho_compression_remark"`
	CreateDate              *time.Time  `gorm:"column:create_date" json:"create_date"`
	CreateBy                *string     `gorm:"column:create_by" json:"create_by"`
	UpdateDate              *time.Time  `gorm:"column:update_date" json:"update_date"`
	UpdateBy                string      `gorm:"column:update_by" json:"update_by"`
	MunicipalityCode        string      `gorm:"column:municipality_code" json:"municipality_code"`
	OrthoGroupMunicipalityCode *int     `gorm:"column:\"ortho_group/municipality_code\"" json:"ortho_group/municipality_code"` // Using backticks for special characters
}
