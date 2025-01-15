package asset_controllers

import (
	"citizen_system_back/services/asset_services"
	"errors"
	"fmt"
	"net/http"

	"citizen_system_back/utils/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary List all assets
// @Description Get a list of all assets with optional filtering
// @Tags assets
// @Security CookieAuth
// @Accept json
// @Produce json
// @Param page query int false "Page number (default: 1)"
// @Param limit query int false "Items per page (default: 10)"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/assets [get]
func ListAssets(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse pagination parameters
		page := 1
		limit := 10
		if pageStr := c.DefaultQuery("page", "1"); pageStr != "" {
			if _, err := fmt.Sscanf(pageStr, "%d", &page); err != nil || page < 1 {
				c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid page parameter", errors.New("page must be a positive integer")))
				return
			}
		}
		if limitStr := c.DefaultQuery("limit", "10"); limitStr != "" {
			if _, err := fmt.Sscanf(limitStr, "%d", &limit); err != nil || limit < 1 {
				c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid limit parameter", errors.New("limit must be a positive integer")))
				return
			}
		}

		assets, err := asset_services.GetAllAssets(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to retrieve assets", err))
			return
		}

		c.JSON(http.StatusOK, response.NewResponse("Assets retrieved successfully", gin.H{
			"assets": assets,
		}))
	}
}

// @Summary Get asset by ID
// @Description Retrieve an asset by its ID
// @Tags assets
// @Security CookieAuth
// @Accept json
// @Produce json
// @Param id path string true "Asset ID (UUID)"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/assets/{id} [get]
func GetAssetByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid asset ID", errors.New("asset ID cannot be empty")))
			return
		}

		asset, err := asset_services.GetAssetByID(db, id)
		if err != nil {
			switch {
			case errors.Is(err, gorm.ErrRecordNotFound):
				c.JSON(http.StatusNotFound, response.NewErrorResponse("Asset not found", err))
			default:
				c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to retrieve asset", err))
			}
			return
		}

		c.JSON(http.StatusOK, response.NewResponse("Asset retrieved successfully", asset))
	}
}
