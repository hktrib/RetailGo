package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// intermediate table for the relationship between Category and Item.
type ClerkUser_Store struct {
	ent.Schema
}

func (ClerkUser_Store) Annotations() []schema.Annotation {
	return []schema.Annotation{}
}

func (ClerkUser_Store) Fields() []ent.Field {
	return []ent.Field{
		field.String("clerk_id"),
		field.Int("store_id"),
	}
}
