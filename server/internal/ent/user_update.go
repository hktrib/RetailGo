// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hktrib/RetailGo/internal/ent/predicate"
	"github.com/hktrib/RetailGo/internal/ent/store"
	"github.com/hktrib/RetailGo/internal/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// SetIsOwner sets the "is_owner" field.
func (uu *UserUpdate) SetIsOwner(b bool) *UserUpdate {
	uu.mutation.SetIsOwner(b)
	return uu
}

// SetNillableIsOwner sets the "is_owner" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIsOwner(b *bool) *UserUpdate {
	if b != nil {
		uu.SetIsOwner(*b)
	}
	return uu
}

// SetStoreID sets the "store_id" field.
func (uu *UserUpdate) SetStoreID(i int) *UserUpdate {
	uu.mutation.ResetStoreID()
	uu.mutation.SetStoreID(i)
	return uu
}

// SetNillableStoreID sets the "store_id" field if the given value is not nil.
func (uu *UserUpdate) SetNillableStoreID(i *int) *UserUpdate {
	if i != nil {
		uu.SetStoreID(*i)
	}
	return uu
}

// AddStoreID adds i to the "store_id" field.
func (uu *UserUpdate) AddStoreID(i int) *UserUpdate {
	uu.mutation.AddStoreID(i)
	return uu
}

// ClearStoreID clears the value of the "store_id" field.
func (uu *UserUpdate) ClearStoreID() *UserUpdate {
	uu.mutation.ClearStoreID()
	return uu
}

// SetClerkUserID sets the "clerk_user_id" field.
func (uu *UserUpdate) SetClerkUserID(s string) *UserUpdate {
	uu.mutation.SetClerkUserID(s)
	return uu
}

// SetNillableClerkUserID sets the "clerk_user_id" field if the given value is not nil.
func (uu *UserUpdate) SetNillableClerkUserID(s *string) *UserUpdate {
	if s != nil {
		uu.SetClerkUserID(*s)
	}
	return uu
}

// ClearClerkUserID clears the value of the "clerk_user_id" field.
func (uu *UserUpdate) ClearClerkUserID() *UserUpdate {
	uu.mutation.ClearClerkUserID()
	return uu
}

// SetFirstName sets the "first_name" field.
func (uu *UserUpdate) SetFirstName(s string) *UserUpdate {
	uu.mutation.SetFirstName(s)
	return uu
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableFirstName(s *string) *UserUpdate {
	if s != nil {
		uu.SetFirstName(*s)
	}
	return uu
}

// ClearFirstName clears the value of the "first_name" field.
func (uu *UserUpdate) ClearFirstName() *UserUpdate {
	uu.mutation.ClearFirstName()
	return uu
}

// SetLastName sets the "last_name" field.
func (uu *UserUpdate) SetLastName(s string) *UserUpdate {
	uu.mutation.SetLastName(s)
	return uu
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableLastName(s *string) *UserUpdate {
	if s != nil {
		uu.SetLastName(*s)
	}
	return uu
}

// ClearLastName clears the value of the "last_name" field.
func (uu *UserUpdate) ClearLastName() *UserUpdate {
	uu.mutation.ClearLastName()
	return uu
}

// AddStoreIDs adds the "store" edge to the Store entity by IDs.
func (uu *UserUpdate) AddStoreIDs(ids ...int) *UserUpdate {
	uu.mutation.AddStoreIDs(ids...)
	return uu
}

// AddStore adds the "store" edges to the Store entity.
func (uu *UserUpdate) AddStore(s ...*Store) *UserUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uu.AddStoreIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearStore clears all "store" edges to the Store entity.
func (uu *UserUpdate) ClearStore() *UserUpdate {
	uu.mutation.ClearStore()
	return uu
}

// RemoveStoreIDs removes the "store" edge to Store entities by IDs.
func (uu *UserUpdate) RemoveStoreIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveStoreIDs(ids...)
	return uu
}

// RemoveStore removes "store" edges to Store entities.
func (uu *UserUpdate) RemoveStore(s ...*Store) *UserUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uu.RemoveStoreIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.IsOwner(); ok {
		_spec.SetField(user.FieldIsOwner, field.TypeBool, value)
	}
	if value, ok := uu.mutation.StoreID(); ok {
		_spec.SetField(user.FieldStoreID, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedStoreID(); ok {
		_spec.AddField(user.FieldStoreID, field.TypeInt, value)
	}
	if uu.mutation.StoreIDCleared() {
		_spec.ClearField(user.FieldStoreID, field.TypeInt)
	}
	if value, ok := uu.mutation.ClerkUserID(); ok {
		_spec.SetField(user.FieldClerkUserID, field.TypeString, value)
	}
	if uu.mutation.ClerkUserIDCleared() {
		_spec.ClearField(user.FieldClerkUserID, field.TypeString)
	}
	if value, ok := uu.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
	}
	if uu.mutation.FirstNameCleared() {
		_spec.ClearField(user.FieldFirstName, field.TypeString)
	}
	if value, ok := uu.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
	}
	if uu.mutation.LastNameCleared() {
		_spec.ClearField(user.FieldLastName, field.TypeString)
	}
	if uu.mutation.StoreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.StoreTable,
			Columns: user.StorePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(store.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedStoreIDs(); len(nodes) > 0 && !uu.mutation.StoreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.StoreTable,
			Columns: user.StorePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(store.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.StoreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.StoreTable,
			Columns: user.StorePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(store.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// SetIsOwner sets the "is_owner" field.
func (uuo *UserUpdateOne) SetIsOwner(b bool) *UserUpdateOne {
	uuo.mutation.SetIsOwner(b)
	return uuo
}

// SetNillableIsOwner sets the "is_owner" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIsOwner(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetIsOwner(*b)
	}
	return uuo
}

// SetStoreID sets the "store_id" field.
func (uuo *UserUpdateOne) SetStoreID(i int) *UserUpdateOne {
	uuo.mutation.ResetStoreID()
	uuo.mutation.SetStoreID(i)
	return uuo
}

// SetNillableStoreID sets the "store_id" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableStoreID(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetStoreID(*i)
	}
	return uuo
}

// AddStoreID adds i to the "store_id" field.
func (uuo *UserUpdateOne) AddStoreID(i int) *UserUpdateOne {
	uuo.mutation.AddStoreID(i)
	return uuo
}

// ClearStoreID clears the value of the "store_id" field.
func (uuo *UserUpdateOne) ClearStoreID() *UserUpdateOne {
	uuo.mutation.ClearStoreID()
	return uuo
}

// SetClerkUserID sets the "clerk_user_id" field.
func (uuo *UserUpdateOne) SetClerkUserID(s string) *UserUpdateOne {
	uuo.mutation.SetClerkUserID(s)
	return uuo
}

// SetNillableClerkUserID sets the "clerk_user_id" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableClerkUserID(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetClerkUserID(*s)
	}
	return uuo
}

// ClearClerkUserID clears the value of the "clerk_user_id" field.
func (uuo *UserUpdateOne) ClearClerkUserID() *UserUpdateOne {
	uuo.mutation.ClearClerkUserID()
	return uuo
}

// SetFirstName sets the "first_name" field.
func (uuo *UserUpdateOne) SetFirstName(s string) *UserUpdateOne {
	uuo.mutation.SetFirstName(s)
	return uuo
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableFirstName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetFirstName(*s)
	}
	return uuo
}

// ClearFirstName clears the value of the "first_name" field.
func (uuo *UserUpdateOne) ClearFirstName() *UserUpdateOne {
	uuo.mutation.ClearFirstName()
	return uuo
}

// SetLastName sets the "last_name" field.
func (uuo *UserUpdateOne) SetLastName(s string) *UserUpdateOne {
	uuo.mutation.SetLastName(s)
	return uuo
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLastName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetLastName(*s)
	}
	return uuo
}

// ClearLastName clears the value of the "last_name" field.
func (uuo *UserUpdateOne) ClearLastName() *UserUpdateOne {
	uuo.mutation.ClearLastName()
	return uuo
}

// AddStoreIDs adds the "store" edge to the Store entity by IDs.
func (uuo *UserUpdateOne) AddStoreIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddStoreIDs(ids...)
	return uuo
}

// AddStore adds the "store" edges to the Store entity.
func (uuo *UserUpdateOne) AddStore(s ...*Store) *UserUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uuo.AddStoreIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearStore clears all "store" edges to the Store entity.
func (uuo *UserUpdateOne) ClearStore() *UserUpdateOne {
	uuo.mutation.ClearStore()
	return uuo
}

// RemoveStoreIDs removes the "store" edge to Store entities by IDs.
func (uuo *UserUpdateOne) RemoveStoreIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveStoreIDs(ids...)
	return uuo
}

// RemoveStore removes "store" edges to Store entities.
func (uuo *UserUpdateOne) RemoveStore(s ...*Store) *UserUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uuo.RemoveStoreIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.IsOwner(); ok {
		_spec.SetField(user.FieldIsOwner, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.StoreID(); ok {
		_spec.SetField(user.FieldStoreID, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedStoreID(); ok {
		_spec.AddField(user.FieldStoreID, field.TypeInt, value)
	}
	if uuo.mutation.StoreIDCleared() {
		_spec.ClearField(user.FieldStoreID, field.TypeInt)
	}
	if value, ok := uuo.mutation.ClerkUserID(); ok {
		_spec.SetField(user.FieldClerkUserID, field.TypeString, value)
	}
	if uuo.mutation.ClerkUserIDCleared() {
		_spec.ClearField(user.FieldClerkUserID, field.TypeString)
	}
	if value, ok := uuo.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
	}
	if uuo.mutation.FirstNameCleared() {
		_spec.ClearField(user.FieldFirstName, field.TypeString)
	}
	if value, ok := uuo.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
	}
	if uuo.mutation.LastNameCleared() {
		_spec.ClearField(user.FieldLastName, field.TypeString)
	}
	if uuo.mutation.StoreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.StoreTable,
			Columns: user.StorePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(store.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedStoreIDs(); len(nodes) > 0 && !uuo.mutation.StoreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.StoreTable,
			Columns: user.StorePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(store.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.StoreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.StoreTable,
			Columns: user.StorePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(store.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
