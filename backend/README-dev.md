# README-dev.md - Guia de Desenvolvimento

## Estrutura do Projeto

Este projeto utiliza **Go** com **Gin** para API REST, **GORM** para ORM e **PostgreSQL** como banco de dados.

```
backend/
├── database/           # Configuração do banco de dados
├── docs/              # Documentação Swagger
├── dto/               # Data Transfer Objects
├── handlers/          # Controladores/Handlers dos endpoints
├── models/            # Modelos de dados (entidades)
├── routes/            # Configuração de rotas
├── main.go            # Ponto de entrada da aplicação
├── go.mod             # Dependências do projeto
└── env.example        # Exemplo de variáveis de ambiente
```

## Pré-requisitos

- Go 1.23+
- PostgreSQL
- Git

## Configuração Inicial

### 1. Configurar Variáveis de Ambiente

Copie o arquivo `env.example` para `.env`:

```bash
cp env.example .env
```

Configure as variáveis no arquivo `.env`:

```env
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=risk_register
DB_PORT=5432
```

### 2. Instalar Dependências

```bash
go mod tidy
```

### 3. Executar a Aplicação

```bash
go run *.go
```

A API estará disponível em `http://localhost:8080`

## Passo a Passo: Criando um Novo CReateUpdateDelete

### Etapa 1: Definir o Modelo (Model)

Primeiro, defina a estrutura de dados no arquivo `models/modelos.go`:

```go
// Exemplo: Modelo para Categoria
type Category struct {
    UID       uuid.UUID `json:"uid" gorm:"type:uuid;primaryKey"`
    Name      string    `json:"name" gorm:"type:varchar(255);not null"`
    Description string  `json:"description" gorm:"type:text"`
    CreatedAt time.Time `json:"created_at" gorm:"type:timestamp"`
    UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp"`
}
```

**Importante:** Adicione o novo modelo na função `AutoMigrate` em `database/database.go`:

```go
err = DB.AutoMigrate(
    &models.User{},
    &models.Risk{},
    &models.Category{}, // Adicione aqui
    // ... outros modelos
)
```

### Etapa 2: Criar o Handler

Crie um novo arquivo `handlers/categories.go`:

```go
package handlers

import (
    "net/http"
    "time"
    
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "github.com/zevjr/senai-projeto-aplicado-I/database"
    "github.com/zevjr/senai-projeto-aplicado-I/models"
)

// GetCategories godoc
// @Summary      Obter todas as categorias
// @Description  Recupera todas as categorias do banco de dados
// @Tags         categories
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Category
// @Failure      500  {object}  map[string]string
// @Router       /api/categories [get]
func GetCategories(c *gin.Context) {
    var categories []models.Category
    if result := database.DB.Find(&categories); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }
    c.JSON(http.StatusOK, categories)
}

// GetCategory godoc
// @Summary      Obter uma categoria específica
// @Description  Recupera uma categoria pelo ID do banco de dados
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        uid   path      string  true  "ID da Categoria"
// @Success      200  {object}  models.Category
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /api/categories/{uid} [get]
func GetCategory(c *gin.Context) {
    id, err := uuid.Parse(c.Param("uid"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    var category models.Category
    if result := database.DB.First(&category, "uid = ?", id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Categoria não encontrada"})
        return
    }

    c.JSON(http.StatusOK, category)
}

// CreateCategory godoc
// @Summary      Criar uma nova categoria
// @Description  Cria uma nova categoria no banco de dados
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        category  body      models.Category  true  "Dados da Categoria"
// @Success      201  {object}  models.Category
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/categories [post]
func CreateCategory(c *gin.Context) {
    var category models.Category
    if err := c.ShouldBindJSON(&category); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Gerar UUID e timestamps
    category.UID = uuid.New()
    category.CreatedAt = time.Now()
    category.UpdatedAt = time.Now()

    if result := database.DB.Create(&category); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusCreated, category)
}

// UpdateCategory godoc
// @Summary      Atualizar uma categoria
// @Description  Atualiza uma categoria existente no banco de dados
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        uid       path      string           true  "ID da Categoria"
// @Param        category  body      models.Category  true  "Dados da Categoria"
// @Success      200  {object}  models.Category
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/categories/{uid} [put]
func UpdateCategory(c *gin.Context) {
    id, err := uuid.Parse(c.Param("uid"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    var category models.Category
    if result := database.DB.First(&category, "uid = ?", id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Categoria não encontrada"})
        return
    }

    var updateData models.Category
    if err := c.ShouldBindJSON(&updateData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Manter o UID original e atualizar timestamp
    updateData.UID = category.UID
    updateData.CreatedAt = category.CreatedAt
    updateData.UpdatedAt = time.Now()

    if result := database.DB.Save(&updateData); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, updateData)
}

// DeleteCategory godoc
// @Summary      Deletar uma categoria
// @Description  Remove uma categoria do banco de dados
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        uid   path      string  true  "ID da Categoria"
// @Success      204  "No Content"
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/categories/{uid} [delete]
func DeleteCategory(c *gin.Context) {
    id, err := uuid.Parse(c.Param("uid"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    var category models.Category
    if result := database.DB.First(&category, "uid = ?", id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Categoria não encontrada"})
        return
    }

    if result := database.DB.Delete(&category); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusNoContent, nil)
}
```

### Etapa 3: Configurar as Rotas

Adicione as rotas no arquivo `routes/routes.go`:

```go
func SetupRouter() *gin.Engine {
    r := gin.Default()

    // Rotas existentes...
    r.GET("/api/health", handlers.GetHealth)

    // Novas rotas para categorias
    r.GET("/api/categories", handlers.GetCategories)
    r.GET("/api/categories/:uid", handlers.GetCategory)
    r.POST("/api/categories", handlers.CreateCategory)
    r.PUT("/api/categories/:uid", handlers.UpdateCategory)
    r.DELETE("/api/categories/:uid", handlers.DeleteCategory)

    // Outras rotas...
    return r
}
```

### Etapa 4: Testar os Endpoints

#### 1. Criar uma categoria (POST)

```bash
curl -X POST http://localhost:8080/api/categories \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Segurança",
    "description": "Categoria para riscos de segurança"
  }'
```

#### 2. Listar todas as categorias (GET)

```bash
curl http://localhost:8080/api/categories
```

#### 3. Obter uma categoria específica (GET)

```bash
curl http://localhost:8080/api/categories/{uid}
```

#### 4. Atualizar uma categoria (PUT)

```bash
curl -X PUT http://localhost:8080/api/categories/{uid} \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Segurança Atualizada",
    "description": "Descrição atualizada"
  }'
```

#### 5. Deletar uma categoria (DELETE)

```bash
curl -X DELETE http://localhost:8080/api/categories/{uid}
```

## Padrões e Boas Práticas

### 1. Estrutura de Resposta

**Sucesso:**
```json
{
  "uid": "123e4567-e89b-12d3-a456-426614174000",
  "name": "Categoria Exemplo",
  "description": "Descrição da categoria",
  "created_at": "2024-01-01T10:00:00Z",
  "updated_at": "2024-01-01T10:00:00Z"
}
```

**Erro:**
```json
{
  "error": "Mensagem de erro descritiva"
}
```

### 2. Códigos de Status HTTP

- `200 OK` - Sucesso na consulta/atualização
- `201 Created` - Recurso criado com sucesso
- `204 No Content` - Recurso deletado com sucesso
- `400 Bad Request` - Dados inválidos
- `404 Not Found` - Recurso não encontrado
- `500 Internal Server Error` - Erro interno do servidor

### 3. Validações

Use tags do GORM para validações básicas:

```go
type Category struct {
    Name string `json:"name" gorm:"type:varchar(255);not null" validate:"required,min=3,max=255"`
    Email string `json:"email" gorm:"type:varchar(255);unique" validate:"email"`
}
```

### 4. Relacionamentos

Para relacionamentos entre tabelas:

```go
// One-to-Many
type Category struct {
    UID   uuid.UUID `json:"uid" gorm:"type:uuid;primaryKey"`
    Name  string    `json:"name"`
    Risks []Risk    `json:"risks" gorm:"foreignKey:CategoryUID"`
}

type Risk struct {
    UID         uuid.UUID `json:"uid" gorm:"type:uuid;primaryKey"`
    CategoryUID uuid.UUID `json:"category_uid" gorm:"type:uuid"`
    Category    Category  `json:"category" gorm:"references:UID"`
}
```

### 5. Paginação

Para endpoints que retornam listas grandes:

```go
func GetCategoriesPaginated(c *gin.Context) {
    page := c.DefaultQuery("page", "1")
    limit := c.DefaultQuery("limit", "10")
    
    pageInt, _ := strconv.Atoi(page)
    limitInt, _ := strconv.Atoi(limit)
    offset := (pageInt - 1) * limitInt
    
    var categories []models.Category
    var total int64
    
    database.DB.Model(&models.Category{}).Count(&total)
    database.DB.Offset(offset).Limit(limitInt).Find(&categories)
    
    c.JSON(http.StatusOK, gin.H{
        "data": categories,
        "total": total,
        "page": pageInt,
        "limit": limitInt,
    })
}
```

## Documentação Swagger

A documentação da API é gerada automaticamente e está disponível em:
`http://localhost:8080/api/swagger/index.html`

### Instalação do Swag CLI

Primeiro, instale a ferramenta `swag`:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### Configurar PATH (Opcional)

Para usar `swag` diretamente sem especificar o caminho completo:

```bash
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc
source ~/.zshrc
```

### Gerar Documentação

Para regenerar a documentação após adicionar novos endpoints:

```bash
# Se configurou o PATH
swag init

# Ou usando o caminho completo
$(go env GOPATH)/bin/swag init
```

### Arquivos Gerados

O comando `swag init` gera automaticamente:
- `docs/docs.go` - Código Go da documentação
- `docs/swagger.json` - Especificação JSON da API
- `docs/swagger.yaml` - Especificação YAML da API

### Acessar a Documentação

 Acesse: http://localhost:8080/api/!swagger/index.html

## Comandos Úteis

### Executar a aplicação
```bash
go run main.go
```