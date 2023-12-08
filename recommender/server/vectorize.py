from data_models import ItemBatch, Item
# from transformers import CLIPModel, AutoTokenizer, AutoProcessor
import requests
import numpy as np
from gensim.models import Word2Vec
import gensim.downloader as downloader

class Vectorizer(object):

    def __init__(self):

        # self.model = CLIPModel.from_pretrained("openai/clip-vit-base-patch32")
        self.model = downloader.load('word2vec-google-news-300') # Word2Vec()
        # self.tokenizer = AutoTokenizer.from_pretrained("openai/clip-vit-base-patch32")
        # self.processor = AutoProcessor.from_pretrained("openai/clip-vit-base-patch32")
        self.dimension = 300 #512

    def get_image_features(self, item):
        if False:
            photo = requests.get(item.Photo, stream = True).raw
            image = Image.open(photo)
            processed_image = self.processor(image, return_tensor = "pt").detach().numpy().astype(np.float16)
            return self.model.get_image_features(**processed_image)
        
        else:
            return np.zeros(self.dimension, dtype = np.float16)

    def get_images_features(self, items):

        return np.array(list(
            map(
                self.get_image_features,
                items
        )))

    def get_text_features(self, batch: ItemBatch):

        print("get_text_features")

        text_vecs = []

        for item in batch.Items:
            print("Item Name", item.name.lower(), " - ", [item_word.lower() for item_word in item.name.split()])
            # try:
            #     # print("Some words recognized:", self.model[])
            # except Exception as error:
            #     print(error)
            #     assert False
            try:
                word_vecs = [self.model[item_word.lower()].reshape(-1, 1) for item_word in item.name.split() if item_word.lower() in self.model]
            except Exception as error:
                print(error)
                assert False
            if len(word_vecs) == 0:
                word_vecs = [np.zeros((self.dimension, 1))]
                print("No Word Vec")

            try:
                word_vecs = np.concatenate(word_vecs, axis = 1)            
                text_vecs.append(np.mean(word_vecs, axis = 1))
            except Exception as error:
                print(error)
                assert False
            
        return np.array(text_vecs)
    
    def vectorize(self, items: ItemBatch):
        # text_vector = self.model.get_text_features(**self.tokenizer([item.Name for item in items], padding = True, return_tensors = "pt")).detach().numpy().astype(np.float16)
        # img_vector = self.get_images_features(items)

        # return text_vector
        print("Text Vector", self.get_text_features(items).shape)
        return self.get_text_features(items)