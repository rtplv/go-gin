package example

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Get godoc
// @Summary Example controller
// @Tags example
// @Produce json
// @Param id path string true " "
// @Success 200 {object} models.Example
// @Failure default {object} response.Errors
// @Router /v1/example/{id} [get]
func Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"action": "example/get",
		"id": ctx.Param("id"),
	})
}
