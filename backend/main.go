package main

// @title           Senai Projeto Aplicado I API
// @version         1.0
// @description     API para o projeto aplicado do Senai
// @host            localhost:8080
// @BasePath        /

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/zevjr/senai-projeto-aplicado-I/database"
	_ "github.com/zevjr/senai-projeto-aplicado-I/docs"
	"github.com/zevjr/senai-projeto-aplicado-I/routes"
)

func main() {
	// Carregar variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente")
	}

	// Configurar conexão com o banco de dados
	database.SetupDB()
	//SeedDatabase(database.DB)
	// Configurar o router Gin
	r := routes.SetupRouter()

	// Iniciar o servidor na porta 8080
	log.Println("Servidor iniciado na porta 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Falha ao iniciar o servidor: %v", err)
	}
}
