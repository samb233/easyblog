// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samb233/easyblog/internal/repo/ent/index"
)

// IndexCreate is the builder for creating a Index entity.
type IndexCreate struct {
	config
	mutation *IndexMutation
	hooks    []Hook
}

// SetContentID sets the "content_id" field.
func (ic *IndexCreate) SetContentID(i int32) *IndexCreate {
	ic.mutation.SetContentID(i)
	return ic
}

// SetTitle sets the "title" field.
func (ic *IndexCreate) SetTitle(s string) *IndexCreate {
	ic.mutation.SetTitle(s)
	return ic
}

// SetDesc sets the "desc" field.
func (ic *IndexCreate) SetDesc(s string) *IndexCreate {
	ic.mutation.SetDesc(s)
	return ic
}

// SetAttr sets the "attr" field.
func (ic *IndexCreate) SetAttr(i int32) *IndexCreate {
	ic.mutation.SetAttr(i)
	return ic
}

// SetNillableAttr sets the "attr" field if the given value is not nil.
func (ic *IndexCreate) SetNillableAttr(i *int32) *IndexCreate {
	if i != nil {
		ic.SetAttr(*i)
	}
	return ic
}

// SetView sets the "view" field.
func (ic *IndexCreate) SetView(i int32) *IndexCreate {
	ic.mutation.SetView(i)
	return ic
}

// SetNillableView sets the "view" field if the given value is not nil.
func (ic *IndexCreate) SetNillableView(i *int32) *IndexCreate {
	if i != nil {
		ic.SetView(*i)
	}
	return ic
}

// SetCreatedAt sets the "created_at" field.
func (ic *IndexCreate) SetCreatedAt(t time.Time) *IndexCreate {
	ic.mutation.SetCreatedAt(t)
	return ic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ic *IndexCreate) SetNillableCreatedAt(t *time.Time) *IndexCreate {
	if t != nil {
		ic.SetCreatedAt(*t)
	}
	return ic
}

// SetUpdatedAt sets the "updated_at" field.
func (ic *IndexCreate) SetUpdatedAt(t time.Time) *IndexCreate {
	ic.mutation.SetUpdatedAt(t)
	return ic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ic *IndexCreate) SetNillableUpdatedAt(t *time.Time) *IndexCreate {
	if t != nil {
		ic.SetUpdatedAt(*t)
	}
	return ic
}

// SetState sets the "state" field.
func (ic *IndexCreate) SetState(i int8) *IndexCreate {
	ic.mutation.SetState(i)
	return ic
}

// SetNillableState sets the "state" field if the given value is not nil.
func (ic *IndexCreate) SetNillableState(i *int8) *IndexCreate {
	if i != nil {
		ic.SetState(*i)
	}
	return ic
}

// SetID sets the "id" field.
func (ic *IndexCreate) SetID(i int32) *IndexCreate {
	ic.mutation.SetID(i)
	return ic
}

// Mutation returns the IndexMutation object of the builder.
func (ic *IndexCreate) Mutation() *IndexMutation {
	return ic.mutation
}

// Save creates the Index in the database.
func (ic *IndexCreate) Save(ctx context.Context) (*Index, error) {
	var (
		err  error
		node *Index
	)
	ic.defaults()
	if len(ic.hooks) == 0 {
		if err = ic.check(); err != nil {
			return nil, err
		}
		node, err = ic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IndexMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ic.check(); err != nil {
				return nil, err
			}
			ic.mutation = mutation
			if node, err = ic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ic.hooks) - 1; i >= 0; i-- {
			if ic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ic *IndexCreate) SaveX(ctx context.Context) *Index {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *IndexCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *IndexCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *IndexCreate) defaults() {
	if _, ok := ic.mutation.Attr(); !ok {
		v := index.DefaultAttr
		ic.mutation.SetAttr(v)
	}
	if _, ok := ic.mutation.View(); !ok {
		v := index.DefaultView
		ic.mutation.SetView(v)
	}
	if _, ok := ic.mutation.CreatedAt(); !ok {
		v := index.DefaultCreatedAt()
		ic.mutation.SetCreatedAt(v)
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		v := index.DefaultUpdatedAt()
		ic.mutation.SetUpdatedAt(v)
	}
	if _, ok := ic.mutation.State(); !ok {
		v := index.DefaultState
		ic.mutation.SetState(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *IndexCreate) check() error {
	if _, ok := ic.mutation.ContentID(); !ok {
		return &ValidationError{Name: "content_id", err: errors.New(`ent: missing required field "content_id"`)}
	}
	if _, ok := ic.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "title"`)}
	}
	if v, ok := ic.mutation.Title(); ok {
		if err := index.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "title": %w`, err)}
		}
	}
	if _, ok := ic.mutation.Desc(); !ok {
		return &ValidationError{Name: "desc", err: errors.New(`ent: missing required field "desc"`)}
	}
	if _, ok := ic.mutation.Attr(); !ok {
		return &ValidationError{Name: "attr", err: errors.New(`ent: missing required field "attr"`)}
	}
	if _, ok := ic.mutation.View(); !ok {
		return &ValidationError{Name: "view", err: errors.New(`ent: missing required field "view"`)}
	}
	if _, ok := ic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "updated_at"`)}
	}
	if _, ok := ic.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "state"`)}
	}
	return nil
}

func (ic *IndexCreate) sqlSave(ctx context.Context) (*Index, error) {
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	return _node, nil
}

func (ic *IndexCreate) createSpec() (*Index, *sqlgraph.CreateSpec) {
	var (
		_node = &Index{config: ic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: index.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: index.FieldID,
			},
		}
	)
	if id, ok := ic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ic.mutation.ContentID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: index.FieldContentID,
		})
		_node.ContentID = value
	}
	if value, ok := ic.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: index.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := ic.mutation.Desc(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: index.FieldDesc,
		})
		_node.Desc = value
	}
	if value, ok := ic.mutation.Attr(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: index.FieldAttr,
		})
		_node.Attr = value
	}
	if value, ok := ic.mutation.View(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: index.FieldView,
		})
		_node.View = value
	}
	if value, ok := ic.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: index.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ic.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: index.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ic.mutation.State(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: index.FieldState,
		})
		_node.State = value
	}
	return _node, _spec
}

// IndexCreateBulk is the builder for creating many Index entities in bulk.
type IndexCreateBulk struct {
	config
	builders []*IndexCreate
}

// Save creates the Index entities in the database.
func (icb *IndexCreateBulk) Save(ctx context.Context) ([]*Index, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Index, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*IndexMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int32(id)
				}
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
func (icb *IndexCreateBulk) SaveX(ctx context.Context) []*Index {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *IndexCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *IndexCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
