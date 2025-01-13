package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"citizen_system_back/controllers/asset_controller"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	_ "citizen_system_back/docs"
)

// RegisterRoutes registers all application routes.
func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	// Register Asset routes with the database passed as a dependency
	assetGroup := router.Group("/api/v1")
	asset_controllers.RegisterAssetRoutes(assetGroup, db)

	// Swagger endpoint at /api/v1/docs
	assetGroup.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
