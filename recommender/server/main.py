from fastapi import FastAPI, status, Response, HTTPException
from data_models import Item, ItemBatch
from vectorize import Vectorizer
from weaviate_writer import WeaviateWriter
import numpy as np

app = FastAPI()

CLIP_vectorizer = Vectorizer()
Writer = WeaviateWriter()

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
            Writer.write_item_vector(item, vector)
            print("Item Write")
            # If you are going to reach a change in store, write the discounted average to the store vector as well
            if current_store_id and item.StoreId != current_store_id:
                print("Hello from Write to store")
                Writer.write_store_vector(item.StoreId, todays_store_vector / todays_sales)
                print("Write Store Vector")
                todays_store_vector *= 0
                print("Reset Store Vector")
                todays_sales = 0
            todays_store_vector += item.NumberSoldSinceUpdate * vector    
            todays_sales += item.NumberSoldSinceUpdate

        # If failure for any particular item, collect the items which failed.
        except:
            failures.add(item.ID)

    try:
        print("Hello from Write to store")
        Writer.write_store_vector(item.StoreId, todays_store_vector / todays_sales)
        print("Write Store Vector")
        todays_store_vector *= 0
        print("Reset Store Vector")
        todays_sales = 0
        print("Write To Store")

    except:
        failures.add(item.ID)

    # Send back a JSON with all the item ids that were not successfully written to Weaviate. (If none, assume all went well).   
    if len(failures) > 0:
        print("Failures:", failures)
        raise HTTPException(status_code = 200) 
    
    return list(failures)