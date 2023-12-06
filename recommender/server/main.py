from fastapi import FastAPI, status, Response, HTTPException
from data_models import Item, ItemBatch
from vectorize import Vectorizer
from weaviate_client import Weaviate
import numpy as np

app = FastAPI()

CLIP_vectorizer = Vectorizer()
WeaviateClient = Weaviate()

@app.post("/vectorizeItems", status_code = 201)
def vectorize_batch(batch: ItemBatch, response: Response):
    # Call a vectorizing function on the whole batch of items

    try:       
        vectors = CLIP_vectorizer.vectorize(batch.items)
    except:
        raise HTTPException(status_code = 500)

    # Iterate through the batch and write the vectors to weaviate

    current_store_id = None
    todays_store_vector = np.zeros(CLIP_vectorizer.dimension)
    todays_sales = 0

    failures = set()

    for item, vector in zip(batch.items, vectors):
        try:
            WeaviateClient.write_item_vector(item, vector)
            # If you are going to reach a change in store, write the discounted average to the store vector as well
            if current_store_id and item.StoreId != current_store_id:
                WeaviateClient.write_store_vector(item.StoreId, todays_store_vector / todays_sales if todays_sales > 0 else todays_store_vector)
                todays_store_vector *= 0
                todays_sales = 0
            todays_store_vector += item.NumberSoldSinceUpdate * vector    
            todays_sales += item.NumberSoldSinceUpdate

        # If failure for any particular item, collect the items which failed.
        except:
            failures.add(item.ID)

    try:
        WeaviateClient.write_store_vector(item.StoreId, todays_store_vector / todays_sales)
        todays_store_vector *= 0
        todays_sales = 0
        
    except:
        failures.add(item.ID)

    # Send back a JSON with all the item ids that were not successfully written to Weaviate. (If none, assume all went well).   
        
    return {"idsFailedToVectorize": list(failures)}

@app.get("/recommend/{store_id}", status_code = 200)

def recommend(store_id: int):
    # Retrieve the store vector
    store_vector = WeaviateClient.get_store_vector(store_id)
    # Search by the store vector for items
    return WeaviateClient.search(store_vector)