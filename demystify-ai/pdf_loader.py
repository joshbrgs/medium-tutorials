from langchain_community.document_loaders import PyPDFLoader


class PDFLoader:
    def __init__(self, file_path):
        self.file_path = file_path

    def load_documents(self):
        loader = PyPDFLoader(self.file_path)
        return loader.load()
