package models

import (
	"github.com/google/uuid"
	"time"
)

// CddpGisOrthoIndex represents the structure of the cddp_gis_ortho_index table
type CddpGisOrthoIndex struct {
	OrthoId uuid.UUID `gorm:"column:ortho_id;primaryKey" json:"ortho_id"`
	OrthoGroupId uuid.UUID `gorm:"column:ortho_group_id" json:"ortho_group_id"`
	OrthoBound *string `gorm:"column:ortho_bound" json:"ortho_bound"`
	OrthoType *string `gorm:"column:ortho_type" json:"ortho_type"`
	OrthoFileName *string `gorm:"column:ortho_file_name" json:"ortho_file_name"`
	OrthoFileExtension *string `gorm:"column:ortho_file_extension" json:"ortho_file_extension"`
	OrthoFileSize *float64 `gorm:"column:ortho_file_size" json:"ortho_file_size"`
	OrthoTilesetId *string `gorm:"column:ortho_tileset_id" json:"ortho_tileset_id"`
	OrthoTilesetLayerId *string `gorm:"column:ortho_tileset_layer_id" json:"ortho_tileset_layer_id"`
	OrthoCompressionRemark *string `gorm:"column:ortho_compression_remark" json:"ortho_compression_remark"`
	OrthoIndexRow *int `gorm:"column:ortho_index_row" json:"ortho_index_row"`
	OrthoIndexCol *int `gorm:"column:ortho_index_col" json:"ortho_index_col"`
	CreateDate *time.Time `gorm:"column:create_date" json:"create_date"`
	CreateBy *string `gorm:"column:create_by" json:"create_by"`
	UpdateDate *time.Time `gorm:"column:update_date" json:"update_date"`
	UpdateBy string `gorm:"column:update_by" json:"update_by"`
}
