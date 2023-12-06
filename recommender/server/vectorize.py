from data_models import ItemBatch, Item
# from transformers import CLIPModel, AutoTokenizer, AutoProcessor
import requests
import numpy as np
from gensim.models import Word2Vec

class Vectorizer(object):

    def __init__(self):

        # self.model = CLIPModel.from_pretrained("openai/clip-vit-base-patch32")
        self.model = Word2Vec()
        # self.tokenizer = AutoTokenizer.from_pretrained("openai/clip-vit-base-patch32")
        # self.processor = AutoProcessor.from_pretrained("openai/clip-vit-base-patch32")
        self.dimension = 100 #512

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

        return list(
            map(
                lambda item: self.model.wv[item.Name],
                batch.items
            )
        )

    def vectorize(self, items: ItemBatch):
        # text_vector = self.model.get_text_features(**self.tokenizer([item.Name for item in items], padding = True, return_tensors = "pt")).detach().numpy().astype(np.float16)
        # img_vector = self.get_images_features(items)

        # return text_vector

        return self.get_text_features(items)