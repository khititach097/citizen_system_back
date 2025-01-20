package routes

import (
	middleware "citizen_system_back/middleware"
	asset_routes "citizen_system_back/modules/assets"
	dev_tools_routes "citizen_system_back/modules/dev_tools"
	users_routes "citizen_system_back/modules/users"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all application routes.
func RegisterRoutes(router *gin.Engine) {
	// Register Asset routes with the database passed as a dependency
	assetGroup := router.Group("/api/v1")
	assetGroup.Use(middleware.AuthMiddleware)
	asset_routes.RegisterAssetRoutes(assetGroup)
	dev_tools_routes.RegisterDevToolsRoutes(assetGroup)
	users_routes.RegisterAssetRoutes(assetGroup)
}
