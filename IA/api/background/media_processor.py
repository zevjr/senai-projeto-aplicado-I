import whisper
import tempfile
import os
from typing import Optional, Tuple
from config import Config


class MediaProcessor:
    """Processa diferentes tipos de mídia e extrai texto"""

    def __init__(self):
        self._load_model()

    def _load_model(self):
        """Carrega o modelo Whisper apenas quando necessário"""
        if self.model is None:
            model_name = Config.get_whisper_model()
            self.model = whisper.load_model(model_name)

    def audio_to_text(self, audio_path: str) -> Optional[str]:
        """Converte áudio em texto usando Whisper"""
        try:
            result = self.model.transcribe(audio_path)
            return result["text"]
        except Exception as e:
            print(f"Erro ao transcrever áudio: {e}")
            return None

    def video_to_text(self, video_path: str) -> Optional[str]:
        """Extrai áudio de vídeo e converte em texto"""
        try:
            with tempfile.NamedTemporaryFile(suffix=".mp3", delete=False) as tmp_audio:
                video = VideoFileClip(video_path)
                video.audio.write_audiofile(tmp_audio.name)
                text = self.audio_to_text(tmp_audio.name)
                video.close()
                os.unlink(tmp_audio.name)
            return text
        except Exception as e:
            print(f"Erro ao processar vídeo: {e}")
            return None

    def process_audio(self, audio_path: str) -> Tuple[Optional[str], int, str]:
        """Processa áudio com armazenamento de metadados"""
        try:
            # 1. Processa o áudio e obtém texto + metadados
            text, audio_metadata = self.media_processor.process_audio(audio_path)
            if not text:
                return None, 5, "Falha na transcrição do áudio"

            # 2. Classifica o risco
            risk_level, justification = self.risk_analyzer.classify_incident(text)

            # 3. Salva o incidente básico
            incident_id = self.db_manager.save_incident(
                text=text,
                risk_level=risk_level,
                justification=justification,
                media_type="audio"
            )
            if not incident_id:
                return None, risk_level, justification

            # 4. Armazena metadados do áudio (NOVO)
            audio_uid = self.db_manager.save_audio_metadata(
                bucket_name="audios",
                bucket_id=audio_metadata['bucket_id'],
                duration=audio_metadata['duration'],
                file_size=audio_metadata['size']
            )

            if audio_uid:
                self.db_manager.link_media_to_incident(
                    incident_uid=incident_id,
                    media_uid=audio_uid,
                    media_type="audio"
                )

            return incident_id, risk_level, justification

        except Exception as e:
            logger.error(f"Erro no processamento de áudio: {e}")
            return None, 5, str(e)

    def process_video(self, video_path: str) -> tuple[Optional[str], Optional[dict]]:
        """Retorna (texto_transcrito, metadados_video)"""
        try:
            # 1. Extrai áudio do vídeo
            with tempfile.NamedTemporaryFile(suffix=".mp3", delete=False) as tmp_audio:
                video = VideoFileClip(video_path)
                video.audio.write_audiofile(tmp_audio.name)
                video.close()

                # 2. Transcreve áudio
                text = self.audio_to_text(tmp_audio.name)

                # 3. Prepara metadados
                metadata = {
                    'original_path': video_path,
                    'duration': video.duration,
                    'size': os.path.getsize(video_path),
                    'bucket_id': f"vid_{os.path.basename(video_path)}"  # Exemplo
                }
                return text, metadata
        finally:
            os.unlink(tmp_audio.name)