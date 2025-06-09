package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetHealth godoc
// @Summary      Verificar saúde da API
// @Description  Retorna status OK se API estiver funcionando corretamente
// @Tags         sistema
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /api/health [get]
func GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "API funcionando corretamente",
	})
}
