// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ItemsColumns holds the columns for the "items" table.
	ItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "photo", Type: field.TypeBytes},
		{Name: "quantity", Type: field.TypeInt},
		{Name: "store_id", Type: field.TypeInt},
		{Name: "category", Type: field.TypeString},
	}
	// ItemsTable holds the schema information for the "items" table.
	ItemsTable = &schema.Table{
		Name:       "items",
		Columns:    ItemsColumns,
		PrimaryKey: []*schema.Column{ItemsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "item_store_id",
				Unique:  false,
				Columns: []*schema.Column{ItemsColumns[4]},
			},
		},
	}
	// StoresColumns holds the columns for the "stores" table.
	StoresColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "store_name", Type: field.TypeString},
	}
	// StoresTable holds the schema information for the "stores" table.
	StoresTable = &schema.Table{
		Name:       "stores",
		Columns:    StoresColumns,
		PrimaryKey: []*schema.Column{StoresColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "store_id",
				Unique:  false,
				Columns: []*schema.Column{StoresColumns[0]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString},
		{Name: "is_owner", Type: field.TypeBool},
		{Name: "real_name", Type: field.TypeString},
		{Name: "store_id", Type: field.TypeInt},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_username_email",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[1], UsersColumns[2]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ItemsTable,
		StoresTable,
		UsersTable,
	}
)

func init() {
}
