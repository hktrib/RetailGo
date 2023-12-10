from typing import List
from data_models import Item
import gensim.downloader as downloader
import numpy as np

class Model(object):
    def preprocess(self, items: List[Item]):
        """ Should return an iterable of preprocessed items, extracting and processing only attributes relevant to the implemented forward."""
        raise NotImplementedError

    def forward(self, items_features):
        """Should return a numpy array of embedding for each item"""
        raise NotImplementedError

class DefaultWord2VecModel(Model):
    def __init__(self, corpus = "word2vec-google-news-300"):
        super().__init__()
        self.w2v = downloader.load(corpus)
        self.dimension = int(corpus.split("-")[-1])
    
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