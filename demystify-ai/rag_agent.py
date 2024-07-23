from langchain.chains import create_retrieval_chain
from langchain.chains.combine_documents import create_stuff_documents_chain


class RAGAgent:
    def __init__(self, retriever, llm, prompt_template):
        self.retriever = retriever
        self.llm = llm
        self.prompt_template = prompt_template

    def create_rag_chain(self):
        question_answer_chain = create_stuff_documents_chain(
            self.llm, self.prompt_template
        )
        return create_retrieval_chain(self.retriever, question_answer_chain)

    def get_answer(self, prompt):
        rag_chain = self.create_rag_chain()
        return rag_chain.invoke({"input": prompt})
