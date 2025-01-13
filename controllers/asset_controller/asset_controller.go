package asset_controllers

import (
	"citizen_system_back/services/asset_services"
	"net/http"
	"citizen_system_back/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ErrorResponse represents a structured error response
type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

// AssetsResponse is the response structure for listing assets
type AssetsResponse struct {
	Message string          `json:"message"`
	Data    []models.AsLand `json:"data"`
}

// RegisterAssetRoutes registers routes for asset-related operations
// @Summary List all assets
// @Description Retrieve all assets in the system
// @Tags assets
// @Accept  json
// @Produce  json
// @Success 200 {object} AssetsResponse
// @Failure 500 {object} ErrorResponse "Failed to retrieve assets"
// @Router /api/v1/assets [get]
func RegisterAssetRoutes(router *gin.RouterGroup, db *gorm.DB) {
	router.GET("/assets", func(c *gin.Context) {
		assets, err := asset_services.GetAllAssets(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Message: "Failed to retrieve assets",
				Error:   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, AssetsResponse{
			Message: "List all assets",
			Data:    assets,
		})
	})
}
