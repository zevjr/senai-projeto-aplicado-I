import psycopg2
from psycopg2 import OperationalError


def testar_conexao():
    try:
        conn = psycopg2.connect(
            dbname="seu_banco",
            user="seu_usuario",
            password="sua_senha",
            host="localhost",
            port="5432",
            connect_timeout=5
        )

        print("✅ Conexão bem-sucedida!")

        # Teste adicional para verificar se está realmente funcional
        with conn.cursor() as cur:
            cur.execute("SELECT version();")
            db_version = cur.fetchone()
            print(f"Versão do PostgreSQL: {db_version[0]}")

        conn.close()

    except OperationalError as e:
        print(f"❌ Falha na conexão: {e}")