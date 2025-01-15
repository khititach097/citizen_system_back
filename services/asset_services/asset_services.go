package asset_services

import (
	"citizen_system_back/models"
	"gorm.io/gorm"
)

// GetAllAssets retrieves all assets from the database
func GetAllAssets(db *gorm.DB) ([]models.AsLands, error) {
	var assets []models.AsLands
	err := db.Limit(10).Find(&assets).Error
	return assets, err
}

// GetAssetByID retrieves a single asset by its ID
//get geometry with WKB (Well-Known Binary)
// func GetAssetByID(db *gorm.DB, id string) (*models.AsLands, error) {
// 	var asset models.AsLands
// 	if err := db.Where("id = ?", id).First(&asset).Error; err != nil {
// 		return nil, err
// 	}
// 	return &asset, nil
// }

//get geometry with GeoJSON 
func GetAssetByID(db *gorm.DB, id string) (*models.AsLands, error) {
	var asset models.AsLands
	if err := db.Select("*, ST_AsGeoJSON(map_geometry) as map_geometry").
		Where("id = ?", id).
		First(&asset).Error; err != nil {
		return nil, err
	}
	return &asset, nil
}

func GetAssetByUserID(db *gorm.DB, user_id string) (*[]models.Citizen, error) {
	var citizen []models.Citizen
	if err := db.Where("citizen_id = ?", user_id).Find(&citizen).Error; err != nil {
		return nil, err
	}
	return &citizen, nil
}
