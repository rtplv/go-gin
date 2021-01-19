package response

import (
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

func InternalServerError(ctx *gin.Context, err error) {
	logrus.Error(err.Error(), "Internal Server Error")

	ctx.JSON(http.StatusInternalServerError, Errors{
		Errors: []ErrorMessage{
			{ Message: "Внутренняя ошибка сервера." },
			{ Message: err.Error() },
		},
	})
}