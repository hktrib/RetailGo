from data_models import ItemBatch
from transformers import CLIPModel, AutoTokenizer, AutoProcessor
from PIL import Image
import requests
import numpy as np

class Vectorizer(object):

    def __init__(self):

        self.model = CLIPModel.from_pretrained("openai/clip-vit-base-patch32")
        self.tokenizer = AutoTokenizer.from_pretrained("openai/clip-vit-base-patch32")
        self.processor = AutoProcessor.from_pretrained("openai/clipi-vit-base-patch32")
        self.dimension = 512

    def get_image_features(self, items):
        return np.array(map(
            lambda item: self.model.get_image_features(**self.processor(requests.get(item.Photo, stream = True).raw), return_tensor = "pt", dtype = np.float16) if len(item.Photo) > 0 
            else np.zeros(self.dimension),
            items
        ))

    def vectorize(self, items: ItemBatch):

        text_vector = np.array(self.model.get_text_features(**self.tokenizer([item.Name for item in items], padding = True, return_tensors = "pt")), dtype = np.float16)
        img_vector = self.get_image_features(items)

        return (text_vector + img_vector)/2