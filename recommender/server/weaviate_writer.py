import weaviate
import os
import numpy as np

class WeaviateWriter(object):
    def __init__(self):
        self.auth_config = weaviate.AuthApiKey(api_key=os.getenv("WEAVIATE_SK"))
        self.client = weaviate.Client(
                url = "https://retailgo-recengine-eb6uzggu.weaviate.network", #os.getenv("WEAVIATE_HOSTNAME"),
                auth_client_secret=self.auth_config,
        )
        self.discount_factor = 0.99

    def write_item_vector(self, item, vector):
        self.client.data_object.update(
            uuid = item.WeaviateID,
            class_name="item",
            vector = vector
        )
    
    def write_store_vector(self, store_id, average_today):

        store_exists = False

        try:
            store = self.client.query.get(
                "store",
                []
            ).with_additional(["id", "vector"])

            curr_vector = store["vector"]
            store_uuid = store["id"]

            store_exists = True

        except:
            print("Store", store_id, "didn't exist")
            pass

        if store_exists:
            self.client.data_object.update(
                uuid = store_uuid,
                class_name = "store",
                vector = self.discount_factor * curr_vector + average_today
            )

        else:
            self.client.data_object.create({
                "store_id": store_id,
            },
            vector = average_today,
            class_name="store"
            )