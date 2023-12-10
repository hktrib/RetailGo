from pydantic import BaseModel
from typing import List

class Item(BaseModel):
    id: int
    weaviate_id: str
    photo: str
    name: str
    store_id: int
    number_sold_since_update: int
    quantity: int
    price: float
    stripe_price_id: str
    category_name: str
    stripe_product_id: str
    edges: dict

class ItemBatch(BaseModel):
    Items: List[Item]
