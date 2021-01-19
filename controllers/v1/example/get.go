package example

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Get godoc
// @Summary Example controller
// @Tags example
// @Produce json
// @Param id path string true "guid джобы"
// @Success 200 {object} models.Example
// @Failure default {object} response.Errors
// @Router /id/{id} [get]
func Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"action": "example",
		"id": c.Param("id"),
	})
}
