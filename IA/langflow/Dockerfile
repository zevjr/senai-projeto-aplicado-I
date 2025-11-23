FROM python:3.11-slim

WORKDIR /app

# Instala dependências do sistema
RUN apt-get update && apt-get install -y git build-essential && rm -rf /var/lib/apt/lists/*

# Instala Langflow
RUN pip install --no-cache-dir langflow

# Copia os fluxos para dentro do container
COPY . /app/flows

# Comando de inicialização com importação do fluxo
# LANGFLOW_IMPORT_PATH pode ser definida externamente via ENV
ENV LANGFLOW_IMPORT_PATH=/app/flows/langflow.json

EXPOSE 7860

CMD ["sh", "-c", "langflow run --import-path $LANGFLOW_IMPORT_PATH"]

