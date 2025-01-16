package middleware

import (
	"citizen_system_back/utils/bedrock_auth"
	"fmt"
	"reflect"
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
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		c.Abort()
		return
	}

	// Attempt to get profile using the provided Authorization header
	result, err := bedrock_auth.GetProfile(authHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request to auth service"})
		c.Abort()
		return
	}

	fmt.Println(" ******* result : ", result)
	fmt.Println(" ******* result reflect: ", reflect.TypeOf(result))

	// // Safely extract the 'data' field from the response
	// jsonResult, ok := result.(map[string]interface{})
	// if !ok {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Missing or invalid 'data' field in response"})
	// 	return
	// }

	// Log extracted data for debugging
	// fmt.Println(" ****** result : ", jsonResult["data"])

	// url := os.Getenv("AUTH_HOST") + "/api/v1/ums/profile?client_id=" + os.Getenv("AUTH_CLIENT_ID") + "&user_type=officer" + "&client_secret=" + os.Getenv("AUTH_CLIENT_SECRET")
	// fmt.Println("**************** auth url : ", url)
	// req, err := http.NewRequest("GET", url, nil)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request to auth service"})
	// 	c.Abort()
	// 	return
	// }

	// // Add Authorization header to the request
	// req.Header.Set("Cookie", authHeader)
	// fmt.Println("**************** auth req : ", req)

	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to auth service"})
	// 	c.Abort()
	// 	return
	// }
	// defer resp.Body.Close()

	// fmt.Println("**************** auth resp : ", resp)

	// // Read the response body
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from auth service"})
	// 	c.Abort()
	// 	return
	// }

	// // Log the response body
	// fmt.Println("**************** auth service response body:", string(body))

	// // Unmarshal the JSON into the struct
	// var authResponse AuthResponse
	// if err := json.Unmarshal(body, &authResponse); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON response from auth service"})
	// 	c.Abort()
	// 	return
	// }

	// // Parse the JSON response
	// var jsonResp map[string]interface{}
	// err = json.Unmarshal(body, &jsonResp)
	// if err != nil {
	// 	fmt.Println("Failed to parse JSON:", err)
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse JSON from auth service"})
	// 	c.Abort()
	// 	return
	// }

	// fmt.Printf("**************** auth parsed resp: %+v\n", jsonResp)

	// // Log the parsed response
	// fmt.Printf("Parsed response: %+v\n", authResponse)

	// // Handle the response
	// if resp.StatusCode != http.StatusOK {
	// 	c.JSON(resp.StatusCode, gin.H{"error": "Failed to authenticate with auth service"})
	// 	return
	// }

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
// func isValidToken(authHeader string) bool {
// 	// Extract token from Authorization header (e.g., "Bearer <token>")
// 	// In this example, we are just checking for a simple placeholder token for demonstration
// 	fmt.Println("auth isValidToken authHeader : ", authHeader)
// 	token := authHeader[len("Bearer "):]
// 	fmt.Println("auth isValidToken token : ", token)
// 	return token != "" // Replace with your token validation logic
// }
