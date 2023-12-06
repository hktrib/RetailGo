from data_models import ItemBatch


class Vectorizer(object):

    def __init__(self):
        # Dummy model
        self.model = lambda item: [0, 1, 2]
        self.dimension = 3
    
    def vectorize(self, items: ItemBatch):
        return list(map(self.model, items))

