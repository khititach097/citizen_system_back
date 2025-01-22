package bedrock_auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	asset_types "citizen_system_back/modules/users/types"
)

type Config struct {
	Host         string
	ClientID     string
	ClientSecret string
	UserType     string
}

// AuthResponse represents the response from the auth service.
type AuthResponse struct {
	Profile asset_types.Profile `json:"data"`
}

// ConnectionV3 returns configuration for the auth service.
func ConnectionV3() Config {
	return Config{
		Host:         os.Getenv("AUTH_HOST"),
		ClientID:     os.Getenv("AUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH_CLIENT_SECRET"),
		UserType:     "officer",
	}
}

// GetProfile fetches the user profile using the provided authorization cookie header.
func GetProfile(cookieHeader string) (asset_types.Profile, error) {
	config := ConnectionV3()

	// Construct the URL for the auth service API
	url := fmt.Sprintf("%s/api/v1/ums/profile?client_id=%s&client_secret=%s&user_type=%s",
		config.Host, config.ClientID, config.ClientSecret, config.UserType)

	// Create HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return asset_types.Profile{}, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Cookie", cookieHeader)

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return asset_types.Profile{}, fmt.Errorf("failed to connect to auth service: %w", err)
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		return asset_types.Profile{}, fmt.Errorf("auth service returned error code: %d", resp.StatusCode)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return asset_types.Profile{}, fmt.Errorf("failed to read response: %w", err)
	}
	bodyStr := string(body)
	fmt.Println("****** Get Profile Body: ", bodyStr)

	// Parse the JSON response into a map
	var authResp AuthResponse
	err = json.Unmarshal(body, &authResp)
	if err != nil {
		fmt.Println(" ****** failed to parse JSON : ", err)
		return asset_types.Profile{}, fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Log the JSON response for debugging
	fmt.Println(" ****** get profile profile : ", authResp.Profile)

	return authResp.Profile, nil
}
