package repositories

import (
	"your-project/models"
	"gorm.io/gorm"
)

// GetAllAssets fetches all assets from the database
func GetAllAssets() ([]models.Asset, error) {
	var assets []models.Asset
	err := db.Find(&assets).Error 
	return assets, err
}

// CreateAsset creates a new asset in the database
func CreateAsset(asset models.Asset) (models.Asset, error) {
	err := db.Create(&asset).Error
	return asset, err
}
