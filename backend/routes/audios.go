package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zevjr/senai-projeto-aplicado-I/handlers"
)

func SetUpAudiosRoutes(routes *gin.Engine) {

	routes.POST("/api/audios", handlers.UploadAudio)
	routes.GET("/api/audios/:uid", handlers.GetAudio)
	routes.GET("/api/audios/", handlers.GetAudios)

}
