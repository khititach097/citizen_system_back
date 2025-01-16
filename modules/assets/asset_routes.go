package asset_routes

import (
	asset_controllers "citizen_system_back/modules/assets/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterAssetRoutes registers routes for asset-related operations
// @Security CookieAuth
func RegisterAssetRoutes(router *gin.RouterGroup) {
	router.GET("/assets", asset_controllers.ListAssets())
	// router.GET("/assets", middleware.AuthMiddleware, asset_controllers.ListAssets())
	router.GET("/assets/:id", asset_controllers.GetAssetByID())
	router.GET("/assets/get_asset_by_user_id/:user_id", asset_controllers.GetAssetByUserID())
}
