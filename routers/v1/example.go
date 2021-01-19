package v1

import (
	"app/controllers/v1/example"
	"github.com/gin-gonic/gin"
)

func setExampleRoutes(r *gin.RouterGroup) {
	taskRouter := r.Group("/example")
	{
		taskRouter.GET("/:guid", example.Get)
	}
}