package s3_routes

import (
	s3_controllers "citizen_system_back/modules/s3/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterAssetRoutes registers routes for asset-related operations
// @Security BearerAuth
// @Security CookieAuth
func RegisterS3Routes(router *gin.RouterGroup) {

	s3Group := router.Group("/s3") // Corrected variable name
	{
		s3Group.GET("/preview/:img_id", s3_controllers.PreviewImageByID())
	}
}
