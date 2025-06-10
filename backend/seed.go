package main

import (
	"github.com/zevjr/senai-projeto-aplicado-I/models"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) {
	log.Println("Iniciando seed do banco de dados...")

	// Verifica se já existem dados
	var userCount int64
	db.Model(&models.User{}).Count(&userCount)
	if userCount > 0 {
		log.Println("Banco de dados já possui dados, pulando seed")
		return
	}

	// Criando usuários
	adminUID := uuid.New()
	operatorUID := uuid.New()

	users := []models.User{
		{
			UID:       adminUID,
			Username:  "admin",
			Role:      "admin",
			CreatedAt: time.Now(),
			Password:  "admin123",
			Email:     "admin@gmail.com",
		},
		{
			UID:       operatorUID,
			Username:  "operador",
			Role:      "operator",
			CreatedAt: time.Now(),
			Password:  "operador123",
			Email:     "operador@gmail.com",
		},
	}

	// Criando riscos
	riskUID := uuid.New()
	riskUIDTwo := uuid.New()
	risks := []models.Risk{
		{
			UID:       riskUID,
			Details:   "Risco de queda",
			CreatedAt: time.Now(),
		},
		{
			UID:       riskUIDTwo,
			Details:   "Risco de corte",
			CreatedAt: time.Now(),
		},
	}

	// Criando registros
	imageUID := uuid.New()
	audioUID := uuid.New()
	registerUID := uuid.New()

	registers := []models.Register{
		{
			UID:       registerUID,
			Title:     "Registro inicial",
			Body:      "Este é um registro de teste criado automaticamente",
			RiskScale: 3,
			Local:     "Área de produção",
			Status:    "aberto",
			ImageUID:  imageUID,
			AudioUID:  audioUID,
			CreatedAt: time.Now(),
		},
	}

	// Criando relações
	userRisks := []models.UserRisk{
		{
			UID:       uuid.New(),
			UserUID:   adminUID,
			RiskUID:   riskUID,
			CreatedAt: time.Now(),
		},
	}

	userRegisters := []models.UserRegister{
		{
			UID:         uuid.New(),
			UserUID:     adminUID,
			RegisterUID: registerUID,
			CreatedAt:   time.Now(),
		},
	}

	// Inserindo dados no banco
	db.Create(&users)
	db.Create(&risks)
	db.Create(&registers)
	db.Create(&userRisks)
	db.Create(&userRegisters)

	log.Println("Seed do banco de dados concluído com sucesso")
}
