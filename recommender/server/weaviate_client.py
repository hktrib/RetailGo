import weaviate
import os
import numpy as np

class Weaviate(object):
    def __init__(self):
        self.auth_config = weaviate.AuthApiKey(api_key=os.getenv("WEAVIATE_SK"))
        self.client = weaviate.Client(
                url = "https://retailgo-recengine-eb6uzggu.weaviate.network",
                auth_client_secret=self.auth_config,
        )
        self.discount_factor = 1

    def write_item_vector(self, item, vector):
        print("Weaviate_Id:", item.weaviate_id)
        self.client.data_object.update(
        uuid = item.weaviate_id,
        data_object={
            "Updated": True
        },
        class_name="item",
        vector = vector
        )
        print("Worked?")

    def write_store_vector(self, store_id, average_today):

        store_exists = False

        print("Average Today:", average_today)

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
            print("Updating existing store", curr_vector.shape, average_today.shape)
            print("Store_UUID:", store_uuid)
            self.client.data_object.update(
                uuid = store_uuid,
                class_name = "store",
                data_object={
                    'completed': True
                },
                vector = self.discount_factor * curr_vector + average_today
            )

        else:
            print("Creating store")
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
        return self.client.query().get("item", ["name", "categoryName", "imageURL", "price"]).with_near_vector({
            'vector': store_vector
        }).with_limit(50).do()
