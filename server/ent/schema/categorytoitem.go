package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// intermediate table for the relationship between Category and Item.
type CategoryItem struct {
	ent.Schema
}

func (CategoryItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("category_id", "item_id"),
	}
}

func (CategoryItem) Fields() []ent.Field {
	return []ent.Field{
		field.Int("category_id"),
		field.Int("item_id"),
	}
}

func (CategoryItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("category", Category.Type).Unique().Required().Field("category_id"),
		edge.To("item", Item.Type).Unique().Required().Field("item_id"),
	}
}
