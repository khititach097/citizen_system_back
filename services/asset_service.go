package services

import (
	"your-project/models"
	"your-project/repositories"
)

// Asset represents the asset data structure
type Asset struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// GetAllAssets retrieves all assets
func GetAllAssets() ([]Asset, error) {
	return repositories.GetAllAssets()
}

// CreateAsset creates a new asset
func CreateAsset(asset Asset) (Asset, error) {
	return repositories.CreateAsset(asset)
}
