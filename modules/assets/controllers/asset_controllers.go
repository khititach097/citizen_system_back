package asset_controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	asset_services "citizen_system_back/modules/assets/services"
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
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/assets [get]
func ListAssets() gin.HandlerFunc {
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

		assets, err := asset_services.GetAllAssets()
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
func GetAssetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid asset ID", errors.New("asset ID cannot be empty")))
			return
		}

		asset, err := asset_services.GetAssetByID(id)
		if err != nil {
			switch {
			case errors.Is(err, gorm.ErrRecordNotFound):
				c.JSON(http.StatusNotFound, response.NewErrorResponse("Asset not found", err))
			default:
				c.JSON(http.StatusInternalServerError, response.NewErrorResponse("Failed to retrieve assets", err))
			}
			return
		}

		c.JSON(http.StatusOK, response.NewResponse("Asset retrieved successfully", asset))
	}
}

// @Summary Get all asset
// @Description Get all asset with user id
// @Tags assets
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param user_id path string true "User ID (UUID)"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/assets/get_asset_by_user_id/{user_id} [get]
func GetAssetByUserID() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")
		parcelType := c.Query("parcel_type")
		muniCode := c.Query("muni_code")
		pageStr := c.Query("page")
		pageSizeStr := c.Query("page_size")

		// Convert page and page_size to integers, with defaults
		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			page = 1 // Default to page 1 if invalid or not provided
		}

		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil || pageSize < 1 {
			pageSize = 10 // Default to page size 10 if invalid or not provided
		}
		//get query name parcel_type and muni_code
		if userId == "" {
			c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid user_id", errors.New("user_id cannot be empty")))
			return
		}

		asset, err := asset_services.GetAssetByUserID(userId, parcelType, muniCode, page, pageSize)
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
