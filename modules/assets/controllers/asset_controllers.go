package asset_controllers

import (
	"errors"
	"fmt"
	"net/http"

	"citizen_system_back/database"
	"citizen_system_back/models"
	asset_services "citizen_system_back/modules/assets/services"
	asset_types "citizen_system_back/modules/assets/types"
	"citizen_system_back/utils/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// @Summary List all assets
// @Description Get a list of all assets with optional filtering
// @Tags assets
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/assets [get]
func ListAssets() gin.HandlerFunc {
	return func(c *gin.Context) {

		// profile, exists := c.Get("profile")
		// if !exists {
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Profile information is missing"})
		// 	return
		// }
		// // Use the profile information
		// fmt.Println("Profile from middleware:", profile)

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
// @Security BearerAuth
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
// @Param user_id path string true "User ID"
// @Param page query int false "Page number"
// @Param page_size query int false "Number of items per page"
// @Param parcel_type query int false "Type of parcel"
// @Param muni_code query string false "Municipality code"
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
		// page, err := strconv.Atoi(pageStr)
		// if err != nil || page < 1 {
		// 	page = 1 // Default to page 1 if invalid or not provided
		// }

		// pageSize, err := strconv.Atoi(pageSizeStr)
		// if err != nil || pageSize < 1 {
		// 	pageSize = 10 // Default to page size 10 if invalid or not provided
		// }
		//get query name parcel_type and muni_code
		if userId == "" {
			c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid user_id", errors.New("user_id cannot be empty")))
			return
		}

		query := asset_types.QueryParams{
			Page:       pageStr,
			PageSize:   pageSizeStr,
			UserID:     userId,
			MuniCode:   muniCode,
			ParcelType: parcelType,
		}

		asset, err := asset_services.GetAssetByUserID(query)
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

// @Summary Get all asset
// @Description Get all asset with land id
// @Tags assets
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param asset_id path string true "Land ID or Condo ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/assets/get_asset_by_asset_id/{asset_id} [get]
func GetAssetByLandID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var db = database.GetDB()
		assetId := c.Param("asset_id")
		if assetId == "" {
			c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid asset id", errors.New("asset id cannot be empty")))
			return
		}

		// Parse assetId into UUID
		assetUUID, err := uuid.Parse(assetId)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid asset id format", err))
			return
		}

		// Check if asset exists in AsLands
		var landCount int64
		isExistOnLand := false
		err = db.Model(&models.AsLands{}).Where("id = ?", assetUUID).Count(&landCount).Error
		if err == nil && landCount > 0 {
			isExistOnLand = true
		}

		// Check if asset exists in AsCondos
		var condoCount int64
		isExistOnCondo := false
		err = db.Model(&models.AsCondos{}).Where("id = ?", assetUUID).Count(&condoCount).Error
		if err == nil && condoCount > 0 {
			isExistOnCondo = true
		}

		var asset map[string]interface{}

		if isExistOnLand {
			asset, err = asset_services.GetAssetByLandId(assetId)
		} else if isExistOnCondo {
			asset, err = asset_services.GetAssetByCondoId(assetId)
		} else {
			c.JSON(http.StatusBadRequest, response.NewErrorResponse("Invalid asset id", errors.New("Not found asset on land and condo")))
			return
		}

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
