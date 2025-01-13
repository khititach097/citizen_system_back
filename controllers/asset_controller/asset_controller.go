package asset_controllers

import (
	"citizen_system_back/services/asset_services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterAssetRoutes registers routes for asset-related operations
func RegisterAssetRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/assets", func(c *gin.Context) {
		assets, err := asset_services.GetAllAssets(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to retrieve assets",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "List all assets",
			"data":    assets,
		})
	})
}