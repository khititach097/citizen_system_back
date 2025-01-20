package middleware

import (
	users_services "citizen_system_back/modules/users/services"
	"fmt"
	"time"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Middleware function
func LoggingMiddleware(c *gin.Context) {
	start := time.Now()
	c.Next() // Process the request
	fmt.Printf("Request: %s %s | Duration: %v", c.Request.Method, c.Request.URL.Path, time.Since(start))
}

// AuthResponse represents the response structure from the authentication API.
type AuthResponse struct {
	Authenticated bool   `json:"authenticated"`
	Message       string `json:"message"`
}

// AuthMiddleware checks user authentication by sending an API request to the auth service.
func AuthMiddleware(c *gin.Context) {
	// Extract the Authorization header
	// authHeader := c.GetHeader("Authorization")
	// if authHeader == "" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
	// 	c.Abort()
	// 	return
	// }

	// // Attempt to get profile using the provided Authorization header
	// result, err := bedrock_auth.GetProfile(authHeader)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request to auth service"})
	// 	c.Abort()
	// 	return
	// }

	// Process User from auth
	user, errUser := users_services.GetUserByEmail()
	if errUser != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user data"})
		c.Abort()
		return
	}
	fmt.Println("User Profile from middleware:", user)
	// Proceed with the request
	c.Set("profile", user)
	c.Next()
}

// isValidToken is a dummy function to check token validity, replace with real implementation
// func isValidToken(authHeader string) bool {
// 	// Extract token from Authorization header (e.g., "Bearer <token>")
// 	// In this example, we are just checking for a simple placeholder token for demonstration
// 	fmt.Println("auth isValidToken authHeader : ", authHeader)
// 	token := authHeader[len("Bearer "):]
// 	fmt.Println("auth isValidToken token : ", token)
// 	return token != "" // Replace with your token validation logic
// }
