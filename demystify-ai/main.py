from pdf_loader import PDFLoader
from document_splitter import DocumentSplitter
from embeddings_generator import EmbeddingsGenerator
from rag_agent import RAGAgent
from utils import create_prompt_template
from langchain_ollama import ChatOllama


def main():
    file_path = input("Please enter the path to the PDF file: ")
    user_prompt = input("Please enter your prompt for the LLM: ")

    pdf_loader = PDFLoader(file_path)
    documents = pdf_loader.load_documents()

    splitter = DocumentSplitter()
    splits = splitter.split_documents(documents)

    embeddings_generator = EmbeddingsGenerator()
    retriever = embeddings_generator.generate_embeddings(splits)

    llm = ChatOllama(model="llama3", temperature=0)
    prompt_template = create_prompt_template()

    rag_agent = RAGAgent(retriever, llm, prompt_template)
    results = rag_agent.get_answer(user_prompt)

    print(results["answer"])

    for document in results["context"]:
        print(document)
        print()


if __name__ == "__main__":
    main()
