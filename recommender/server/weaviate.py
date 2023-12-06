import weaviate
import os
import numpy as np

class Weaviate(object):
    def __init__(self):
        self.auth_config = weaviate.auth.AuthApiKey(api_key=os.getenv("WEAVIATE_SK"))
        self.client = weaviate.Client(
                url = os.getenv("WEAVIATE_HOSTNAME"),
                auth_client_secret=self.auth_config,
        )
        self.discount_factor = 0.99

    def write_item_vector(self, item, vector):
        self.client.data_object.update(
        uuid = item.WeaviateID,
        data_object={},
        class_name="item",
        vector = vector
        )

    def write_store_vector(self, store_id, average_today):

        store_exists = False

        try:
            store = self.client.query.get(
                "store",
                []
            ).with_where(
                {"path": ["store_id"], "operator": "Equal", "valueNumber": store_id}
            ).with_additional(["id", "vector"]).with_limit(1).do()["data"]["Get"]["Store"][0]["_additional"]

            curr_vector = np.array(store["vector"])
            store_uuid = store["id"]

            store_exists = True
        except:
            pass

        if store_exists:
            self.client.data_object.update(
                uuid = store_uuid,
                class_name = "store",
                data_object={},
                vector = self.discount_factor * curr_vector + average_today
            )

        else:
            _ = self.client.data_object.create({
                "store_id": store_id,
            },
            vector = average_today,
            class_name="store"
            )
    
    def get_store_vector(self, store_id):
        store = self.client.query.get(
                    "store",
                    []
                ).with_where(
                    {"path": ["store_id"], "operator": "Equal", "valueNumber": store_id}
                ).with_additional(["id", "vector"]).with_limit(1).do()["data"]["Get"]["Store"][0]["_additional"]

        store_vector = np.array(store["vector"])
        
        return store_vector

    def search(self, store_vector):
        return self.client.query().get("item", ["name", "categoryName", "imageURL", "price"]).with_limit(50).do()
