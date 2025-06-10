package handlers

import (
	"github.com/zevjr/senai-projeto-aplicado-I/dto"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/zevjr/senai-projeto-aplicado-I/database"
	"github.com/zevjr/senai-projeto-aplicado-I/models"
)

// UploadImage godoc
// @Summary      Carrega uma imagem
// @Description  Salva uma imagem no banco de dados
// @Tags         images
// @Accept       multipart/form-data
// @Produce      json
// @Param        file  formData  file  true  "Image to upload"
// @Success      201  {object}  models.Image
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/images [post]
func UploadImage(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot open file"})
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot read file"})
		return
	}
	img := models.Image{
		UID:       uuid.New(),
		Name:      fileHeader.Filename,
		MimeType:  fileHeader.Header.Get("Content-Type"),
		Data:      data,
		CreatedAt: time.Now(),
	}
	if result := database.DB.Create(&img); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, img)
}

// GetImage godoc
// @Summary      Baixa uma imagem
// @Description  Recupera uma imagem pelo UID do banco de dados
// @Tags         images
// @Produce      application/octet-stream
// @Param        uid  path      string  true  "Image UID"
// @Success      200  {file}    file
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /api/images/{uid} [get]
func GetImage(c *gin.Context) {
	uid, err := uuid.Parse(c.Param("uid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UID"})
		return
	}
	var img models.Image
	if result := database.DB.First(&img, "uid = ?", uid); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}
	c.Header("Content-Disposition", "attachment; filename=\""+img.Name+"\"")
	c.Header("Content-Type", img.MimeType)
	c.Data(http.StatusOK, img.MimeType, img.Data)
}

// GetImages godoc
// @Summary      Retorna todas as imagens
// @Description  Retorna todas as imagens sem os dados do arquivo
// @Tags         images
// @Produce      json
// @Success      200  {array}  dto.ImageWithoutData
// @Failure      500  {object}  map[string]string
// @Router       /api/images [get]
func GetImages(c *gin.Context) {
	var images []dto.ImageWithoutData
	if result := database.DB.Model(&models.Image{}).
		Select("uid, name, mime_type, created_at").
		Find(&images); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, images)
}

// GetImages godoc
// @Summary      Retorna todos os audios
// @Description  Retorna todos os audios sem os dados do arquivo
// @Tags         audios
// @Produce      json
// @Success      200  {array}  dto.AudioWithoutData
// @Failure      500  {object}  map[string]string
// @Router       /api/audios [get]
func GetAudios(c *gin.Context) {
	var audios []dto.AudioWithoutData
	if result := database.DB.Model(&models.Audio{}).
		Select("uid, name, mime_type, created_at").
		Find(&audios); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
	}

	c.JSON(http.StatusOK, audios)
}

// UploadAudio godoc
// @Summary      Carrega um arquivo de áudio
// @Description  Salva um arquivo de áudio no banco de dados
// @Tags         audios
// @Accept       multipart/form-data
// @Produce      json
// @Param        file  formData  file  true  "Audio to upload"
// @Success      201  {object}  models.Audio
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/audios [post]
func UploadAudio(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot open file"})
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot read file"})
		return
	}
	audio := models.Audio{
		UID:       uuid.New(),
		Name:      fileHeader.Filename,
		MimeType:  fileHeader.Header.Get("Content-Type"),
		Data:      data,
		CreatedAt: time.Now(),
	}
	if result := database.DB.Create(&audio); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, audio)
}

// GetAudio godoc
// @Summary      Baixa um arquivo de áudio
// @Description  Recupera um arquivo de áudio pelo UID do banco de dados
// @Tags         audios
// @Produce      application/octet-stream
// @Param        uid  path      string  true  "Audio UID"
// @Success      200  {file}    file
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /api/audios/{uid} [get]
func GetAudio(c *gin.Context) {
	uid, err := uuid.Parse(c.Param("uid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UID"})
		return
	}
	var audio models.Audio
	if result := database.DB.First(&audio, "uid = ?", uid); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Audio not found"})
		return
	}
	c.Header("Content-Disposition", "attachment; filename=\""+audio.Name+"\"")
	c.Header("Content-Type", audio.MimeType)
	c.Data(http.StatusOK, audio.MimeType, audio.Data)
}
