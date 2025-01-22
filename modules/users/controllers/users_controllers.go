package users_controllers

import (
	users_services "citizen_system_back/modules/users/services"
	"citizen_system_back/utils/response"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get Profile
// @Description Get Citizen User Profile
// @Tags users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/users/profile [get]
func GetProfile() gin.HandlerFunc {
	return func(c *gin.Context) {

		profile, exists := c.Get("profile")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Profile information is missing"})
			return
		}
		// Use the profile information
		fmt.Println("Profile from middleware:", profile)

		result := users_services.GetProfile()

		c.JSON(http.StatusOK, response.NewResponse("Get user profile successfully", gin.H{
			"profile": result,
		}))
	}
}
