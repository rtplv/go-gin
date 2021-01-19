package validate

import (
	"app/core/response"
	"strings"
)

// GetValidationMessages extract errors from error type
func GetValidationMessages(err error) []response.ErrorMessage {
	errorMessages := make([]response.ErrorMessage, 0)

	for _, msg := range strings.Split(err.Error(), "\n") {
		errorMessages = append(errorMessages, response.ErrorMessage{
			Message: msg,
		})
	}

	return errorMessages
}