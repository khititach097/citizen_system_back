package asset_routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"citizen_system_back/controllers/asset_controllers"
)

// RegisterAssetRoutes registers routes for asset-related operations
func RegisterAssetRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/assets", asset_controllers.ListAssets(db))
	router.GET("/assets/:id", asset_controllers.GetAssetByID(db))
}