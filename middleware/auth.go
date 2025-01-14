package middleware

import (
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
	authHeader := c.GetHeader("Authorization")
	fmt.Println("authHeader : ", authHeader)
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		c.Abort()
		return
	}

	if authHeader == "" || !isValidToken(authHeader) {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	// Send a request to the auth service
	_, err := http.NewRequest("GET", "http://localhost:5005/auth", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request to auth service"})
		c.Abort()
		return
	}

	// // Add Authorization header to the request
	// req.Header.Set("Authorization", authHeader)

	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to auth service"})
	// 	c.Abort()
	// 	return
	// }
	// defer resp.Body.Close()

	// // Read and parse the response
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from auth service"})
	// 	c.Abort()
	// 	return
	// }

	// var authResponse AuthResponse
	// if err := json.Unmarshal(body, &authResponse); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response from auth service"})
	// 	c.Abort()
	// 	return
	// }

	// // Check if the user is authenticated
	// if !authResponse.Authenticated {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": authResponse.Message})
	// 	c.Abort()
	// 	return
	// }

	// Proceed with the request
	c.Next()
}

// isValidToken is a dummy function to check token validity, replace with real implementation
func isValidToken(authHeader string) bool {
	// Extract token from Authorization header (e.g., "Bearer <token>")
	// In this example, we are just checking for a simple placeholder token for demonstration
	token := authHeader[len("Bearer "):]
	return token == "your-valid-token" // Replace with your token validation logic
}
