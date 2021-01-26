package response

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type Errors struct {
	Errors []ErrorMessage `json:"errors"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

func ValidationError(ctx *gin.Context, err error) {
	errorMessages := make([]ErrorMessage, 0)
	logrus.Error(err.Error(), "Client request validation error")

	for _, msg := range strings.Split(err.Error(), "\n") {
		errorMessages = append(errorMessages, ErrorMessage{
			Message: msg,
		})
	}

	ctx.JSON(http.StatusBadRequest, Errors{
		Errors: errorMessages,
	})
}

func InternalServerError(ctx *gin.Context, err error) {
	logrus.Error(err.Error(), "Internal Server Error")

	ctx.JSON(http.StatusInternalServerError, Errors{
		Errors: []ErrorMessage{
			{Message: "Internal Server Error."},
			{Message: err.Error()},
		},
	})
}
