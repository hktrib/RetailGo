// Code generated by ent, DO NOT EDIT.

package item

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the item type in the database.
	Label = "item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPhoto holds the string denoting the photo field in the database.
	FieldPhoto = "photo"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldStoreID holds the string denoting the store_id field in the database.
	FieldStoreID = "store_id"
	// FieldStripePriceID holds the string denoting the stripe_price_id field in the database.
	FieldStripePriceID = "stripe_price_id"
	// FieldStripeProductID holds the string denoting the stripe_product_id field in the database.
	FieldStripeProductID = "stripe_product_id"
	// FieldCategoryName holds the string denoting the category_name field in the database.
	FieldCategoryName = "category_name"
	// FieldNumberSold holds the string denoting the number_sold field in the database.
	FieldNumberSold = "number_sold"
	// FieldDateLastSold holds the string denoting the date_last_sold field in the database.
	FieldDateLastSold = "date_last_sold"
	// EdgeCategory holds the string denoting the category edge name in mutations.
	EdgeCategory = "category"
	// EdgeStore holds the string denoting the store edge name in mutations.
	EdgeStore = "store"
	// EdgeCategoryItem holds the string denoting the category_item edge name in mutations.
	EdgeCategoryItem = "category_item"
	// Table holds the table name of the item in the database.
	Table = "items"
	// CategoryTable is the table that holds the category relation/edge. The primary key declared below.
	CategoryTable = "category_items"
	// CategoryInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	CategoryInverseTable = "categories"
	// StoreTable is the table that holds the store relation/edge.
	StoreTable = "items"
	// StoreInverseTable is the table name for the Store entity.
	// It exists in this package in order to avoid circular dependency with the "store" package.
	StoreInverseTable = "stores"
	// StoreColumn is the table column denoting the store relation/edge.
	StoreColumn = "store_id"
	// CategoryItemTable is the table that holds the category_item relation/edge.
	CategoryItemTable = "category_items"
	// CategoryItemInverseTable is the table name for the CategoryItem entity.
	// It exists in this package in order to avoid circular dependency with the "categoryitem" package.
	CategoryItemInverseTable = "category_items"
	// CategoryItemColumn is the table column denoting the category_item relation/edge.
	CategoryItemColumn = "item_id"
)

// Columns holds all SQL columns for item fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldPhoto,
	FieldQuantity,
	FieldPrice,
	FieldStoreID,
	FieldStripePriceID,
	FieldStripeProductID,
	FieldCategoryName,
	FieldNumberSold,
	FieldDateLastSold,
}

var (
	// CategoryPrimaryKey and CategoryColumn2 are the table columns denoting the
	// primary key for the category relation (M2M).
	CategoryPrimaryKey = []string{"category_id", "item_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Item queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByQuantity orders the results by the quantity field.
func ByQuantity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQuantity, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByStoreID orders the results by the store_id field.
func ByStoreID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStoreID, opts...).ToFunc()
}

// ByStripePriceID orders the results by the stripe_price_id field.
func ByStripePriceID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStripePriceID, opts...).ToFunc()
}

// ByStripeProductID orders the results by the stripe_product_id field.
func ByStripeProductID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStripeProductID, opts...).ToFunc()
}

// ByCategoryName orders the results by the category_name field.
func ByCategoryName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCategoryName, opts...).ToFunc()
}

// ByNumberSold orders the results by the number_sold field.
func ByNumberSold(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNumberSold, opts...).ToFunc()
}

// ByDateLastSold orders the results by the date_last_sold field.
func ByDateLastSold(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDateLastSold, opts...).ToFunc()
}

// ByCategoryCount orders the results by category count.
func ByCategoryCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCategoryStep(), opts...)
	}
}

// ByCategory orders the results by category terms.
func ByCategory(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCategoryStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByStoreField orders the results by store field.
func ByStoreField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStoreStep(), sql.OrderByField(field, opts...))
	}
}

// ByCategoryItemCount orders the results by category_item count.
func ByCategoryItemCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCategoryItemStep(), opts...)
	}
}

// ByCategoryItem orders the results by category_item terms.
func ByCategoryItem(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCategoryItemStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCategoryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CategoryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, CategoryTable, CategoryPrimaryKey...),
	)
}
func newStoreStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StoreInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, StoreTable, StoreColumn),
	)
}
func newCategoryItemStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CategoryItemInverseTable, CategoryItemColumn),
		sqlgraph.Edge(sqlgraph.O2M, true, CategoryItemTable, CategoryItemColumn),
	)
}
