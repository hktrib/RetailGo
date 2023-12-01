// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CategoriesColumns holds the columns for the "categories" table.
	CategoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "photo", Type: field.TypeBytes},
		{Name: "store_id", Type: field.TypeInt},
	}
	// CategoriesTable holds the schema information for the "categories" table.
	CategoriesTable = &schema.Table{
		Name:       "categories",
		Columns:    CategoriesColumns,
		PrimaryKey: []*schema.Column{CategoriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "categories_stores_categories",
				Columns:    []*schema.Column{CategoriesColumns[3]},
				RefColumns: []*schema.Column{StoresColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "category_name",
				Unique:  true,
				Columns: []*schema.Column{CategoriesColumns[1]},
			},
		},
	}
	// CategoryItemsColumns holds the columns for the "category_items" table.
	CategoryItemsColumns = []*schema.Column{
		{Name: "category_id", Type: field.TypeInt},
		{Name: "item_id", Type: field.TypeInt},
	}
	// CategoryItemsTable holds the schema information for the "category_items" table.
	CategoryItemsTable = &schema.Table{
		Name:       "category_items",
		Columns:    CategoryItemsColumns,
		PrimaryKey: []*schema.Column{CategoryItemsColumns[0], CategoryItemsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "category_items_categories_category",
				Columns:    []*schema.Column{CategoryItemsColumns[0]},
				RefColumns: []*schema.Column{CategoriesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "category_items_items_item",
				Columns:    []*schema.Column{CategoryItemsColumns[1]},
				RefColumns: []*schema.Column{ItemsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ItemsColumns holds the columns for the "items" table.
	ItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "photo", Type: field.TypeString},
		{Name: "quantity", Type: field.TypeInt},
		{Name: "price", Type: field.TypeFloat64, SchemaType: map[string]string{"postgres": "decimal(10,2)"}},
		{Name: "stripe_price_id", Type: field.TypeString},
		{Name: "stripe_product_id", Type: field.TypeString},
		{Name: "category_name", Type: field.TypeString},
		{Name: "weaviate_id", Type: field.TypeString},
		{Name: "vectorized", Type: field.TypeBool},
		{Name: "store_id", Type: field.TypeInt},
	}
	// ItemsTable holds the schema information for the "items" table.
	ItemsTable = &schema.Table{
		Name:       "items",
		Columns:    ItemsColumns,
		PrimaryKey: []*schema.Column{ItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "items_stores_items",
				Columns:    []*schema.Column{ItemsColumns[10]},
				RefColumns: []*schema.Column{StoresColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "item_id",
				Unique:  false,
				Columns: []*schema.Column{ItemsColumns[0]},
			},
			{
				Name:    "item_stripe_product_id",
				Unique:  false,
				Columns: []*schema.Column{ItemsColumns[6]},
			},
		},
	}
	// StoresColumns holds the columns for the "stores" table.
	StoresColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeString, Unique: true},
		{Name: "store_name", Type: field.TypeString},
		{Name: "created_by", Type: field.TypeString},
		{Name: "owner_email", Type: field.TypeString, Nullable: true},
		{Name: "store_address", Type: field.TypeString, Nullable: true},
		{Name: "store_phone", Type: field.TypeString, Nullable: true},
		{Name: "store_type", Type: field.TypeString, Nullable: true},
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
		{Name: "email", Type: field.TypeString},
		{Name: "is_owner", Type: field.TypeBool},
		{Name: "store_id", Type: field.TypeInt, Nullable: true},
		{Name: "clerk_user_id", Type: field.TypeString, Nullable: true},
		{Name: "first_name", Type: field.TypeString, Nullable: true},
		{Name: "last_name", Type: field.TypeString, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_email",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[1]},
			},
		},
	}
	// UserToStoresColumns holds the columns for the "user_to_stores" table.
	UserToStoresColumns = []*schema.Column{
		{Name: "permission_level", Type: field.TypeInt, Nullable: true},
		{Name: "joined_at", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "store_id", Type: field.TypeInt},
	}
	// UserToStoresTable holds the schema information for the "user_to_stores" table.
	UserToStoresTable = &schema.Table{
		Name:       "user_to_stores",
		Columns:    UserToStoresColumns,
		PrimaryKey: []*schema.Column{UserToStoresColumns[2], UserToStoresColumns[3]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_to_stores_users_user",
				Columns:    []*schema.Column{UserToStoresColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "user_to_stores_stores_store",
				Columns:    []*schema.Column{UserToStoresColumns[3]},
				RefColumns: []*schema.Column{StoresColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CategoriesTable,
		CategoryItemsTable,
		ItemsTable,
		StoresTable,
		UsersTable,
		UserToStoresTable,
	}
)

func init() {
	CategoriesTable.ForeignKeys[0].RefTable = StoresTable
	CategoryItemsTable.ForeignKeys[0].RefTable = CategoriesTable
	CategoryItemsTable.ForeignKeys[1].RefTable = ItemsTable
	ItemsTable.ForeignKeys[0].RefTable = StoresTable
	UserToStoresTable.ForeignKeys[0].RefTable = UsersTable
	UserToStoresTable.ForeignKeys[1].RefTable = StoresTable
}
