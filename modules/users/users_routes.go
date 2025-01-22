package users_routes

import (
	users_controllers "citizen_system_back/modules/users/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterAssetRoutes registers routes for asset-related operations
// @Security BearerAuth
func RegisterAssetRoutes(router *gin.RouterGroup) {
	usersGroup := router.Group("/users") // Corrected variable name
	// router.GET("/assets", middleware.AuthMiddleware, asset_controllers.ListAssets())
	{
		usersGroup.GET("/profile", users_controllers.GetProfile())
	}
}
