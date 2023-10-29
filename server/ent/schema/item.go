package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
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
		field.Float("price").
			SchemaType(map[string]string{
				dialect.Postgres: "decimal(10,2)", // Override MySQL.
			}),
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
		index.Fields("store_id", "category").Annotations(entsql.IndexWhere("store_id IS NOT NULL AND category IS NOT NULL")),
		index.Fields("store_id"),
		index.Fields("category"),
	}
}
