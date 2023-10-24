package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Item holds the schema definition for the Item entity.
type Item struct {
	ent.Schema
}

// Fields of the Item.
func (Item) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.String("name"),
		field.Bytes("photo"),
		field.Int("quantity"),
		field.Int("store_id"),
		field.String("category"),
	}
}

// Edges of the Item.
func (Item) Edges() []ent.Edge {
	return nil
}

func (Item) Indexes() []ent.Index {
	return []ent.Index{
		// index.Fields("item_name").
		// 	Edges("store_name").
		// 	Unique(),
		index.Fields("store_id").
			Unique(),
	}
}
