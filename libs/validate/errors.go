package validate

import (
	"app/core/response"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func ValidationError(ctx *gin.Context, err error) {
	logrus.Error(err.Error(), "Client request validation error")

	ctx.JSON(http.StatusBadRequest, response.Errors{
		Errors: GetValidationErrors(err),
	})
}
