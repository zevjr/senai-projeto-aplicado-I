package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente")
	}

	// Configurar conexão com o banco de dados
	SetupDB()

	// Configurar o router Gin
	r := gin.Default()

	// Rota para verificar se a API está funcionando
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "API funcionando corretamente",
		})
	})

	// Rotas para usuários
	r.GET("/users", GetUsers)
	r.GET("/users/:id", GetUser)
	r.POST("/users", CreateUser)

	// Rotas para registros
	r.GET("/registers", GetRegisters)
	r.POST("/registers", CreateRegister)

	// Iniciar o servidor na porta 8080
	log.Println("Servidor iniciado na porta 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Falha ao iniciar o servidor: %v", err)
	}
}

// Handlers para usuários
func GetUsers(c *gin.Context) {
	var users []User
	if result := DB.Find(&users); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var user User
	if result := DB.First(&user, "uid = ?", id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Gerar um novo UUID se não for fornecido
	if user.UID == uuid.Nil {
		user.UID = uuid.New()
	}

	if result := DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Handlers para registros
func GetRegisters(c *gin.Context) {
	var registers []Register
	if result := DB.Find(&registers); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, registers)
}

func CreateRegister(c *gin.Context) {
	var register Register
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

	if result := DB.Create(&register); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, register)
}
