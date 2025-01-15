package asset_routes

import (
	"citizen_system_back/controllers/asset_controllers"
	// "citizen_system_back/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterAssetRoutes registers routes for asset-related operations
// @Security BearerAuth
func RegisterAssetRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/assets", asset_controllers.ListAssets(db))
	// router.GET("/assets", middleware.AuthMiddleware, asset_controllers.ListAssets(db))
	router.GET("/assets/:id", asset_controllers.GetAssetByID(db))
	router.GET("/assets/get_asset_by_user_id/:user_id", asset_controllers.GetAssetByUserID(db))
}
