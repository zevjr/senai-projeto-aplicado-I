import os
from dotenv import load_dotenv
from typing import Optional

load_dotenv()


class Config:
    """Centraliza todas as configurações do sistema"""

    @staticmethod
    def get_db_config() -> dict:
        """Retorna configuração do banco de dados"""
        return {
            'dbname': os.getenv('DB_NAME'),
            'user': os.getenv('DB_USER'),
            'password': os.getenv('DB_PASSWORD'),
            'host': os.getenv('DB_HOST'),
            'port': os.getenv('DB_PORT')
        }

    @staticmethod
    def get_openai_key() -> str:
        """Retorna chave da API OpenAI"""
        key = os.getenv('OPENAI_API_KEY')
        if not key:
            raise ValueError("OPENAI_API_KEY não encontrada nas variáveis de ambiente")
        return key

    @staticmethod
    def get_whisper_model() -> str:
        """Retorna modelo Whisper a ser usado"""
        return os.getenv('WHISPER_MODEL', 'base')  # Default para 'base'