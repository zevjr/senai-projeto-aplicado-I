# PASSO A PASSO CONEXÃO IA COM PYTHON E LANGCHAIN

## Arquitetura do Sistema

**Diagrama:**  
![Diagrama do Sistema](media/image1.png)

**Código:**  
![Código](media/image2.png)

---

## 1. Recebimento do Registro

- **Input (Registro)**: O colaborador envia dados via:
  - **Texto**: Formulário direto.
  - **Áudio/Vídeo**: Arquivos (ex: `.mp3`, `.mp4`).

![Exemplo de Input](media/image3.png)

---

## 2. Pré-Processamento de Mídia

- **Áudio/Vídeo → Texto**: Usamos bibliotecas como:
  - `whisper` (OpenAI) para áudio.
  - `moviepy` + `whisper` para vídeo.

![Processamento de Mídia](media/image4.png)

---

## 3. Classificação de Risco com LangChain

- **Prompt Template**: Direciona a IA para classificar com base em dados históricos.
- **Conexão com BD**: Contexto armazenado em tabelas SQL.

![Classificação de Risco](media/image5.png)

---

## 4. Armazenamento no Banco de Dados

- **Estrutura da Tabela**:
  ![Estrutura da Tabela](media/image6.png)

- **Inserção via Python**:
  ![Código de Inserção](media/image7.png)

---

## 5. Geração de Gráficos

- **Bibliotecas**: `matplotlib` ou `plotly`.
- **Dados**: Contagem de registros por nível de risco.

![Exemplo de Gráfico](media/image8.png)

---

## Fluxo de Execução Completo

![Fluxo de Execução](media/image9.png)

---

## Tecnologias-Chave

| Função                      | Tecnologias               |
|-----------------------------|---------------------------|
| Processamento de Áudio/Vídeo | `whisper`, `moviepy`      |
| IA e Classificação          | LangChain + OpenAI (ou Llama2) |
| Banco de Dados              | SQLite/PostgreSQL         |
| Visualização                | Matplotlib/Plotly         |

---

## Saída Esperada

1. **Terminal**:  
   ![Saída no Terminal](media/image10.png)

2. **Gráfico**:  
   ![Gráfico de Risco](media/image8.png)