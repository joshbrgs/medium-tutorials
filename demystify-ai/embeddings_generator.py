from langchain_chroma import Chroma
from langchain_ollama import OllamaEmbeddings


class EmbeddingsGenerator:
    def __init__(self, model="llama3"):
        self.model = model

    def generate_embeddings(self, documents):
        embeddings = OllamaEmbeddings(model=self.model)
        vectorstore = Chroma.from_documents(documents=documents, embedding=embeddings)
        return vectorstore.as_retriever()
