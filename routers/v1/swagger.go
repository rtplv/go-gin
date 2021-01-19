package v1

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setSwaggerRoutes(r *gin.RouterGroup) {
	url := ginSwagger.URL("/v1/docs/swagger.json")

	r.StaticFile("/docs/swagger.json", "./docs/swagger.json")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
