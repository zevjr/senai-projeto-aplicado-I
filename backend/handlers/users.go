package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zevjr/senai-projeto-aplicado-I/database"
	"github.com/zevjr/senai-projeto-aplicado-I/models"
	"net/http"
)

// GetUsers godoc
// @Summary      Get all users
// @Description  Retrieves all users from the database
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {array}  models.User
// @Router       /api/users [get]
func GetUsers(c *gin.Context) {
	var users []models.User
	if result := database.DB.Find(&users); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// GetUser godoc
// @Summary      Obter um usuário específico
// @Description  Recupera um usuário pelo ID do banco de dados
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "ID do Usuário"
// @Success      200  {object}  models.User
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /api/users/{id} [get]
func GetUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var user models.User
	if result := database.DB.First(&user, "uid = ?", id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// CreateUser godoc
// @Summary      Criar um novo usuário
// @Description  Cria um novo usuário no banco de dados
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "Dados do Usuário"
// @Success      201  {object}  models.User
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/users [post]
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Gerar um novo UUID se não for fornecido
	if user.UID == uuid.Nil {
		user.UID = uuid.New()
	}

	if result := database.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
