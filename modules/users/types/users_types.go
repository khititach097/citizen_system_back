package users_types

// Profile structure contains user profile details.
//
//	type Profile struct {
//		ID          string      `json:"id"`
//		Title       string      `json:"title"`
//		FirstName   string      `json:"firstName"`
//		LastName    string      `json:"lastName"`
//		PhoneNumber string      `json:"phoneNumber"`
//		Picture     interface{} `json:"picture"`
//		Accounts    []Account   `json:"accounts"`
//		Status      string      `json:"status"`
//	}
type Profile struct {
	ID           string  `json:"id"`
	Title        string  `json:"title"`
	FirstName    string  `json:"firstName"`
	LastName     string  `json:"lastName"`
	PhoneNumber  string  `json:"phoneNumber"`
	Picture      *string `json:"picture"` // Can be null, so pointer type is used
	DocumentType string  `json:"documentType"`
	NationalID   string  `json:"nationalId"`
	PassportNo   *string `json:"passportNo"` // Can be null, so pointer type is used
	Status       string  `json:"status"`
	// Accounts     []Account `json:"accounts"`
}

// Account represents a user account linked to the profile.
type Account struct {
	Provider string `json:"provider"`
	Type     string `json:"type"`
	Account  string `json:"account"`
	Status   string `json:"status"`
}
