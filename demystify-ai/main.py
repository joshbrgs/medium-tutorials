from langchain_community.document_loaders import PyPDFLoader
from langchain_text_splitters import RecursiveCharacterTextSplitter
from langchain_chroma import Chroma
from langchain_ollama import OllamaEmbeddings

file_path = "./robinhood-2023-annual-report.pdf"
loader = PyPDFLoader(file_path)

docs = loader.load()

text_splitter = RecursiveCharacterTextSplitter(chunk_size=1000, chunk_overlap=200)
splits = text_splitter.split_documents(docs)

embeddings = OllamaEmbeddings(model="llama3")

vectorstore = Chroma.from_documents(documents=splits, embedding=embeddings)

query = "What is the 2023 net profit?"
search = vectorstore.similarity_search(query)

print(search[0].page_content)
