package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Item holds the schema definition for the Item entity.
type Item struct {
	ent.Schema
}

func (Item) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.String("name"),
		field.String("photo"),
		field.Int("quantity"),
		field.Float("price").
			SchemaType(map[string]string{
				dialect.Postgres: "decimal(10,2)",
			}),
		field.Int("store_id"),
		field.String("stripe_price_id").Optional(),
		field.String("category_name").Optional(),
		field.String("stripe_product_id").Optional(),
		field.String("weaviate_id").Optional(),
		field.Bool("vectorized").Optional(),
		field.Int("number_sold_since_update").Optional(),
		field.String("date_last_sold").Optional(),
	}
}

func (Item) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("category", Category.Type).Ref("items").Through("category_item", CategoryItem.Type),
		edge.From("store", Store.Type).Ref("items").Required().Field("store_id").Unique(),
	}
}

func (Item) Indexes() []ent.Index {
	return []ent.Index{

		index.Fields("id"),
		index.Fields("stripe_product_id"),
	}
}
