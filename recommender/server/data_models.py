from pydantic import BaseModel
from typing import List

class Item(BaseModel):
    ID: int
    WeaviateID: str
    Photo: str
    Name: str
    StoreId: int

class ItemBatch(BaseModel):
    items: List[Item]
