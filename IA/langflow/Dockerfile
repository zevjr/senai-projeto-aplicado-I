# FROM python:3.11-slim

# WORKDIR /app

# # Instala dependências do sistema
# RUN apt-get update && apt-get install -y git build-essential && rm -rf /var/lib/apt/lists/*

# # Instala Langflow
# RUN pip install --no-cache-dir langflow

# # Copia os fluxos para dentro do container
# COPY . /app/flows

# # Comando de inicialização com importação do fluxo
# # LANGFLOW_IMPORT_PATH pode ser definida externamente via ENV
# ENV LANGFLOW_IMPORT_PATH=/app/flows/langflow.json

# EXPOSE 7860

# CMD ["sh", "-c", "langflow run --import-path $LANGFLOW_IMPORT_PATH"]

# Use the official Langflow image as base
FROM langflowai/langflow:latest

# RUN pip install --no-cache-dir <your-packages-here>

COPY langflow.json /app/config/langflow.json

# OPTIONAL — Set environment variables
ENV LANGFLOW_PORT=7860 \
    LANGFLOW_HOST=0.0.0.0 \
    LANGFLOW_LOG_LEVEL=info \
    LANGFLOW_IMPORT_PATH=/app/config/langflow.json


# Expose Langflow default port
EXPOSE 7860

# Default startup command
CMD ["langflow", "run", "--host", "0.0.0.0", "--port", "7860"]
