# Development stage (com hot reload)
FROM node:20-alpine

WORKDIR /app

# Copiar arquivos de dependências (incluindo package-lock.json)
COPY package*.json ./

# Instalar dependências
# Usar npm install ao invés de npm ci para evitar problemas com cache
RUN npm install

# Copiar código fonte
COPY . .

# Sobrescrever o arquivo proxy.conf.json para usar o service name do Docker
RUN echo '{ \
  "/api": { \
    "target": "http://backend:8080", \
    "secure": false, \
    "changeOrigin": true, \
    "logLevel": "debug" \
  } \
}' > /app/src/proxy.conf.json

# Expor a porta
EXPOSE 4200

# Comando para executar em modo desenvolvimento com proxy
CMD ["npm", "start", "--", "--host", "0.0.0.0", "--proxy-config", "src/proxy.conf.json"]
