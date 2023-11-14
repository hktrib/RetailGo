package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Store holds the schema definition for the Store entity.
type Store struct {
	ent.Schema
}

// Fields of the Store.
func (Store) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.String("store_name"),
		field.String("owner_Email").Optional(),
	}
}

// Edges of the Store.
func (Store) Edges() []ent.Edge {
	return []ent.Edge{
		// establish a many to many relationship with the Item entity.
		edge.To("items", Item.Type),
		edge.To("categories", Category.Type),
		edge.From("user", User.Type).Ref("store").Through("UserToStore", UserToStore.Type),
	}
}

func (Store) Indexes() []ent.Index {
	return []ent.Index{
		// index.Fields("item_name").
		// 	Edges("store_name").
		// 	Unique(),
		index.Fields("id"),
	}
}
