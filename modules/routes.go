package routes

import (
	middleware "citizen_system_back/middleware"
	asset_routes "citizen_system_back/modules/assets"
	dev_tools_routes "citizen_system_back/modules/dev_tools"
	s3_routes "citizen_system_back/modules/s3"
	users_routes "citizen_system_back/modules/users"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all application routes.
func RegisterRoutes(router *gin.Engine) {
	// Register Asset routes with the database passed as a dependency
	mainGroup := router.Group("/api/v1")
	mainGroup.Use(middleware.AuthMiddleware)
	asset_routes.RegisterAssetRoutes(mainGroup)
	dev_tools_routes.RegisterDevToolsRoutes(mainGroup)
	users_routes.RegisterAssetRoutes(mainGroup)
	s3_routes.RegisterS3Routes(mainGroup)
}
