package routes

import (
	"github.com/gin-gonic/gin"

	"citizen_system_back/routes/asset_routes"
	"citizen_system_back/routes/dev_tools_routes"
)

// RegisterRoutes registers all application routes.
// @Security CookieAuth
func RegisterRoutes(router *gin.Engine) {
	// Register Asset routes with the database passed as a dependency
	assetGroup := router.Group("/api/v1")
	// assetGroup.Use(middleware.AuthMiddleware)
	asset_routes.RegisterAssetRoutes(assetGroup)
	dev_tools_routes.RegisterDevToolsRoutes(assetGroup)
}
