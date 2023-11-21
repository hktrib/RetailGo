// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hktrib/RetailGo/internal/ent/category"
	"github.com/hktrib/RetailGo/internal/ent/item"
	"github.com/hktrib/RetailGo/internal/ent/predicate"
	"github.com/hktrib/RetailGo/internal/ent/store"
)

// ItemUpdate is the builder for updating Item entities.
type ItemUpdate struct {
	config
	hooks    []Hook
	mutation *ItemMutation
}

// Where appends a list predicates to the ItemUpdate builder.
func (iu *ItemUpdate) Where(ps ...predicate.Item) *ItemUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetName sets the "name" field.
func (iu *ItemUpdate) SetName(s string) *ItemUpdate {
	iu.mutation.SetName(s)
	return iu
}

// SetPhoto sets the "photo" field.
func (iu *ItemUpdate) SetPhoto(b []byte) *ItemUpdate {
	iu.mutation.SetPhoto(b)
	return iu
}

// SetQuantity sets the "quantity" field.
func (iu *ItemUpdate) SetQuantity(i int) *ItemUpdate {
	iu.mutation.ResetQuantity()
	iu.mutation.SetQuantity(i)
	return iu
}

// AddQuantity adds i to the "quantity" field.
func (iu *ItemUpdate) AddQuantity(i int) *ItemUpdate {
	iu.mutation.AddQuantity(i)
	return iu
}

// SetPrice sets the "price" field.
func (iu *ItemUpdate) SetPrice(f float64) *ItemUpdate {
	iu.mutation.ResetPrice()
	iu.mutation.SetPrice(f)
	return iu
}

// AddPrice adds f to the "price" field.
func (iu *ItemUpdate) AddPrice(f float64) *ItemUpdate {
	iu.mutation.AddPrice(f)
	return iu
}

// SetStoreID sets the "store_id" field.
func (iu *ItemUpdate) SetStoreID(i int) *ItemUpdate {
	iu.mutation.SetStoreID(i)
	return iu
}

// SetStripePriceID sets the "stripe_price_id" field.
func (iu *ItemUpdate) SetStripePriceID(s string) *ItemUpdate {
	iu.mutation.SetStripePriceID(s)
	return iu
}

// SetStripeProductID sets the "stripe_product_id" field.
func (iu *ItemUpdate) SetStripeProductID(s string) *ItemUpdate {
	iu.mutation.SetStripeProductID(s)
	return iu
}

// AddCategoryIDs adds the "category" edge to the Category entity by IDs.
func (iu *ItemUpdate) AddCategoryIDs(ids ...int) *ItemUpdate {
	iu.mutation.AddCategoryIDs(ids...)
	return iu
}

// AddCategory adds the "category" edges to the Category entity.
func (iu *ItemUpdate) AddCategory(c ...*Category) *ItemUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return iu.AddCategoryIDs(ids...)
}

// SetStore sets the "store" edge to the Store entity.
func (iu *ItemUpdate) SetStore(s *Store) *ItemUpdate {
	return iu.SetStoreID(s.ID)
}

// Mutation returns the ItemMutation object of the builder.
func (iu *ItemUpdate) Mutation() *ItemMutation {
	return iu.mutation
}

// ClearCategory clears all "category" edges to the Category entity.
func (iu *ItemUpdate) ClearCategory() *ItemUpdate {
	iu.mutation.ClearCategory()
	return iu
}

// RemoveCategoryIDs removes the "category" edge to Category entities by IDs.
func (iu *ItemUpdate) RemoveCategoryIDs(ids ...int) *ItemUpdate {
	iu.mutation.RemoveCategoryIDs(ids...)
	return iu
}

// RemoveCategory removes "category" edges to Category entities.
func (iu *ItemUpdate) RemoveCategory(c ...*Category) *ItemUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return iu.RemoveCategoryIDs(ids...)
}

// ClearStore clears the "store" edge to the Store entity.
func (iu *ItemUpdate) ClearStore() *ItemUpdate {
	iu.mutation.ClearStore()
	return iu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *ItemUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *ItemUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *ItemUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *ItemUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *ItemUpdate) check() error {
	if _, ok := iu.mutation.StoreID(); iu.mutation.StoreCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Item.store"`)
	}
	return nil
}

func (iu *ItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := iu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(item.Table, item.Columns, sqlgraph.NewFieldSpec(item.FieldID, field.TypeInt))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.Name(); ok {
		_spec.SetField(item.FieldName, field.TypeString, value)
	}
	if value, ok := iu.mutation.Photo(); ok {
		_spec.SetField(item.FieldPhoto, field.TypeBytes, value)
	}
	if value, ok := iu.mutation.Quantity(); ok {
		_spec.SetField(item.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := iu.mutation.AddedQuantity(); ok {
		_spec.AddField(item.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := iu.mutation.Price(); ok {
		_spec.SetField(item.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := iu.mutation.AddedPrice(); ok {
		_spec.AddField(item.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := iu.mutation.StripePriceID(); ok {
		_spec.SetField(item.FieldStripePriceID, field.TypeString, value)
	}
	if value, ok := iu.mutation.StripeProductID(); ok {
		_spec.SetField(item.FieldStripeProductID, field.TypeString, value)
	}
	if iu.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   item.CategoryTable,
			Columns: item.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedCategoryIDs(); len(nodes) > 0 && !iu.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   item.CategoryTable,
			Columns: item.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   item.CategoryTable,
			Columns: item.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.StoreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.StoreTable,
			Columns: []string{item.StoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(store.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.StoreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.StoreTable,
			Columns: []string{item.StoreColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{item.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// ItemUpdateOne is the builder for updating a single Item entity.
type ItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ItemMutation
}

// SetName sets the "name" field.
func (iuo *ItemUpdateOne) SetName(s string) *ItemUpdateOne {
	iuo.mutation.SetName(s)
	return iuo
}

// SetPhoto sets the "photo" field.
func (iuo *ItemUpdateOne) SetPhoto(b []byte) *ItemUpdateOne {
	iuo.mutation.SetPhoto(b)
	return iuo
}

// SetQuantity sets the "quantity" field.
func (iuo *ItemUpdateOne) SetQuantity(i int) *ItemUpdateOne {
	iuo.mutation.ResetQuantity()
	iuo.mutation.SetQuantity(i)
	return iuo
}

// AddQuantity adds i to the "quantity" field.
func (iuo *ItemUpdateOne) AddQuantity(i int) *ItemUpdateOne {
	iuo.mutation.AddQuantity(i)
	return iuo
}

// SetPrice sets the "price" field.
func (iuo *ItemUpdateOne) SetPrice(f float64) *ItemUpdateOne {
	iuo.mutation.ResetPrice()
	iuo.mutation.SetPrice(f)
	return iuo
}

// AddPrice adds f to the "price" field.
func (iuo *ItemUpdateOne) AddPrice(f float64) *ItemUpdateOne {
	iuo.mutation.AddPrice(f)
	return iuo
}

// SetStoreID sets the "store_id" field.
func (iuo *ItemUpdateOne) SetStoreID(i int) *ItemUpdateOne {
	iuo.mutation.SetStoreID(i)
	return iuo
}

// SetStripePriceID sets the "stripe_price_id" field.
func (iuo *ItemUpdateOne) SetStripePriceID(s string) *ItemUpdateOne {
	iuo.mutation.SetStripePriceID(s)
	return iuo
}

// SetStripeProductID sets the "stripe_product_id" field.
func (iuo *ItemUpdateOne) SetStripeProductID(s string) *ItemUpdateOne {
	iuo.mutation.SetStripeProductID(s)
	return iuo
}

// AddCategoryIDs adds the "category" edge to the Category entity by IDs.
func (iuo *ItemUpdateOne) AddCategoryIDs(ids ...int) *ItemUpdateOne {
	iuo.mutation.AddCategoryIDs(ids...)
	return iuo
}

// AddCategory adds the "category" edges to the Category entity.
func (iuo *ItemUpdateOne) AddCategory(c ...*Category) *ItemUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return iuo.AddCategoryIDs(ids...)
}

// SetStore sets the "store" edge to the Store entity.
func (iuo *ItemUpdateOne) SetStore(s *Store) *ItemUpdateOne {
	return iuo.SetStoreID(s.ID)
}

// Mutation returns the ItemMutation object of the builder.
func (iuo *ItemUpdateOne) Mutation() *ItemMutation {
	return iuo.mutation
}

// ClearCategory clears all "category" edges to the Category entity.
func (iuo *ItemUpdateOne) ClearCategory() *ItemUpdateOne {
	iuo.mutation.ClearCategory()
	return iuo
}

// RemoveCategoryIDs removes the "category" edge to Category entities by IDs.
func (iuo *ItemUpdateOne) RemoveCategoryIDs(ids ...int) *ItemUpdateOne {
	iuo.mutation.RemoveCategoryIDs(ids...)
	return iuo
}

// RemoveCategory removes "category" edges to Category entities.
func (iuo *ItemUpdateOne) RemoveCategory(c ...*Category) *ItemUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return iuo.RemoveCategoryIDs(ids...)
}

// ClearStore clears the "store" edge to the Store entity.
func (iuo *ItemUpdateOne) ClearStore() *ItemUpdateOne {
	iuo.mutation.ClearStore()
	return iuo
}

// Where appends a list predicates to the ItemUpdate builder.
func (iuo *ItemUpdateOne) Where(ps ...predicate.Item) *ItemUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *ItemUpdateOne) Select(field string, fields ...string) *ItemUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Item entity.
func (iuo *ItemUpdateOne) Save(ctx context.Context) (*Item, error) {
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *ItemUpdateOne) SaveX(ctx context.Context) *Item {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *ItemUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *ItemUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *ItemUpdateOne) check() error {
	if _, ok := iuo.mutation.StoreID(); iuo.mutation.StoreCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Item.store"`)
	}
	return nil
}

func (iuo *ItemUpdateOne) sqlSave(ctx context.Context) (_node *Item, err error) {
	if err := iuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(item.Table, item.Columns, sqlgraph.NewFieldSpec(item.FieldID, field.TypeInt))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Item.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, item.FieldID)
		for _, f := range fields {
			if !item.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != item.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.Name(); ok {
		_spec.SetField(item.FieldName, field.TypeString, value)
	}
	if value, ok := iuo.mutation.Photo(); ok {
		_spec.SetField(item.FieldPhoto, field.TypeBytes, value)
	}
	if value, ok := iuo.mutation.Quantity(); ok {
		_spec.SetField(item.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := iuo.mutation.AddedQuantity(); ok {
		_spec.AddField(item.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := iuo.mutation.Price(); ok {
		_spec.SetField(item.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := iuo.mutation.AddedPrice(); ok {
		_spec.AddField(item.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := iuo.mutation.StripePriceID(); ok {
		_spec.SetField(item.FieldStripePriceID, field.TypeString, value)
	}
	if value, ok := iuo.mutation.StripeProductID(); ok {
		_spec.SetField(item.FieldStripeProductID, field.TypeString, value)
	}
	if iuo.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   item.CategoryTable,
			Columns: item.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedCategoryIDs(); len(nodes) > 0 && !iuo.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   item.CategoryTable,
			Columns: item.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   item.CategoryTable,
			Columns: item.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.StoreCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.StoreTable,
			Columns: []string{item.StoreColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(store.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.StoreIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.StoreTable,
			Columns: []string{item.StoreColumn},
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
	_node = &Item{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{item.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}
