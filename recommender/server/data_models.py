from pydantic import BaseModel, Field
from typing import List, Optional

class Item(BaseModel):
    id: int
    weaviate_id: str
    photo: Optional[str]
    name: str
    store_id: int
    number_sold_since_update: Optional[int] = Field(default = 0)
    quantity: int
    price: float
    stripe_price_id: str
    category_name: str
    stripe_product_id: str
    edges: dict

class ItemBatch(BaseModel):
    Items: List[Item]
