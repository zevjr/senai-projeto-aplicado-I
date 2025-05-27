import logging
import sys
from typing import Optional, Tuple

logging.basicConfig(
    level=logging.INFO,
    format='%(asctime)s - %(levelname)s - %(message)s',
    handlers=[
        logging.FileHandler('incident_analysis.log', encoding='utf-8'),
        logging.StreamHandler()
    ]
)
logger = logging.getLogger(__name__)
class IncidentAnalysisSystem:
    def __init__(self):
        self.logger = logging.getLogger(f"{__name__}.IncidentAnalysisSystem")
        try:
            self.media_processor = MediaProcessor()
            self.risk_analyzer = RiskAnalyzer()
            self.db_manager = DatabaseManager()
            self.logger.info("Sistema inicializado com sucesso")
        except Exception as e:
            self.logger.critical(f"Falha na inicialização: {e}")

    def process_text(self, text: str) -> tuple[Optional[int], int, str]:
        """Processa texto com verificação de componentes"""
        try:
            risk_level, justification = self.risk_analyzer.classify_incident(text=text)

            incident_id = self.db_manager.save_incident(
                text=text,
                risk_level=risk_level,
                justification=justification,
                media_type="text"
            )

            return incident_id, risk_level, justification

        except ValueError as ve:
            logger.error(f"Erro de validação: {ve}")
        except Exception as e:
            logger.error(f"Erro ao processar texto: {e}")
            return None, 5, "Erro no processamento"

    def process_video(self, video_path: str) -> Tuple[Optional[str], int, str]:
        """Processa vídeo com armazenamento de metadados"""
        try:
            # 1. Processa o vídeo e obtém texto + metadados
            text, video_metadata = self.media_processor.process_video(video_path)
            if not text:
                return None, 5, "Falha na transcrição do vídeo"

            # 2. Classifica o risco
            risk_level, justification = self.risk_analyzer.classify_incident(text)

            # 3. Salva o incidente básico
            incident_id = self.db_manager.save_incident(
                text=text,
                risk_level=risk_level,
                justification=justification,
                media_type="video"
            )
            if not incident_id:
                return None, risk_level, justification

            # 4. Armazena metadados do vídeo e associa ao incidente (NOVO)
            video_uid = self.db_manager.save_video_metadata(
                bucket_name="videos",
                bucket_id=video_metadata['bucket_id'],
                duration=video_metadata['duration'],
                file_size=video_metadata['size']
            )

            if video_uid:
                self.db_manager.link_media_to_incident(
                    incident_uid=incident_id,
                    media_uid=video_uid,
                    media_type="video"
                )

            return incident_id, risk_level, justification

        except Exception as e:
            logger.error(f"Erro no processamento de vídeo: {e}")
            return None, 5, str(e)

def main() -> int:
    system = None

    try:
        logger.info("Iniciando o sistema de análise de incidentes")

        system = IncidentAnalysisSystem()
        logger.info("Sistema Inicializando com sucesso")

        sample_text = "Vazamento de produto químico"
        logger.debug(f"Processando texto: {sample_text}")

        result = system.process_text(sample_text)
        logger.info(f"Resultado: {result}")
        return 0

    except Exception as e:
        logger.critical(f"Falha catastrófica no sistema", exc_info=True)
        return 1
    finally:
        logger.info("Encerrando o sistema de análise")
        if system is not None and hasattr(system, 'db_manager'):
            try:
                system.db_manager.close()
            except Exception as e:
                logger.error(f"Erro ao fechar conexão: {e}")


if __name__ == "__main__":
    sys.exit(main())
