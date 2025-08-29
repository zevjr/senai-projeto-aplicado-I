package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zevjr/senai-projeto-aplicado-I/handlers"
)

func SetUpUsersRoutes(routes *gin.Engine) {

	routes.GET("/api/users", handlers.GetUsers)
	routes.GET("/api/users/:uid", handlers.GetUser)
	routes.POST("/api/users", handlers.CreateUser)

}
