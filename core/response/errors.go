package response

import (
	"app/libs/validate"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Errors struct {
	Errors []ErrorMessage `json:"errors"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

func ValidationError(ctx *gin.Context, err error) {
	logrus.Error(err.Error(), "Client request validation error")

	ctx.JSON(http.StatusBadRequest, Errors{
		Errors: validate.GetValidationMessages(err),
	})
}

func InternalServerError(ctx *gin.Context, err error) {
	logrus.Error(err.Error(), "Internal Server Error")

	ctx.JSON(http.StatusInternalServerError, Errors{
		Errors: []ErrorMessage{
			{ Message: "Внутренняя ошибка сервера." },
			{ Message: err.Error() },
		},
	})
}