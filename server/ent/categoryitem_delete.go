// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/hktrib/RetailGo/ent/categoryitem"
	"github.com/hktrib/RetailGo/ent/predicate"
)

// CategoryItemDelete is the builder for deleting a CategoryItem entity.
type CategoryItemDelete struct {
	config
	hooks    []Hook
	mutation *CategoryItemMutation
}

// Where appends a list predicates to the CategoryItemDelete builder.
func (cid *CategoryItemDelete) Where(ps ...predicate.CategoryItem) *CategoryItemDelete {
	cid.mutation.Where(ps...)
	return cid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cid *CategoryItemDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cid.sqlExec, cid.mutation, cid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cid *CategoryItemDelete) ExecX(ctx context.Context) int {
	n, err := cid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cid *CategoryItemDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(categoryitem.Table, nil)
	if ps := cid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cid.mutation.done = true
	return affected, err
}

// CategoryItemDeleteOne is the builder for deleting a single CategoryItem entity.
type CategoryItemDeleteOne struct {
	cid *CategoryItemDelete
}

// Where appends a list predicates to the CategoryItemDelete builder.
func (cido *CategoryItemDeleteOne) Where(ps ...predicate.CategoryItem) *CategoryItemDeleteOne {
	cido.cid.mutation.Where(ps...)
	return cido
}

// Exec executes the deletion query.
func (cido *CategoryItemDeleteOne) Exec(ctx context.Context) error {
	n, err := cido.cid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{categoryitem.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cido *CategoryItemDeleteOne) ExecX(ctx context.Context) {
	if err := cido.Exec(ctx); err != nil {
		panic(err)
	}
}
