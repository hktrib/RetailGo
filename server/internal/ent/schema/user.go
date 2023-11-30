package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.String("email"),
		field.Bool("is_owner"),
		field.Int("store_id").Optional(),
		field.String("clerk_user_id").Optional(),
		field.String("first_name").Optional(),
		field.String("last_name").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("store", Store.Type).Through("UserToStore", UserToStore.Type),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email"),
	}
}
