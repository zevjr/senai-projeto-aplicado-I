import logging
import psycopg2
from psycopg2 import errors
from typing import Optional, Dict


logger = logging.getLogger(__name__)




class DatabaseManager:
    """Gerencia todas as operações com o banco de dados"""

    def __init__(self):
        self._connect()

    def _connect(self):
        """Estabelece conexão com o banco de dados"""
        try:
            self.conn = psycopg2.connect(
                dbname="PROJETO",
                user="postgres",
                password="postgres",
                host="localhost",
                port="5432"
            )
        except Exception as e:
            logger.error(f"Erro ao conectar: {e}")
            raise

    def is_connected(self) -> bool:
        """Verifica se a conexão está ativa"""
        return self.conn is not None and self.conn.closed == 0

    def save_incident(self, text: str, risk_level: int, justification: str, media_type: str) ->Optional[str]:
        try:
            with self.conn.cursor() as cur:
                cur.execute(
                    """
                    INSERT INTO registers
                    (tite, body, risk_scale, local, status)
                    VALUES (%s, %s, %s, %s, %s)
                    RETURNING uid
                    """,
                    (
                        f"Incidente {media_type}",
                        f"{text}\n\\Justificativa: {justification}",
                        risk_level,
                        "Local não especificado",
                        "evaluated"
                    )
                )
                incidente_id = cur.fetchone()
                self.conn.commit()
                return str(incidente_id)
        except Exception as e:
            self.conn.rollback()
            logger.error(f"Erro ao salvar incidente: {e}")
            return None


    def get_risk_distribution(self) -> Dict[int, int]:
        """Obtém a distribuição de níveis de risco"""
        if not self.is_connected():
            raise ConnectionError("Não conectado ao banco de dados")

        try:
            with self.conn.cursor() as cur:
                cur.execute("""
                    SELECT risk_scale, COUNT(*) 
                    FROM registers 
                    GROUP BY risk_scale
                    ORDER BY risk_scale
                """)
                return {row[0]: row[1] for row in cur.fetchall()}
        except errors.Error as e:
            raise RuntimeError(f"Erro ao buscar distribuição de risco: {e}")

    def close(self):
        """Fecha a conexão com o banco de dados"""
        if self.conn and not self.conn.closed:
            self.conn.close()
            logger.info("Conexão fechada com sucesso")


    def __del__(self):
        """Garante o fechamento da conexão"""
        self.close()

    def link_media_to_incident(self, incident_uid: str, media_uid: str, media_type: str) -> bool:
        """Associa mídia (áudio/vídeo/imagem) a um registro de incidente"""
        valid_types = ['audio', 'video', 'image']
        if media_type not in valid_types:
            logger.error(f"Tipo de mídia inválido: {media_type}")
            return False

        try:
            with self.conn.cursor() as cur:
                cur.execute(
                    f"""
                    UPDATE registers
                    SET {media_type}_uid = %s
                    WHERE uid = %s
                    """,
                    (media_uid, incident_uid))
                self.conn.commit()
                return True
        except Exception as e:
            self.conn.rollback()
            logger.error(f"Erro ao vincular mídia: {e}")
            return False

    def save_video_metadata(self, bucket_name: str, bucket_id: str, duration: float, file_size: int) -> Optional[str]:
        try:
            with self.conn.cursor() as cur:
                cur.execute(
                    """
                    INSERT INTO videos 
                    (bucket_name, bucket_id, duration, file_size)
                    VALUES (%s, %s, %s, %s)
                    RETURNING uid
                    """,
                    (bucket_name, bucket_id, duration, file_size)
                )
                video_uid = cur.fetchone()
                self.conn.commit()
                return str(video_uid)
        except Exception as e:
            self.conn.rollback()
            logger.error(f"Erro ao salvar vídeo: {e}")
            return None

    def save_audio_metadata(self, bucket_name: str, bucket_id: str, duration: float, file_size: int):
        """Similar ao save_video_metadata, mas para áudios"""
        ...

    def get_incident_with_media(self, incident_uid: str) -> dict | None:
        """Recupera um incidente com informações de mídia associada"""
        try:
            with self.conn.cursor() as cur:
                cur.execute("""
                    SELECT r.*, a.bucket_id as audio_id, v.bucket_id as video_id, i.bucket_id as image_id
                    FROM registers r
                    LEFT JOIN audios a ON r.audio_uid = a.uid
                    LEFT JOIN videos v ON r.video_uid = v.uid
                    LEFT JOIN images i ON r.image_uid = i.uid
                    WHERE r.uid = %s
                """, (incident_uid,))
                result = cur.fetchone()
                return dict(result) if result else None
        except Exception as e:
            logger.error(f"Erro ao buscar incidente: {e}")
            return None