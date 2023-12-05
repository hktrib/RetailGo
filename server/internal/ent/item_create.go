// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hktrib/RetailGo/internal/ent/category"
	"github.com/hktrib/RetailGo/internal/ent/item"
	"github.com/hktrib/RetailGo/internal/ent/store"
)

// ItemCreate is the builder for creating a Item entity.
type ItemCreate struct {
	config
	mutation *ItemMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ic *ItemCreate) SetName(s string) *ItemCreate {
	ic.mutation.SetName(s)
	return ic
}

// SetPhoto sets the "photo" field.
func (ic *ItemCreate) SetPhoto(s string) *ItemCreate {
	ic.mutation.SetPhoto(s)
	return ic
}

// SetQuantity sets the "quantity" field.
func (ic *ItemCreate) SetQuantity(i int) *ItemCreate {
	ic.mutation.SetQuantity(i)
	return ic
}

// SetPrice sets the "price" field.
func (ic *ItemCreate) SetPrice(f float64) *ItemCreate {
	ic.mutation.SetPrice(f)
	return ic
}

// SetStoreID sets the "store_id" field.
func (ic *ItemCreate) SetStoreID(i int) *ItemCreate {
	ic.mutation.SetStoreID(i)
	return ic
}

// SetStripePriceID sets the "stripe_price_id" field.
func (ic *ItemCreate) SetStripePriceID(s string) *ItemCreate {
	ic.mutation.SetStripePriceID(s)
	return ic
}

// SetNillableStripePriceID sets the "stripe_price_id" field if the given value is not nil.
func (ic *ItemCreate) SetNillableStripePriceID(s *string) *ItemCreate {
	if s != nil {
		ic.SetStripePriceID(*s)
	}
	return ic
}

// SetCategoryName sets the "category_name" field.
func (ic *ItemCreate) SetCategoryName(s string) *ItemCreate {
	ic.mutation.SetCategoryName(s)
	return ic
}

// SetNillableCategoryName sets the "category_name" field if the given value is not nil.
func (ic *ItemCreate) SetNillableCategoryName(s *string) *ItemCreate {
	if s != nil {
		ic.SetCategoryName(*s)
	}
	return ic
}

// SetStripeProductID sets the "stripe_product_id" field.
func (ic *ItemCreate) SetStripeProductID(s string) *ItemCreate {
	ic.mutation.SetStripeProductID(s)
	return ic
}

// SetNillableStripeProductID sets the "stripe_product_id" field if the given value is not nil.
func (ic *ItemCreate) SetNillableStripeProductID(s *string) *ItemCreate {
	if s != nil {
		ic.SetStripeProductID(*s)
	}
	return ic
}

// SetWeaviateID sets the "weaviate_id" field.
func (ic *ItemCreate) SetWeaviateID(s string) *ItemCreate {
	ic.mutation.SetWeaviateID(s)
	return ic
}

// SetNillableWeaviateID sets the "weaviate_id" field if the given value is not nil.
func (ic *ItemCreate) SetNillableWeaviateID(s *string) *ItemCreate {
	if s != nil {
		ic.SetWeaviateID(*s)
	}
	return ic
}

// SetVectorized sets the "vectorized" field.
func (ic *ItemCreate) SetVectorized(b bool) *ItemCreate {
	ic.mutation.SetVectorized(b)
	return ic
}

// SetNillableVectorized sets the "vectorized" field if the given value is not nil.
func (ic *ItemCreate) SetNillableVectorized(b *bool) *ItemCreate {
	if b != nil {
		ic.SetVectorized(*b)
	}
	return ic
}

// SetNumberSoldSinceUpdate sets the "number_sold_since_update" field.
func (ic *ItemCreate) SetNumberSoldSinceUpdate(i int) *ItemCreate {
	ic.mutation.SetNumberSoldSinceUpdate(i)
	return ic
}

// SetNillableNumberSoldSinceUpdate sets the "number_sold_since_update" field if the given value is not nil.
func (ic *ItemCreate) SetNillableNumberSoldSinceUpdate(i *int) *ItemCreate {
	if i != nil {
		ic.SetNumberSoldSinceUpdate(*i)
	}
	return ic
}

// SetDateLastSold sets the "date_last_sold" field.
func (ic *ItemCreate) SetDateLastSold(s string) *ItemCreate {
	ic.mutation.SetDateLastSold(s)
	return ic
}

// SetNillableDateLastSold sets the "date_last_sold" field if the given value is not nil.
func (ic *ItemCreate) SetNillableDateLastSold(s *string) *ItemCreate {
	if s != nil {
		ic.SetDateLastSold(*s)
	}
	return ic
}

// SetID sets the "id" field.
func (ic *ItemCreate) SetID(i int) *ItemCreate {
	ic.mutation.SetID(i)
	return ic
}

// AddCategoryIDs adds the "category" edge to the Category entity by IDs.
func (ic *ItemCreate) AddCategoryIDs(ids ...int) *ItemCreate {
	ic.mutation.AddCategoryIDs(ids...)
	return ic
}

// AddCategory adds the "category" edges to the Category entity.
func (ic *ItemCreate) AddCategory(c ...*Category) *ItemCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ic.AddCategoryIDs(ids...)
}

// SetStore sets the "store" edge to the Store entity.
func (ic *ItemCreate) SetStore(s *Store) *ItemCreate {
	return ic.SetStoreID(s.ID)
}

// Mutation returns the ItemMutation object of the builder.
func (ic *ItemCreate) Mutation() *ItemMutation {
	return ic.mutation
}

// Save creates the Item in the database.
func (ic *ItemCreate) Save(ctx context.Context) (*Item, error) {
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *ItemCreate) SaveX(ctx context.Context) *Item {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *ItemCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *ItemCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *ItemCreate) check() error {
	if _, ok := ic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Item.name"`)}
	}
	if _, ok := ic.mutation.Photo(); !ok {
		return &ValidationError{Name: "photo", err: errors.New(`ent: missing required field "Item.photo"`)}
	}
	if _, ok := ic.mutation.Quantity(); !ok {
		return &ValidationError{Name: "quantity", err: errors.New(`ent: missing required field "Item.quantity"`)}
	}
	if _, ok := ic.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "Item.price"`)}
	}
	if _, ok := ic.mutation.StoreID(); !ok {
		return &ValidationError{Name: "store_id", err: errors.New(`ent: missing required field "Item.store_id"`)}
	}
	if _, ok := ic.mutation.StoreID(); !ok {
		return &ValidationError{Name: "store", err: errors.New(`ent: missing required edge "Item.store"`)}
	}
	return nil
}

func (ic *ItemCreate) sqlSave(ctx context.Context) (*Item, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *ItemCreate) createSpec() (*Item, *sqlgraph.CreateSpec) {
	var (
		_node = &Item{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(item.Table, sqlgraph.NewFieldSpec(item.FieldID, field.TypeInt))
	)
	if id, ok := ic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ic.mutation.Name(); ok {
		_spec.SetField(item.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ic.mutation.Photo(); ok {
		_spec.SetField(item.FieldPhoto, field.TypeString, value)
		_node.Photo = value
	}
	if value, ok := ic.mutation.Quantity(); ok {
		_spec.SetField(item.FieldQuantity, field.TypeInt, value)
		_node.Quantity = value
	}
	if value, ok := ic.mutation.Price(); ok {
		_spec.SetField(item.FieldPrice, field.TypeFloat64, value)
		_node.Price = value
	}
	if value, ok := ic.mutation.StripePriceID(); ok {
		_spec.SetField(item.FieldStripePriceID, field.TypeString, value)
		_node.StripePriceID = value
	}
	if value, ok := ic.mutation.CategoryName(); ok {
		_spec.SetField(item.FieldCategoryName, field.TypeString, value)
		_node.CategoryName = value
	}
	if value, ok := ic.mutation.StripeProductID(); ok {
		_spec.SetField(item.FieldStripeProductID, field.TypeString, value)
		_node.StripeProductID = value
	}
	if value, ok := ic.mutation.WeaviateID(); ok {
		_spec.SetField(item.FieldWeaviateID, field.TypeString, value)
		_node.WeaviateID = value
	}
	if value, ok := ic.mutation.Vectorized(); ok {
		_spec.SetField(item.FieldVectorized, field.TypeBool, value)
		_node.Vectorized = value
	}
	if value, ok := ic.mutation.NumberSoldSinceUpdate(); ok {
		_spec.SetField(item.FieldNumberSoldSinceUpdate, field.TypeInt, value)
		_node.NumberSoldSinceUpdate = value
	}
	if value, ok := ic.mutation.DateLastSold(); ok {
		_spec.SetField(item.FieldDateLastSold, field.TypeString, value)
		_node.DateLastSold = value
	}
	if nodes := ic.mutation.CategoryIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.StoreIDs(); len(nodes) > 0 {
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
		_node.StoreID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ItemCreateBulk is the builder for creating many Item entities in bulk.
type ItemCreateBulk struct {
	config
	err      error
	builders []*ItemCreate
}

// Save creates the Item entities in the database.
func (icb *ItemCreateBulk) Save(ctx context.Context) ([]*Item, error) {
	if icb.err != nil {
		return nil, icb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Item, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ItemMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *ItemCreateBulk) SaveX(ctx context.Context) []*Item {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *ItemCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *ItemCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
