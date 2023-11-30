package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// intermediate table for the relationship between Category and Item.
type ClerkUser_Store struct {
	ent.Schema
}

func (ClerkUser_Store) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("clerk_user_id", "store_id"),
	}
}

func (ClerkUser_Store) Fields() []ent.Field {
	return []ent.Field{
		field.String("clerk_user_id"),
		field.String("store_id"),
	}
}

func (ClerkUser_Store) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("store", Store.Type).Unique().Required().Field("store_id"),
	}
}
