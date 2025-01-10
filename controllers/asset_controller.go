package controllers

import (
	"net/http"
	"your-project/services"
	"github.com/gin-gonic/gin"
)

// GetAssets handles the GET request for fetching all assets
func GetAssets(c *gin.Context) {
	assets, err := services.GetAllAssets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, assets)
}

// CreateAsset handles the POST request for creating a new asset
func CreateAsset(c *gin.Context) {
	var asset services.Asset
	if err := c.ShouldBindJSON(&asset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdAsset, err := services.CreateAsset(asset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdAsset)
}
