package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// intermediate table for the relationship between Category and Item.
type UserToStore struct {
	ent.Schema
}

func (UserToStore) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("user_id", "store_id"),
	}
}

func (UserToStore) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("store_id"),
		field.String("clerk_user_id").Optional(),
		field.Int("permission_level").Optional(),
		field.Int("joined_at").Optional(),
	}
}

func (UserToStore) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique().Required().Field("user_id"),
		edge.To("store", Store.Type).Unique().Required().Field("store_id"),
	}
}
