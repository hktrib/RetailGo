package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Category struct {
	ent.Schema
}

func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique(),
		field.String("name"),
		field.Bytes("photo"),
		field.Int("store_id").Unique(),
	}
}

func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		// establish a many to many relationship with the Item entity.
		edge.To("items", Item.Type).Through("category_item", CategoryItem.Type),
		edge.From("store", Store.Type).Ref("categories").Field("store_id").Unique().Required(),
	}
}

func (Category) Indexes() []ent.Index {
	return []ent.Index{
		// index from name
		index.Fields("name").Unique(),
	}
}
