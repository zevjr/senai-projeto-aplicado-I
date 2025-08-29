package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/zevjr/senai-projeto-aplicado-I/handlers"
)

func SetupRouter() *gin.Engine {
	routes := gin.Default()

	SetUpUsersRoutes(routes)
	SetUpRegistersRoutes(routes)
	SetUpImagesRoutes(routes)
	SetUpAudiosRoutes(routes)

	// telemetry
	routes.GET("/api/health", handlers.GetHealth)

	// Swagger
	routes.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return routes
}
