// asset_controllers/asset_controller.go
package asset_controllers

import (
	"citizen_system_back/services/asset_services"
	"net/http"
	"citizen_system_back/models"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ErrorResponse represents a structured error response
type ErrorResponse struct {
	Message string `json:"message" example:"Error message"`
	Error   string `json:"error" example:"Detailed error description"`
}

// AssetsResponse is the response structure for listing assets
type AssetsResponse struct {
	Message string          `json:"message" example:"List all assets"`
	Data    []models.AsLand `json:"data"`
}

// GetAssetResponse is the response structure for retrieving a single asset
type GetAssetResponse struct {
	Message string         `json:"message" example:"Asset retrieved successfully"`
	Data    *models.AsLand `json:"data"`
}

// @Summary List all assets
	// @Description Get a list of all assets
	// @Tags assets
	// @Accept json
	// @Produce json
	// @Success 200 {object} AssetsResponse
	// @Failure 500 {object} ErrorResponse
	// @Router /api/v1/assets [get]
func ListAssets(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
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
	}
}

// @Summary Get asset by ID
	// @Description Retrieve an asset by its ID
	// @Tags assets
	// @Accept json
	// @Produce json
	// @Param id path string true "Asset ID (UUID)"
	// @Success 200 {object} GetAssetResponse
	// @Failure 404 {object} ErrorResponse
	// @Failure 500 {object} ErrorResponse
	// @Router /api/v1/assets/{id} [get]
func GetAssetByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		
		asset, err := asset_services.GetAssetByID(db, id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, ErrorResponse{
					Message: "Asset not found",
					Error:   err.Error(),
				})
				return
			}
	
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Message: "Failed to retrieve asset",
				Error:   err.Error(),
			})
			return
		}
	
		c.JSON(http.StatusOK, GetAssetResponse{
			Message: "Asset retrieved successfully",
			Data:    asset,
		})
	}
}