import numpy as np
from data_models import Item
from typing import List
from models import Model

class Recommender(object):
    def __init__(self, model: Model):
        # Load the model
        self.model = model

    def embed_batch(self, items: List[Item], batch_size = 64):
        batch_item_vectors = []

        # Filter only relevant attributes (model.preprocess), run the model.forward over them in batches
        for minibatch_start in range(0, len(items), batch_size):
            preprocessed_items = self.model.preprocess(items[minibatch_start : minibatch_start + batch_size])
            minibatch_item_vectors = self.model.forward(preprocessed_items)

            batch_item_vectors.append(minibatch_item_vectors)

        return np.concatenate(batch_item_vectors, axis = 0)
    
    def filter(self, candidates, store_id):
        print("Candidates:", candidates)
        return list(
            filter(
                lambda candidate: candidate["storeId"] == store_id if "storeId" in candidate else True,
                candidates
            )
        )

    def rerank(self, candidates, final_candidates = 15):
        # Maximize diversity among candidates
        return candidates[:final_candidates]