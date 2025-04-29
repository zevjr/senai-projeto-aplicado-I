package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDB() {
	// Configuração da conexão com o banco de dados
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Conexão com o banco de dados
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
	}

	// Automigração das tabelas
	err = DB.AutoMigrate(
		&User{},
		&Risk{},
		&UserRisk{},
		&Register{},
		&UserRegister{},
		&Image{},
		&Audio{},
		&Preference{},
		&Configuration{},
	)
	if err != nil {
		log.Fatalf("Falha na migração do banco de dados: %v", err)
	}

	log.Println("Conexão com o banco de dados estabelecida e migração concluída")
}
