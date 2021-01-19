package example

import (
	"app/core/request"
	"app/core/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Create godoc
// @Summary Example controller
// @Tags example
// @Produce json
// @Param params body request.Example true " "
// @Success 200 {object} models.Example
// @Failure default {object} response.Errors
// @Router /v1/example [post]
func Create(ctx *gin.Context) {
	var req request.Example

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		response.ValidationError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"action": "example/create",
		"body": req,
	})
}
