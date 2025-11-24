package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zevjr/senai-projeto-aplicado-I/database"
	"github.com/zevjr/senai-projeto-aplicado-I/models"
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
	go func(reg models.Register) {
    // 1. Monta o input para IA
		inputValue := reg.Title + " " + reg.Body

		payload := map[string]string{
			"message": inputValue,
		}

		payloadBytes, _ := json.Marshal(payload)

		// 2. Faz a requisição POST
		resp, err := http.Post(
			"http://api:5000/ia/call",
			"application/json",
			bytes.NewBuffer(payloadBytes),
		)
		if err != nil {
			log.Printf("Erro ao chamar IA: %v", err)
			return
		}
		defer resp.Body.Close()

		// 3. Lê o número retornado
		var result struct {
			Response string `json:"response"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			log.Printf("Erro ao decodificar resposta da IA: %v", err)
			return
		}

		parts := strings.SplitN(result.Response, " - ", 2)
		if len(parts) != 2 {
				log.Printf("Formato inesperado: %s", result.Response)
				return
		}

		riskValue, err := strconv.Atoi(strings.TrimSpace(parts[0])) 
		if err != nil {
				log.Printf("Erro ao converter o risco, %s", err)
				return
		}
		riskText := strings.TrimSpace(parts[1])
		newBody := reg.Body + 
		"\nIA Response:" + riskText +
	  "\nEscala de risco retificada pela IA"

		// 4. Atualiza o registro no banco
		updateData := map[string]interface{}{
			"risk_scale": riskValue,
			"status":     "Em análise",
			"body":       newBody,
		}

		if err := database.DB.Model(&models.Register{}).Where("uid = ?", reg.UID).Updates(updateData).Error; err != nil {
			log.Printf("Erro ao atualizar registro com IA: %v", err)
			return
		}

		log.Printf("Registro %s atualizado com IA. RiskScale=%d", reg.UID, riskValue)

	}(register)


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

// UpdateRegister godoc
// @Summary      Atualizar um registro
// @Description  Atualiza um registro existente no banco de dados
// @Tags         registers
// @Accept       json
// @Produce      json
// @Param        uid      path      string         true  "ID do Registro"
// @Param        register  body      models.Register  true  "Dados do Registro"
// @Success      200  {object}  models.Register
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/registers/{uid} [put]
func UpdateRegister(c *gin.Context) {
	id, err := uuid.Parse(c.Param("uid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var existingRegister models.Register
	if result := database.DB.First(&existingRegister, "uid = ?", id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Registro não encontrado"})
		return
	}

	var updatedRegister models.Register
	if err := c.ShouldBindJSON(&updatedRegister); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedRegister.UID = id
	updatedRegister.CreatedAt = existingRegister.CreatedAt

	if result := database.DB.Model(&existingRegister).Updates(updatedRegister); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedRegister)
}

// DeleteRegister godoc
// @Summary      Deletar um registro
// @Description  Remove um registro do banco de dados pelo UID
// @Tags         registers
// @Accept       json
// @Produce      json
// @Param        uid   path      string  true  "ID do Registro"
// @Success      204  "No Content"
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/registers/{uid} [delete]
func DeleteRegister(c *gin.Context) {
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

	if result := database.DB.Delete(&register); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
