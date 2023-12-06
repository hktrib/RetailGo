from data_models import ItemBatch
from transformers import CLIPModel, AutoTokenizer, AutoProcessor
from PIL import Image
import requests
import numpy as np

class Vectorizer(object):

    def __init__(self):

        self.model = CLIPModel.from_pretrained("openai/clip-vit-base-patch32")
        self.tokenizer = AutoTokenizer.from_pretrained("openai/clip-vit-base-patch32")
        self.processor = AutoProcessor.from_pretrained("openai/clip-vit-base-patch32")
        self.dimension = 512

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

    def vectorize(self, items: ItemBatch):
        text_vector = self.model.get_text_features(**self.tokenizer([item.Name for item in items], padding = True, return_tensors = "pt")).detach().numpy().astype(np.float16)
        img_vector = self.get_images_features(items)

        return text_vector