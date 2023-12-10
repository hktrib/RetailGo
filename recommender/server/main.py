from fastapi import FastAPI, status, Response, HTTPException
from data_models import Item, ItemBatch
from vdb import DB
from recsys import Recommender
from models import DefaultWord2VecModel, DefaultGloveModel
import uvicorn
import os

app = FastAPI()

# A Recommender Class that does embedding, ranking and reranking.
try:
    default_w2v = DefaultGloveModel()
    recommender = Recommender(default_w2v)
except Exception as error:
    print("Initialization failed:", error)
    assert False

print("Recommender initialized")

# DB() to init the vector database
database = DB(discount_factor=0.99, dimension=recommender.dimension)
print("DB initialized")


@app.post("/vectorizeItems", status_code = 201)
def vectorize_batch(batch: ItemBatch, response: Response):

    if len(batch.Items) == 0:
        return []

    embedding_failed = []
    batch_size = len(batch.Items)

    # Embeds full batch of items.
    try:
        vectors = recommender.embed_batch(batch.Items)
    except Exception as error:
        print("Failed to embed items:", error)
        raise HTTPException(500)

    # Then, iterate through all items, and write them, one by one to Weaviate (need an interface to do this, something like db.write_item(item))

    written_items = 0

    for item, vector in zip(batch.Items, vectors):
        vector = vector.reshape(-1, 1)
        success = database.write_item(item, vector)

        if not success:
            embedding_failed.append(item.id)
        
        written_items += 1

        if written_items == batch_size or batch.Items[written_items].store_id != item.store_id:
            database.write_store()

    print("Out of", batch_size, "items,", len(embedding_failed), "failed to be embedded.")

    # Return the list of failed items' ids.
    return embedding_failed

@app.get("/recommend/{store_id}", status_code = 200)
def recommend(store_id: int):
    # Retrieve the store vector; Search by the store vector for items - db.find_similar_items, raise a 404 if the store is not found on Weaviate.
    results = database.retrieve_candidates_to_recommend(store_id)
    
    if not results:
        print("Failed to retrieve candidates.")
        raise HTTPException(500)    

    # Filter them
    try:
        results = recommender.filter(results, store_id)
    except Exception as error:
        print("Failed to Filter:", error)
        raise HTTPException(500)

    # Rerank them
    try:
        results = recommender.rerank(results)
    except Exception as error:
        print("Failed to rerank:", error)
        raise HTTPException(500)

    # Return all
    return results