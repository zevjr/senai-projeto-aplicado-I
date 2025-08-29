package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zevjr/senai-projeto-aplicado-I/handlers"
)

func SetUpImagesRoutes(routes *gin.Engine) {

	routes.POST("/api/images", handlers.UploadImage)
	routes.GET("/api/images/:uid", handlers.GetImage)
	routes.GET("/api/images", handlers.GetImages)

}