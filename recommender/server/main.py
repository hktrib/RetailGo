from fastapi import FastAPI, status, Response, HTTPException
from data_models import Item, ItemBatch
from vectorize import Vectorizer
from weaviate_client import Weaviate
import numpy as np

app = FastAPI()

vectorizer = Vectorizer()
WeaviateClient = Weaviate()

@app.post("/vectorizeItems", status_code = 201)
def vectorize_batch(batch: ItemBatch, response: Response):
    # Call a vectorizing function on the whole batch of items

    print("Batch:", batch)
    if len(batch.Items) == 0:
        return {"idsFailedToVectorize": []}
    
    try:       
        vectors = vectorizer.vectorize(batch)
    except:
        print("Vectorizer Error")
        raise HTTPException(status_code = 500)

    # Iterate through the batch and write the vectors to weaviate

    current_store_id = None
    todays_store_vector = np.zeros(vectorizer.dimension)
    todays_sales = 0

    failures = set()

    for item, vector in zip(batch.Items, vectors):
        try:
            WeaviateClient.write_item_vector(item, vector)
            # If you are going to reach a change in store, write the discounted average to the store vector as well
            if current_store_id and item.store_id != current_store_id:
                WeaviateClient.write_store_vector(item.store_id, todays_store_vector / todays_sales if todays_sales > 0 else todays_store_vector)
                todays_store_vector *= 0
                todays_sales = 0
                current_store_id = item.store_id

            todays_store_vector += item.number_sold_since_update * vector    
            todays_sales += item.number_sold_since_update

        # If failure for any particular item, collect the items which failed.
        except Exception as error:
            print("Error:", error)
            failures.add(item.id)

    try:
        WeaviateClient.write_store_vector(item.store_id, todays_store_vector / todays_sales if todays_sales > 0 else todays_store_vector)
        todays_store_vector *= 0
        todays_sales = 0
        
    except Exception as error:
        print("Error:", error)
        print("Failure:", item.id)
        failures.add(item.id)

    # Send back a JSON with all the item ids that were not successfully written to Weaviate. (If none, assume all went well).   
        
    return {"idsFailedToVectorize": list(failures)}

@app.get("/recommend/{store_id}", status_code = 200)
def recommend(store_id: int):
    # Retrieve the store vector
    store_vector = WeaviateClient.get_store_vector(store_id)
    # Search by the store vector for items
    return WeaviateClient.search(store_vector)