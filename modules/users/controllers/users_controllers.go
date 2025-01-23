package users_controllers

import (
	"citizen_system_back/models"
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
		fmt.Println(" 32 GetProfile Profile from middleware:", profile)

		// Perform a type assertion to convert profile to models.CitizenUser
		userProfile, ok := profile.(models.CitizenUser)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid profile type"})
			return
		}

		result, err := users_services.GetProfile(userProfile)
		if err != nil {
			c.JSON(http.StatusNotFound, response.NewResponse("user not found", gin.H{
				"profile": nil,
			}))
			return
		}

		c.JSON(http.StatusOK, response.NewResponse("Get user profile successfully", gin.H{
			"profile": result,
		}))
	}
}
