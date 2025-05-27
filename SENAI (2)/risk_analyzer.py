from langchain.chains import LLMChain
from langchain.prompts import PromptTemplate
from langchain_openai import OpenAI
from config import Config
from typing import Tuple


class RiskAnalyzer:
    """Classifica incidentes com base no texto fornecido"""

    def __init__(self):
        self.llm = OpenAI(
            temperature=0,
            openai_api_key=Config.get_openai_key(),
            max_retries=3
        )

        self.classification_prompt = PromptTemplate(
            input_variables=["text"],
            template="""
            Analise o seguinte relato de incidente e classifique-o conforme:

            Níveis de Risco:
            1-3: Baixo (procedimentos padrão)
            4-6: Médio (requer atenção gerencial)
            7-10: Alto (ação imediata necessária)

            Texto: {text}

            Retorne APENAS um número entre 1 e 10 seguido de uma breve justificativa (máx. 20 palavras).
            Exemplo: "5 - Queda de equipamento sem feridos"
            """
        )

    def classify_incident(self, text: str) -> tuple[int, str]:
        """Classifica o incidente e retorna nível e justificativa"""
        if not isinstance(text, str) or not text.strip():
            return 5, "Texto inválido"

        try:
            chain = LLMChain(
                llm=self.llm,
                prompt=self.classification_prompt,
                verbose=True
            )

            result = chain.run({"text": text})

            # Processa o resultado
            parts = result.split(" - ", 1)
            risk_level = int(parts[0].strip())
            justification = parts[1].strip() if len(parts) > 1 else "Sem justificativa"

            return self._parse_classification(result)

        except Exception as e:
            logger.error(f"Erro na classificação: {e}")
            return 5, "Erro na análise"  # Valor padrão

    def _parse_classification(self, classification: str) -> Tuple[int, str]:
        """Extrai nível e justificativa da resposta do modelo"""
        parts = classification.split(" - ", 1)
        try:
            level = int(parts[0].strip())
            justification = parts[1].strip() if len(parts) > 1 else "Sem justificativa"
            return level, justification
        except (ValueError, IndexError):
            return 5, classification  # Fallback