package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zevjr/senai-projeto-aplicado-I/database"
	"github.com/zevjr/senai-projeto-aplicado-I/models"
	"net/http"
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

	// Gerar um novo UUID se não for fornecido
	if register.UID == uuid.Nil {
		register.UID = uuid.New()
	}

	//TODO: Enviar as fotos para o Bucket e armazenar o ID na tabela de imagens
	//TODO: Enviar os audios para o Bucket e armazenar o ID na tabela de imagens
	//TODO: Criar a rotina para consumir os registros e enviar los para a AI

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
