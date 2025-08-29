package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zevjr/senai-projeto-aplicado-I/database"
	"github.com/zevjr/senai-projeto-aplicado-I/models"
	"net/http"
	"time"
)

// CreateRegister godoc
// @Summary      Criar um novo registro
// @Description  Cria um novo registro no banco de dados
// @Tags         registers
// @Accept       json
// @Produce      json
// @Param        register  body      models.Register  true  "Dados do Registro"
// @Success      201  {object}  models.Register
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/registers [post]
func CreateRegister(c *gin.Context) {
	var register models.Register
	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	register.UID = uuid.New()
	register.CreatedAt = time.Now()

	//register.Body
	//database.DB.Find(&models.Image{}, "uid = ?", register.ImageUID)
	// criar uma go routine para trasncrever audio e enviar para IA para analise de risco
	//TODO: IA

	if result := database.DB.Create(&register); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, register)
}

// GetRegisters godoc
// @Summary      Obter todos os registros
// @Description  Recupera todos os registros do banco de dados
// @Tags         registers
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Register
// @Failure      500  {object}  map[string]string
// @Router       /api/registers [get]
func GetRegisters(c *gin.Context) {
	var registers []models.Register
	if result := database.DB.Find(&registers); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, registers)
}

// GetUser godoc
// @Summary      Obter um registro específico
// @Description  Recupera um registro pelo ID do banco de dados
// @Tags         registers
// @Accept       json
// @Produce      json
// @Param        uid   path      string  true  "ID do Registro"
// @Success      200  {object}  models.Register
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /api/registers/{uid} [get]
func GetRegister(c *gin.Context) {
	id, err := uuid.Parse(c.Param("uid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var register models.Register

	if result := database.DB.First(&register, "uid = ?", id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Registro não encontrado"})
		return
	}

	c.JSON(http.StatusOK, register)
}
