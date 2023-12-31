// Code generated by ent, DO NOT EDIT.

package store

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the store type in the database.
	Label = "store"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldStoreName holds the string denoting the store_name field in the database.
	FieldStoreName = "store_name"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldOwnerEmail holds the string denoting the owner_email field in the database.
	FieldOwnerEmail = "owner_email"
	// FieldStoreAddress holds the string denoting the store_address field in the database.
	FieldStoreAddress = "store_address"
	// FieldStorePhone holds the string denoting the store_phone field in the database.
	FieldStorePhone = "store_phone"
	// FieldStripeAccountID holds the string denoting the stripe_account_id field in the database.
	FieldStripeAccountID = "stripe_account_id"
	// FieldStoreType holds the string denoting the store_type field in the database.
	FieldStoreType = "store_type"
	// FieldIsAuthorized holds the string denoting the is_authorized field in the database.
	FieldIsAuthorized = "is_authorized"
	// EdgeItems holds the string denoting the items edge name in mutations.
	EdgeItems = "items"
	// EdgeCategories holds the string denoting the categories edge name in mutations.
	EdgeCategories = "categories"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeUserToStore holds the string denoting the usertostore edge name in mutations.
	EdgeUserToStore = "UserToStore"
	// Table holds the table name of the store in the database.
	Table = "stores"
	// ItemsTable is the table that holds the items relation/edge.
	ItemsTable = "items"
	// ItemsInverseTable is the table name for the Item entity.
	// It exists in this package in order to avoid circular dependency with the "item" package.
	ItemsInverseTable = "items"
	// ItemsColumn is the table column denoting the items relation/edge.
	ItemsColumn = "store_id"
	// CategoriesTable is the table that holds the categories relation/edge.
	CategoriesTable = "categories"
	// CategoriesInverseTable is the table name for the Category entity.
	// It exists in this package in order to avoid circular dependency with the "category" package.
	CategoriesInverseTable = "categories"
	// CategoriesColumn is the table column denoting the categories relation/edge.
	CategoriesColumn = "store_id"
	// UserTable is the table that holds the user relation/edge. The primary key declared below.
	UserTable = "user_to_stores"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserToStoreTable is the table that holds the UserToStore relation/edge.
	UserToStoreTable = "user_to_stores"
	// UserToStoreInverseTable is the table name for the UserToStore entity.
	// It exists in this package in order to avoid circular dependency with the "usertostore" package.
	UserToStoreInverseTable = "user_to_stores"
	// UserToStoreColumn is the table column denoting the UserToStore relation/edge.
	UserToStoreColumn = "store_id"
)

// Columns holds all SQL columns for store fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldStoreName,
	FieldCreatedBy,
	FieldOwnerEmail,
	FieldStoreAddress,
	FieldStorePhone,
	FieldStripeAccountID,
	FieldStoreType,
	FieldIsAuthorized,
}

var (
	// UserPrimaryKey and UserColumn2 are the table columns denoting the
	// primary key for the user relation (M2M).
	UserPrimaryKey = []string{"user_id", "store_id"}
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

var (
	// DefaultIsAuthorized holds the default value on creation for the "is_authorized" field.
	DefaultIsAuthorized bool
)

// OrderOption defines the ordering options for the Store queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUUID orders the results by the uuid field.
func ByUUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUUID, opts...).ToFunc()
}

// ByStoreName orders the results by the store_name field.
func ByStoreName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStoreName, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByOwnerEmail orders the results by the owner_email field.
func ByOwnerEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOwnerEmail, opts...).ToFunc()
}

// ByStoreAddress orders the results by the store_address field.
func ByStoreAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStoreAddress, opts...).ToFunc()
}

// ByStorePhone orders the results by the store_phone field.
func ByStorePhone(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStorePhone, opts...).ToFunc()
}

// ByStripeAccountID orders the results by the stripe_account_id field.
func ByStripeAccountID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStripeAccountID, opts...).ToFunc()
}

// ByStoreType orders the results by the store_type field.
func ByStoreType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStoreType, opts...).ToFunc()
}

// ByIsAuthorized orders the results by the is_authorized field.
func ByIsAuthorized(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsAuthorized, opts...).ToFunc()
}

// ByItemsCount orders the results by items count.
func ByItemsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newItemsStep(), opts...)
	}
}

// ByItems orders the results by items terms.
func ByItems(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newItemsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCategoriesCount orders the results by categories count.
func ByCategoriesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCategoriesStep(), opts...)
	}
}

// ByCategories orders the results by categories terms.
func ByCategories(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCategoriesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserCount orders the results by user count.
func ByUserCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserStep(), opts...)
	}
}

// ByUser orders the results by user terms.
func ByUser(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserToStoreCount orders the results by UserToStore count.
func ByUserToStoreCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserToStoreStep(), opts...)
	}
}

// ByUserToStore orders the results by UserToStore terms.
func ByUserToStore(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserToStoreStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newItemsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ItemsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ItemsTable, ItemsColumn),
	)
}
func newCategoriesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CategoriesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CategoriesTable, CategoriesColumn),
	)
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, UserTable, UserPrimaryKey...),
	)
}
func newUserToStoreStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserToStoreInverseTable, UserToStoreColumn),
		sqlgraph.Edge(sqlgraph.O2M, true, UserToStoreTable, UserToStoreColumn),
	)
}
