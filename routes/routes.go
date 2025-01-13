package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"citizen_system_back/controllers/asset_controller"
)

// RegisterRoutes registers all application routes.
func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	// Register Asset routes with the database passed as a dependency
	assetGroup := router.Group("/api/v1")
	asset_controllers.RegisterAssetRoutes(assetGroup, db)
}
