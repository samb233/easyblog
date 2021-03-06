// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/samb233/easyblog/internal/repo/ent/index"
	"github.com/samb233/easyblog/internal/repo/ent/predicate"
)

// IndexUpdate is the builder for updating Index entities.
type IndexUpdate struct {
	config
	hooks    []Hook
	mutation *IndexMutation
}

// Where appends a list predicates to the IndexUpdate builder.
func (iu *IndexUpdate) Where(ps ...predicate.Index) *IndexUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetContentID sets the "content_id" field.
func (iu *IndexUpdate) SetContentID(i int32) *IndexUpdate {
	iu.mutation.ResetContentID()
	iu.mutation.SetContentID(i)
	return iu
}

// AddContentID adds i to the "content_id" field.
func (iu *IndexUpdate) AddContentID(i int32) *IndexUpdate {
	iu.mutation.AddContentID(i)
	return iu
}

// SetTitle sets the "title" field.
func (iu *IndexUpdate) SetTitle(s string) *IndexUpdate {
	iu.mutation.SetTitle(s)
	return iu
}

// SetDesc sets the "desc" field.
func (iu *IndexUpdate) SetDesc(s string) *IndexUpdate {
	iu.mutation.SetDesc(s)
	return iu
}

// SetAttr sets the "attr" field.
func (iu *IndexUpdate) SetAttr(i int32) *IndexUpdate {
	iu.mutation.ResetAttr()
	iu.mutation.SetAttr(i)
	return iu
}

// SetNillableAttr sets the "attr" field if the given value is not nil.
func (iu *IndexUpdate) SetNillableAttr(i *int32) *IndexUpdate {
	if i != nil {
		iu.SetAttr(*i)
	}
	return iu
}

// AddAttr adds i to the "attr" field.
func (iu *IndexUpdate) AddAttr(i int32) *IndexUpdate {
	iu.mutation.AddAttr(i)
	return iu
}

// SetView sets the "view" field.
func (iu *IndexUpdate) SetView(i int32) *IndexUpdate {
	iu.mutation.ResetView()
	iu.mutation.SetView(i)
	return iu
}

// SetNillableView sets the "view" field if the given value is not nil.
func (iu *IndexUpdate) SetNillableView(i *int32) *IndexUpdate {
	if i != nil {
		iu.SetView(*i)
	}
	return iu
}

// AddView adds i to the "view" field.
func (iu *IndexUpdate) AddView(i int32) *IndexUpdate {
	iu.mutation.AddView(i)
	return iu
}

// SetUpdatedAt sets the "updated_at" field.
func (iu *IndexUpdate) SetUpdatedAt(t time.Time) *IndexUpdate {
	iu.mutation.SetUpdatedAt(t)
	return iu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (iu *IndexUpdate) SetNillableUpdatedAt(t *time.Time) *IndexUpdate {
	if t != nil {
		iu.SetUpdatedAt(*t)
	}
	return iu
}

// SetState sets the "state" field.
func (iu *IndexUpdate) SetState(i int8) *IndexUpdate {
	iu.mutation.ResetState()
	iu.mutation.SetState(i)
	return iu
}

// SetNillableState sets the "state" field if the given value is not nil.
func (iu *IndexUpdate) SetNillableState(i *int8) *IndexUpdate {
	if i != nil {
		iu.SetState(*i)
	}
	return iu
}

// AddState adds i to the "state" field.
func (iu *IndexUpdate) AddState(i int8) *IndexUpdate {
	iu.mutation.AddState(i)
	return iu
}

// Mutation returns the IndexMutation object of the builder.
func (iu *IndexUpdate) Mutation() *IndexMutation {
	return iu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *IndexUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(iu.hooks) == 0 {
		if err = iu.check(); err != nil {
			return 0, err
		}
		affected, err = iu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IndexMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iu.check(); err != nil {
				return 0, err
			}
			iu.mutation = mutation
			affected, err = iu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iu.hooks) - 1; i >= 0; i-- {
			if iu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iu *IndexUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *IndexUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *IndexUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *IndexUpdate) check() error {
	if v, ok := iu.mutation.Title(); ok {
		if err := index.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	return nil
}

func (iu *IndexUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   index.Table,
			Columns: index.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: index.FieldID,
			},
		},
	}
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.ContentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: index.FieldContentID,
		})
	}
	if value, ok := iu.mutation.AddedContentID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: index.FieldContentID,
		})
	}
	if value, ok := iu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: index.FieldTitle,
		})
	}
	if value, ok := iu.mutation.Desc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: index.FieldDesc,
		})
	}
	if value, ok := iu.mutation.Attr(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: index.FieldAttr,
		})
	}
	if value, ok := iu.mutation.AddedAttr(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: index.FieldAttr,
		})
	}
	if value, ok := iu.mutation.View(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: index.FieldView,
		})
	}
	if value, ok := iu.mutation.AddedView(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: index.FieldView,
		})
	}
	if value, ok := iu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: index.FieldUpdatedAt,
		})
	}
	if value, ok := iu.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: index.FieldState,
		})
	}
	if value, ok := iu.mutation.AddedState(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: index.FieldState,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{index.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// IndexUpdateOne is the builder for updating a single Index entity.
type IndexUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *IndexMutation
}

// SetContentID sets the "content_id" field.
func (iuo *IndexUpdateOne) SetContentID(i int32) *IndexUpdateOne {
	iuo.mutation.ResetContentID()
	iuo.mutation.SetContentID(i)
	return iuo
}

// AddContentID adds i to the "content_id" field.
func (iuo *IndexUpdateOne) AddContentID(i int32) *IndexUpdateOne {
	iuo.mutation.AddContentID(i)
	return iuo
}

// SetTitle sets the "title" field.
func (iuo *IndexUpdateOne) SetTitle(s string) *IndexUpdateOne {
	iuo.mutation.SetTitle(s)
	return iuo
}

// SetDesc sets the "desc" field.
func (iuo *IndexUpdateOne) SetDesc(s string) *IndexUpdateOne {
	iuo.mutation.SetDesc(s)
	return iuo
}

// SetAttr sets the "attr" field.
func (iuo *IndexUpdateOne) SetAttr(i int32) *IndexUpdateOne {
	iuo.mutation.ResetAttr()
	iuo.mutation.SetAttr(i)
	return iuo
}

// SetNillableAttr sets the "attr" field if the given value is not nil.
func (iuo *IndexUpdateOne) SetNillableAttr(i *int32) *IndexUpdateOne {
	if i != nil {
		iuo.SetAttr(*i)
	}
	return iuo
}

// AddAttr adds i to the "attr" field.
func (iuo *IndexUpdateOne) AddAttr(i int32) *IndexUpdateOne {
	iuo.mutation.AddAttr(i)
	return iuo
}

// SetView sets the "view" field.
func (iuo *IndexUpdateOne) SetView(i int32) *IndexUpdateOne {
	iuo.mutation.ResetView()
	iuo.mutation.SetView(i)
	return iuo
}

// SetNillableView sets the "view" field if the given value is not nil.
func (iuo *IndexUpdateOne) SetNillableView(i *int32) *IndexUpdateOne {
	if i != nil {
		iuo.SetView(*i)
	}
	return iuo
}

// AddView adds i to the "view" field.
func (iuo *IndexUpdateOne) AddView(i int32) *IndexUpdateOne {
	iuo.mutation.AddView(i)
	return iuo
}

// SetUpdatedAt sets the "updated_at" field.
func (iuo *IndexUpdateOne) SetUpdatedAt(t time.Time) *IndexUpdateOne {
	iuo.mutation.SetUpdatedAt(t)
	return iuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (iuo *IndexUpdateOne) SetNillableUpdatedAt(t *time.Time) *IndexUpdateOne {
	if t != nil {
		iuo.SetUpdatedAt(*t)
	}
	return iuo
}

// SetState sets the "state" field.
func (iuo *IndexUpdateOne) SetState(i int8) *IndexUpdateOne {
	iuo.mutation.ResetState()
	iuo.mutation.SetState(i)
	return iuo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (iuo *IndexUpdateOne) SetNillableState(i *int8) *IndexUpdateOne {
	if i != nil {
		iuo.SetState(*i)
	}
	return iuo
}

// AddState adds i to the "state" field.
func (iuo *IndexUpdateOne) AddState(i int8) *IndexUpdateOne {
	iuo.mutation.AddState(i)
	return iuo
}

// Mutation returns the IndexMutation object of the builder.
func (iuo *IndexUpdateOne) Mutation() *IndexMutation {
	return iuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *IndexUpdateOne) Select(field string, fields ...string) *IndexUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Index entity.
func (iuo *IndexUpdateOne) Save(ctx context.Context) (*Index, error) {
	var (
		err  error
		node *Index
	)
	if len(iuo.hooks) == 0 {
		if err = iuo.check(); err != nil {
			return nil, err
		}
		node, err = iuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IndexMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = iuo.check(); err != nil {
				return nil, err
			}
			iuo.mutation = mutation
			node, err = iuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iuo.hooks) - 1; i >= 0; i-- {
			if iuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *IndexUpdateOne) SaveX(ctx context.Context) *Index {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *IndexUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *IndexUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *IndexUpdateOne) check() error {
	if v, ok := iuo.mutation.Title(); ok {
		if err := index.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	return nil
}

func (iuo *IndexUpdateOne) sqlSave(ctx context.Context) (_node *Index, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   index.Table,
			Columns: index.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: index.FieldID,
			},
		},
	}
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Index.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, index.FieldID)
		for _, f := range fields {
			if !index.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != index.FieldID {
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
	if value, ok := iuo.mutation.ContentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: index.FieldContentID,
		})
	}
	if value, ok := iuo.mutation.AddedContentID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: index.FieldContentID,
		})
	}
	if value, ok := iuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: index.FieldTitle,
		})
	}
	if value, ok := iuo.mutation.Desc(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: index.FieldDesc,
		})
	}
	if value, ok := iuo.mutation.Attr(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: index.FieldAttr,
		})
	}
	if value, ok := iuo.mutation.AddedAttr(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: index.FieldAttr,
		})
	}
	if value, ok := iuo.mutation.View(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: index.FieldView,
		})
	}
	if value, ok := iuo.mutation.AddedView(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: index.FieldView,
		})
	}
	if value, ok := iuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: index.FieldUpdatedAt,
		})
	}
	if value, ok := iuo.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: index.FieldState,
		})
	}
	if value, ok := iuo.mutation.AddedState(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: index.FieldState,
		})
	}
	_node = &Index{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{index.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
