package asset_routes

import (
	asset_controllers "citizen_system_back/modules/assets/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterAssetRoutes registers routes for asset-related operations
// @Security BearerAuth
// @Security CookieAuth
func RegisterAssetRoutes(router *gin.RouterGroup) {

	assetGroup := router.Group("/assets") // Corrected variable name
	{
		assetGroup.GET("/", asset_controllers.ListAssets())
		// router.GET("/assets", middleware.AuthMiddleware, asset_controllers.ListAssets())
		assetGroup.GET("/:id", asset_controllers.GetAssetByID())
		assetGroup.GET("/get_asset_by_user_id/:user_id", asset_controllers.GetAssetByUserID())
		assetGroup.GET("/get_asset_by_asset_id/:asset_id", asset_controllers.GetAssetByLandID())
	}
}
