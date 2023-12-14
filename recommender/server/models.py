from typing import List
from data_models import Item
import gensim.downloader as downloader
from gensim.models import FastText, KeyedVectors
import numpy as np

class Model(object):
    def preprocess(self, items: List[Item]):
        """ Should return an iterable of preprocessed items, extracting and processing only attributes relevant to the implemented forward."""
        raise NotImplementedError

    def forward(self, items_features):
        """Should return a numpy array of embedding for each item"""
        raise NotImplementedError

class DefaultWord2VecModel(Model):
    def __init__(self, corpus = "word2vec-google-news-300", model_file = None):
        super().__init__()
        if not model_file:
            self.w2v = downloader.load(corpus)
            self.dimension = int(corpus.split("-")[-1])

        else:
            self.w2v = KeyedVectors.load(model_file, mmap = 'r').wv
            self.dimension = self.w2v.vector_size

    def preprocess(self, items: List[Item]):
        return map(
            lambda item: [item.category_name] + item.name.split(),
            items
        )
    
    def forward_single(self, item_features):
        word_vectors = [self.w2v[word].reshape(-1, 1) for word in item_features if word in self.w2v]

        if len(word_vectors) == 0:
            word_vectors = [np.zeros((self.dimension, 1)).astype(np.float16)]

        return np.mean(
            np.concatenate(
                word_vectors, axis = 1
            ), axis = 1
        )

    def forward(self, items_features):
        return np.array([self.forward_single(item_features) for item_features in items_features]).astype(np.float16)
    

class DefaultGloveModel(Model):
    def __init__(self, corpus = "glove-twitter-25"):
        super().__init__()
        self.glove = downloader.load(corpus)
        self.dimension = int(corpus.split("-")[-1])
    
    def preprocess(self, items: List[Item]):
        return map(
            lambda item: [item.category_name] + item.name.split(),
            items
        )
    
    def forward_single(self, item_features):
        word_vectors = [self.glove[word].reshape(-1, 1) for word in item_features if word in self.glove]

        if len(word_vectors) == 0:
            word_vectors = [np.zeros((self.dimension, 1)).astype(np.float16)]

        return np.mean(
            np.concatenate(
                word_vectors, axis = 1
            ), axis = 1
        )

    def forward(self, items_features):
        return np.array([self.forward_single(item_features) for item_features in items_features]).astype(np.float16)
    
class ProductFastTextModel(Model):
    def __init__(self, corpus = None, model_file = "./models/Product FastText/fast_text_product_descriptions_10000-25-5-5.model"):
        super().__init__()
        if corpus:
            self.w2v = downloader.load(corpus)
            self.dimension = int(corpus.split("-")[-1])

        else:
            self.w2v = FastText.load(model_file).wv
            self.dimension = self.w2v.vector_size

    def preprocess(self, items: List[Item]):
        return map(
            lambda item: [item.category_name] + item.name.split(),
            items
        )
    
    def forward_single(self, item_features):
        word_vectors = [self.w2v[word].reshape(-1, 1) for word in item_features if word in self.w2v]

        if len(word_vectors) == 0:
            word_vectors = [np.zeros((self.dimension, 1)).astype(np.float16)]

        return np.mean(
            np.concatenate(
                word_vectors, axis = 1
            ), axis = 1
        )

    def forward(self, items_features):
        return np.array([self.forward_single(item_features) for item_features in items_features]).astype(np.float16)
