import weaviate
from data_models import Item

import os
import numpy as np

RESULT_LIMIT = 10

class DB(object):
    def __init__(self, discount_factor, dimension = 300):
        try:
            self.client = weaviate.Client(
                    url = "https://retailgo-recengine-eb6uzggu.weaviate.network",
                    auth_client_secret = weaviate.AuthApiKey(api_key="isYZjIQAxvMOhFkTUt0bI5xqETVAHGHqO6fU"),
            )
        except Exception as error:
            print("Failed to connect to Vector Database at {WEAVIATE_HOSTNAME}, with Secret Key {WEAVIATE_SK}".format(os.getenv("WEAVIATE_HOSTNAME"), os.getenv("WEAVIATE_SK")), error)
            assert False

        self.dimension = dimension

        self.discount_factor = discount_factor
        self.current_store = None
        self.current_store_vector = np.zeros((self.dimension, 1)).astype(np.float16)
        self.current_sales = 0
    
    # Item Operations: Add item to Store and Search by vector for similar items
    
    def add_item(self, store_id, vector):
        self.current_store = store_id
        self.current_store_vector += vector
        self.current_sales += 1

    def find_item_near_vector(self, vector, result_limit = 50):

        try:
            nearby_items = self.client.query.get("item", ["name", "categoryName", "imageURL", "price"]).with_near_vector({
                'vector': vector
            }).with_limit(result_limit).with_additional(["distance"]).do()

            return nearby_items

        except Exception as error:
            print("Failed to search for items near vector:", error)

    # Store Ops: Get the id and vector of a store, Create a new store, Update the store's vector.

    def get_store(self, store_id):
        store = self.client.query.get(
            "store",
            []
        ).with_where(
            {"path": ["store_id"], "operator": "Equal", "valueNumber": store_id}
        ).with_additional(["id", "vector"]).with_limit(1).do()["data"]["Get"]["Store"]

        if len(store) == 0:
            print("No such store", store_id, "exists on Weaviate as of now.")
            return None

        return store[0]["_additional"]

    def create_store(self, store_id):
        zero_vector = np.zeros((self.dimension, 1)).astype(np.float16)
        try:
            store_uuid = self.client.data_object.create({
                    "store_id": store_id,
                },
                vector = zero_vector,
                class_name="store"
                )
        except Exception as error:
            print("Failed to create new store {store_id}:".format(store_id), error)

            return {
                "success": False
            }
        

        return {
            "vector": zero_vector,
            "id": store_uuid,
            "success": True
        }

    def update_store(self, store_uuid: str, prev_store_vector, curr_store_vector):
        try:
            self.client.data_object.update(
                    uuid = store_uuid,
                    class_name = "store",
                    data_object={},
                    vector = (self.discount_factor * prev_store_vector + curr_store_vector)
                )

        except Exception as error:
            print(store_uuid)
            print("Failed to update store {store_uuid}'s vector:".format(store_uuid), error)


    # Write Item.
    def write_item(self, item: Item, vector: np.ndarray):
        success = True

        # If item doesn't have weaviate id, return error

        if item.weaviate_id == "":
            success = False
            print("Item {id} has no weaviate id:".format(item.id), error)
            return success

        # Else, try weaviate update. If that fails, return error.
        try:
            self.client.data_object.update(
            uuid = item.weaviate_id,
            data_object={},
            class_name="item",
            vector = vector
            )
        except Exception as error:
            itemId = item.id
            print("Error updating vector for item", itemId, error)

            success = False

        # Update the store's periodic total to account for the new item
        if item.number_sold_since_update != 0:
            self.add_item(item.store_id, vector)

        # If it didn't fail, return True for success
        return success

    # Write Store.
    def write_store(self):
        # Skip the day if no sales were made
        if self.current_sales == 0:
            return
        
        # Calculate average of all items:
        average_item_vector = self.current_store_vector / self.current_sales

        # Find the corresponding store.
        store = self.get_store(self.current_store)
        
        # The store doesn't yet exist, so create it
        if not store:
            store = self.create_store(self.current_store)

            if not store["success"]:
                return
        
        if len(store["vector"]) == 0:
            store["vector"] = np.zeros((self.dimension, 1))

        # Retrieve the previous store vector
        prev_store_vector = np.array(store["vector"])
        store_uuid = store['id']

        # Update the store vector
        self.update_store(store_uuid, prev_store_vector, average_item_vector)

        # Clear the store state
        self.current_store = None
        self.current_sales = 0
        self.current_store_vector *= 0

    # Candidate Retrieval for a given store
    def retrieve_candidates_to_recommend(self, store_id):

        # Find the store's vector

        try:
            store = self.get_store(store_id)
        except Exception as error:
            print("Failed to retrieve store object:", error)
            return None

        # If it doesn't exist, start with no preferences - reranking will improve these cold start recommendations

        if not store or len(store["vector"]) == 0:
            store_vector = np.zeros((self.dimension, 1)).astype(np.float16)
        else:
            store_vector = store["vector"]
        
        # Search for nearby items, candidates to be recommended

        try:
            candidates = self.find_item_near_vector(store_vector, result_limit=RESULT_LIMIT)
        except Exception as error:
            print("Failed to find nearby items:", error)
            return None

        return candidates["data"]["Get"]["Item"]