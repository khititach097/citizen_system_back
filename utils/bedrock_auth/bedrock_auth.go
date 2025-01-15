package bedrock_auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Config struct {
	Host         string
	ClientID     string
	ClientSecret string
	UserType     string
}

// Profile structure contains user profile details.
type Profile struct {
	ID          string      `json:"id"`
	Title       string      `json:"title"`
	FirstName   string      `json:"firstName"`
	LastName    string      `json:"lastName"`
	PhoneNumber string      `json:"phoneNumber"`
	Picture     interface{} `json:"picture"`
	Accounts    []Account   `json:"accounts"`
	Status      string      `json:"status"`
}

// Account represents a user account linked to the profile.
type Account struct {
	Provider string `json:"provider"`
	Type     string `json:"type"`
	Account  string `json:"account"`
	Status   string `json:"status"`
}

// AuthResponse represents the response from the auth service.
type AuthResponse struct {
	Profile Profile `json:"profile"`
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
func GetProfile(cookieHeader string) (interface{}, error) {
	config := ConnectionV3()

	// Construct the URL for the auth service API
	url := fmt.Sprintf("%s/api/v1/ums/profile?client_id=%s&client_secret=%s&user_type=%s",
		config.Host, config.ClientID, config.ClientSecret, config.UserType)

	// Create HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Cookie", cookieHeader)

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to auth service: %w", err)
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("auth service returned error code: %d", resp.StatusCode)
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	// Parse the JSON response into a map
	var jsonResp map[string]interface{}
	err = json.Unmarshal(body, &jsonResp)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Log the JSON response for debugging
	fmt.Println(" ****** get profile jsonResp : ", jsonResp)

	return jsonResp, nil
}
