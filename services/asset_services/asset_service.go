package asset_services

import (
	"citizen_system_back/models"
	"gorm.io/gorm"
)

// GetAllAssets retrieves all assets from the database
func GetAllAssets(db *gorm.DB) ([]models.Asset, error) {
	var assets []models.Asset
	// Replace this with actual database fetch logic
	err := db.Find(&assets).Error
	return assets, err
}
