package asset_services

import (
	"citizen_system_back/models"
	"gorm.io/gorm"
)

// GetAllAssets retrieves all assets from the database
func GetAllAssets(db *gorm.DB) ([]models.AsLand, error) {
	var assets []models.AsLand
	err := db.Limit(10).Find(&assets).Error
	return assets, err
}

// GetAssetByID retrieves a single asset by its ID
func GetAssetByID(db *gorm.DB, id string) (*models.AsLand, error) {
	var asset models.AsLand
	if err := db.Where("id = ?", id).First(&asset).Error; err != nil {
		return nil, err
	}
	return &asset, nil
}
