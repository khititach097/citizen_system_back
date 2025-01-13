package response

import (

)


// Response represents the standard API response structure
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message" example:"Operation successful"`
	Error   *string     `json:"error,omitempty" example:"Detailed error description"`
	Data    interface{} `json:"data,omitempty"`
}

// newResponse creates a new successful Response instance
func NewResponse(message string, data interface{}) Response {
	return Response{
		Success: true,
		Message: message,
		Data:    data,
	}
}

// newErrorResponse creates a new error Response instance
func NewErrorResponse(message string, err error) Response {
	errStr := err.Error()
	return Response{
		Success: false,
		Message: message,
		Error:   &errStr,
	}
}