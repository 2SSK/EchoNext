package errs

import "encoding/json"

// APIError represents a structured API error
type APIError struct {
	Status  int         `json:"status"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

// Error implements the error interface
func (e *APIError) Error() string {
	return e.Message
}

// ToJSON serializes the error to JSON
func (e *APIError) ToJSON() ([]byte, error) {
	return json.Marshal(e)
}

// NewAPIError creates a new APIError
func NewAPIError(status int, code, message string, details ...interface{}) *APIError {
	err := &APIError{
		Status:  status,
		Code:    code,
		Message: message,
	}
	if len(details) > 0 {
		err.Details = details[0]
	}
	return err
}
