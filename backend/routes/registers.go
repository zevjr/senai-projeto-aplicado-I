package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zevjr/senai-projeto-aplicado-I/handlers"
)

func SetUpRegistersRoutes(routes *gin.Engine) {

	routes.GET("/api/registers", handlers.GetRegisters)
	routes.POST("/api/registers", handlers.CreateRegister)
	routes.GET("/api/registers/:uid", handlers.GetRegister)
	routes.PUT("/api/registers/:uid", handlers.UpdateRegister)
	routes.DELETE("/api/registers/:uid", handlers.DeleteRegister)

}
