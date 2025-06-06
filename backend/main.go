package main

// @title           Senai Projeto Aplicado I API
// @version         1.0
// @description     API para o projeto aplicado do Senai
// @host            localhost:8080
// @BasePath        /

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/zevjr/senai-projeto-aplicado-I/docs"
)

func main() {
	// Carregar variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente")
	}

	// Configurar conexão com o banco de dados
	SetupDB()
	SeedDatabase(DB)
	// Configurar o router Gin
	r := gin.Default()

	r.GET("/api/health", GetHealth)

	// Rotas para usuários
	r.GET("/api/users", GetUsers)
	r.GET("/api/users/:id", GetUser)
	r.POST("/users", CreateUser)

	// Rotas para registros
	r.GET("/api/registers", GetRegisters)
	r.POST("/api/registers", CreateRegister)
	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Iniciar o servidor na porta 8080
	log.Println("Servidor iniciado na porta 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Falha ao iniciar o servidor: %v", err)
	}
}

// GetHealth godoc
// @Summary      Verificar saúde da API
// @Description  Retorna status OK se a API estiver funcionando corretamente
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
}1

// GetUsers godoc
// @Summary      Get all users
// @Description  Retrieves all users from the database
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {array}  User
// @Router       /api/users [get]
func GetUsers(c *gin.Context) {
	var users []User
	if result := DB.Find(&users); result.Error != nil {
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
// @Success      200  {object}  User
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /api/users/{id} [get]
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

// CreateUser godoc
// @Summary      Criar um novo usuário
// @Description  Cria um novo usuário no banco de dados
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      User  true  "Dados do Usuário"
// @Success      201  {object}  User
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/users [post]
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

// GetRegisters godoc
// @Summary      Obter todos os registros
// @Description  Recupera todos os registros do banco de dados
// @Tags         registers
// @Accept       json
// @Produce      json
// @Success      200  {array}   Register
// @Failure      500  {object}  map[string]string
// @Router       /api/registers [get]
func GetRegisters(c *gin.Context) {
	var registers []Register
	if result := DB.Find(&registers); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, registers)
}

// CreateRegister godoc
// @Summary      Criar um novo registro
// @Description  Cria um novo registro no banco de dados
// @Tags         registers
// @Accept       json
// @Produce      json
// @Param        register  body      Register  true  "Dados do Registro"
// @Success      201  {object}  Register
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/registers [post]
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
