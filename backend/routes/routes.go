package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github.com/zevjr/senai-projeto-aplicado-I/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/api/health", handlers.GetHealth)

	// Rotas para usuários
	r.GET("/api/users", handlers.GetUsers)
	r.GET("/api/users/:id", handlers.GetUser)
	r.POST("/api/users", handlers.CreateUser)

	// Rotas para registros
	r.GET("/api/registers", handlers.GetRegisters)
	r.POST("/api/registers", handlers.CreateRegister)

	// Rotas para arquivos
	r.POST("/api/images", handlers.UploadImage)
	r.GET("/api/images/:uid", handlers.GetImage)
	r.GET("/api/images", handlers.GetImages)

	r.POST("/api/audios", handlers.UploadAudio)
	r.GET("/api/audios/:uid", handlers.GetAudio)
	r.GET("/api/audios/", handlers.GetAudios)

	// Swagger
	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
