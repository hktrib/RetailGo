from fastapi import FastAPI, status, Response, HTTPException
from data_models import Item, ItemBatch
from vdb import DB
from recsys import Recommender
from models import DefaultWord2VecModel


app = FastAPI()

# DB() to init the vector database
database = DB(discount_factor=0.99)
print("DB initialized")
# A Recommender Class that does embedding, ranking and reranking.
recommender = Recommender(DefaultWord2VecModel())
print("Recommender initialized")

@app.post("/vectorizeItems", status_code = 201)
def vectorize_batch(batch: ItemBatch, response: Response):

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

    # Return the list of failed items' ids.
    return embedding_failed

@app.get("/recommend/{store_id}", status_code = 200)
def recommend(store_id: int):
    # Retrieve the store vector; Search by the store vector for items - db.find_similar_items, raise a 404 if the store is not found on Weaviate.
    results = database.retrieve_candidates_to_recommend(store_id)
    
    if not results:
        raise HTTPException(500)

    # Rerank them
    try:
        results = recommender.rerank(results)
    except:
        raise HTTPException(500)

    # Return all
    return results