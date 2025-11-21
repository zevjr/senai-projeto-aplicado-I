# Build stage
FROM golang:1.24-alpine AS builder

# Instalar dependências necessárias
RUN apk add --no-cache git gcc musl-dev

# Definir diretório de trabalho
WORKDIR /app

# Copiar arquivos de dependências
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar o código fonte (exceto .env que será gerado)
COPY . .

# Build da aplicação
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main .

# Production stage
FROM alpine:latest

# Instalar ca-certificates para HTTPS e wget para healthcheck
RUN apk --no-cache add ca-certificates wget

WORKDIR /root/

# Copiar o binário compilado
COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

# Variáveis de ambiente padrão (podem ser sobrescritas pelo docker-compose)
ENV DB_HOST=postgres \
    DB_USER=postgres \
    DB_PASSWORD=postgres \
    DB_NAME=risk_register \
    DB_PORT=5432 \
    GIN_MODE=release \
    PORT=8080

# Expor a porta
EXPOSE 8080

# Comando para executar
# As variáveis de ambiente são lidas diretamente pela aplicação Go
CMD ["./main"]
