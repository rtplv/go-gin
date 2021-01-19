package validate

import (
	"app/core/response"
	"strings"
)

// GetValidationErrors extract errors from error type
func GetValidationErrors(err error) []response.ErrorMessage {
	errorMessages := make([]response.ErrorMessage, 0)

	for _, msg := range strings.Split(err.Error(), "\n") {
		errorMessages = append(errorMessages, response.ErrorMessage{
			Message: msg,
		})
	}

	return errorMessages
}