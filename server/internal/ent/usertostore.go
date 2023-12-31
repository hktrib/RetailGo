// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/hktrib/RetailGo/internal/ent/store"
	"github.com/hktrib/RetailGo/internal/ent/user"
	"github.com/hktrib/RetailGo/internal/ent/usertostore"
)

// UserToStore is the model entity for the UserToStore schema.
type UserToStore struct {
	config `json:"-"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// StoreID holds the value of the "store_id" field.
	StoreID int `json:"store_id,omitempty"`
	// StoreName holds the value of the "store_name" field.
	StoreName string `json:"store_name,omitempty"`
	// ClerkUserID holds the value of the "clerk_user_id" field.
	ClerkUserID string `json:"clerk_user_id,omitempty"`
	// PermissionLevel holds the value of the "permission_level" field.
	PermissionLevel int `json:"permission_level,omitempty"`
	// JoinedAt holds the value of the "joined_at" field.
	JoinedAt int `json:"joined_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserToStoreQuery when eager-loading is set.
	Edges        UserToStoreEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserToStoreEdges holds the relations/edges for other nodes in the graph.
type UserToStoreEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Store holds the value of the store edge.
	Store *Store `json:"store,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserToStoreEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// StoreOrErr returns the Store value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserToStoreEdges) StoreOrErr() (*Store, error) {
	if e.loadedTypes[1] {
		if e.Store == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: store.Label}
		}
		return e.Store, nil
	}
	return nil, &NotLoadedError{edge: "store"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserToStore) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case usertostore.FieldUserID, usertostore.FieldStoreID, usertostore.FieldPermissionLevel, usertostore.FieldJoinedAt:
			values[i] = new(sql.NullInt64)
		case usertostore.FieldStoreName, usertostore.FieldClerkUserID:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserToStore fields.
func (uts *UserToStore) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case usertostore.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				uts.UserID = int(value.Int64)
			}
		case usertostore.FieldStoreID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field store_id", values[i])
			} else if value.Valid {
				uts.StoreID = int(value.Int64)
			}
		case usertostore.FieldStoreName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field store_name", values[i])
			} else if value.Valid {
				uts.StoreName = value.String
			}
		case usertostore.FieldClerkUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field clerk_user_id", values[i])
			} else if value.Valid {
				uts.ClerkUserID = value.String
			}
		case usertostore.FieldPermissionLevel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field permission_level", values[i])
			} else if value.Valid {
				uts.PermissionLevel = int(value.Int64)
			}
		case usertostore.FieldJoinedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field joined_at", values[i])
			} else if value.Valid {
				uts.JoinedAt = int(value.Int64)
			}
		default:
			uts.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserToStore.
// This includes values selected through modifiers, order, etc.
func (uts *UserToStore) Value(name string) (ent.Value, error) {
	return uts.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the UserToStore entity.
func (uts *UserToStore) QueryUser() *UserQuery {
	return NewUserToStoreClient(uts.config).QueryUser(uts)
}

// QueryStore queries the "store" edge of the UserToStore entity.
func (uts *UserToStore) QueryStore() *StoreQuery {
	return NewUserToStoreClient(uts.config).QueryStore(uts)
}

// Update returns a builder for updating this UserToStore.
// Note that you need to call UserToStore.Unwrap() before calling this method if this UserToStore
// was returned from a transaction, and the transaction was committed or rolled back.
func (uts *UserToStore) Update() *UserToStoreUpdateOne {
	return NewUserToStoreClient(uts.config).UpdateOne(uts)
}

// Unwrap unwraps the UserToStore entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (uts *UserToStore) Unwrap() *UserToStore {
	_tx, ok := uts.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserToStore is not a transactional entity")
	}
	uts.config.driver = _tx.drv
	return uts
}

// String implements the fmt.Stringer.
func (uts *UserToStore) String() string {
	var builder strings.Builder
	builder.WriteString("UserToStore(")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", uts.UserID))
	builder.WriteString(", ")
	builder.WriteString("store_id=")
	builder.WriteString(fmt.Sprintf("%v", uts.StoreID))
	builder.WriteString(", ")
	builder.WriteString("store_name=")
	builder.WriteString(uts.StoreName)
	builder.WriteString(", ")
	builder.WriteString("clerk_user_id=")
	builder.WriteString(uts.ClerkUserID)
	builder.WriteString(", ")
	builder.WriteString("permission_level=")
	builder.WriteString(fmt.Sprintf("%v", uts.PermissionLevel))
	builder.WriteString(", ")
	builder.WriteString("joined_at=")
	builder.WriteString(fmt.Sprintf("%v", uts.JoinedAt))
	builder.WriteByte(')')
	return builder.String()
}

// UserToStores is a parsable slice of UserToStore.
type UserToStores []*UserToStore
