// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/samb233/easyblog/internal/data/ent/content"
	"github.com/samb233/easyblog/internal/data/ent/index"
	"github.com/samb233/easyblog/internal/data/ent/predicate"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeContent = "Content"
	TypeIndex   = "Index"
)

// ContentMutation represents an operation that mutates the Content nodes in the graph.
type ContentMutation struct {
	config
	op            Op
	typ           string
	id            *int32
	content       *string
	created_at    *time.Time
	updated_at    *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Content, error)
	predicates    []predicate.Content
}

var _ ent.Mutation = (*ContentMutation)(nil)

// contentOption allows management of the mutation configuration using functional options.
type contentOption func(*ContentMutation)

// newContentMutation creates new mutation for the Content entity.
func newContentMutation(c config, op Op, opts ...contentOption) *ContentMutation {
	m := &ContentMutation{
		config:        c,
		op:            op,
		typ:           TypeContent,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withContentID sets the ID field of the mutation.
func withContentID(id int32) contentOption {
	return func(m *ContentMutation) {
		var (
			err   error
			once  sync.Once
			value *Content
		)
		m.oldValue = func(ctx context.Context) (*Content, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Content.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withContent sets the old Content of the mutation.
func withContent(node *Content) contentOption {
	return func(m *ContentMutation) {
		m.oldValue = func(context.Context) (*Content, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ContentMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ContentMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Content entities.
func (m *ContentMutation) SetID(id int32) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ContentMutation) ID() (id int32, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetContent sets the "content" field.
func (m *ContentMutation) SetContent(s string) {
	m.content = &s
}

// Content returns the value of the "content" field in the mutation.
func (m *ContentMutation) Content() (r string, exists bool) {
	v := m.content
	if v == nil {
		return
	}
	return *v, true
}

// OldContent returns the old "content" field's value of the Content entity.
// If the Content object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ContentMutation) OldContent(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldContent is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldContent requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldContent: %w", err)
	}
	return oldValue.Content, nil
}

// ResetContent resets all changes to the "content" field.
func (m *ContentMutation) ResetContent() {
	m.content = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *ContentMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *ContentMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Content entity.
// If the Content object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ContentMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *ContentMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *ContentMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *ContentMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Content entity.
// If the Content object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ContentMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *ContentMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// Where appends a list predicates to the ContentMutation builder.
func (m *ContentMutation) Where(ps ...predicate.Content) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *ContentMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Content).
func (m *ContentMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ContentMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.content != nil {
		fields = append(fields, content.FieldContent)
	}
	if m.created_at != nil {
		fields = append(fields, content.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, content.FieldUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ContentMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case content.FieldContent:
		return m.Content()
	case content.FieldCreatedAt:
		return m.CreatedAt()
	case content.FieldUpdatedAt:
		return m.UpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ContentMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case content.FieldContent:
		return m.OldContent(ctx)
	case content.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case content.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Content field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ContentMutation) SetField(name string, value ent.Value) error {
	switch name {
	case content.FieldContent:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetContent(v)
		return nil
	case content.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case content.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Content field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ContentMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ContentMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ContentMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Content numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ContentMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ContentMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ContentMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Content nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ContentMutation) ResetField(name string) error {
	switch name {
	case content.FieldContent:
		m.ResetContent()
		return nil
	case content.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case content.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Content field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ContentMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ContentMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ContentMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ContentMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ContentMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ContentMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ContentMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Content unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ContentMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Content edge %s", name)
}

// IndexMutation represents an operation that mutates the Index nodes in the graph.
type IndexMutation struct {
	config
	op            Op
	typ           string
	id            *int32
	content_id    *int32
	addcontent_id *int32
	title         *string
	desc          *string
	attr          *int32
	addattr       *int32
	view          *int32
	addview       *int32
	created_at    *time.Time
	updated_at    *time.Time
	state         *int8
	addstate      *int8
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Index, error)
	predicates    []predicate.Index
}

var _ ent.Mutation = (*IndexMutation)(nil)

// indexOption allows management of the mutation configuration using functional options.
type indexOption func(*IndexMutation)

// newIndexMutation creates new mutation for the Index entity.
func newIndexMutation(c config, op Op, opts ...indexOption) *IndexMutation {
	m := &IndexMutation{
		config:        c,
		op:            op,
		typ:           TypeIndex,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withIndexID sets the ID field of the mutation.
func withIndexID(id int32) indexOption {
	return func(m *IndexMutation) {
		var (
			err   error
			once  sync.Once
			value *Index
		)
		m.oldValue = func(ctx context.Context) (*Index, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Index.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withIndex sets the old Index of the mutation.
func withIndex(node *Index) indexOption {
	return func(m *IndexMutation) {
		m.oldValue = func(context.Context) (*Index, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m IndexMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m IndexMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Index entities.
func (m *IndexMutation) SetID(id int32) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *IndexMutation) ID() (id int32, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetContentID sets the "content_id" field.
func (m *IndexMutation) SetContentID(i int32) {
	m.content_id = &i
	m.addcontent_id = nil
}

// ContentID returns the value of the "content_id" field in the mutation.
func (m *IndexMutation) ContentID() (r int32, exists bool) {
	v := m.content_id
	if v == nil {
		return
	}
	return *v, true
}

// OldContentID returns the old "content_id" field's value of the Index entity.
// If the Index object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *IndexMutation) OldContentID(ctx context.Context) (v int32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldContentID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldContentID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldContentID: %w", err)
	}
	return oldValue.ContentID, nil
}

// AddContentID adds i to the "content_id" field.
func (m *IndexMutation) AddContentID(i int32) {
	if m.addcontent_id != nil {
		*m.addcontent_id += i
	} else {
		m.addcontent_id = &i
	}
}

// AddedContentID returns the value that was added to the "content_id" field in this mutation.
func (m *IndexMutation) AddedContentID() (r int32, exists bool) {
	v := m.addcontent_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetContentID resets all changes to the "content_id" field.
func (m *IndexMutation) ResetContentID() {
	m.content_id = nil
	m.addcontent_id = nil
}

// SetTitle sets the "title" field.
func (m *IndexMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the value of the "title" field in the mutation.
func (m *IndexMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "title" field's value of the Index entity.
// If the Index object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *IndexMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle resets all changes to the "title" field.
func (m *IndexMutation) ResetTitle() {
	m.title = nil
}

// SetDesc sets the "desc" field.
func (m *IndexMutation) SetDesc(s string) {
	m.desc = &s
}

// Desc returns the value of the "desc" field in the mutation.
func (m *IndexMutation) Desc() (r string, exists bool) {
	v := m.desc
	if v == nil {
		return
	}
	return *v, true
}

// OldDesc returns the old "desc" field's value of the Index entity.
// If the Index object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *IndexMutation) OldDesc(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldDesc is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldDesc requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDesc: %w", err)
	}
	return oldValue.Desc, nil
}

// ResetDesc resets all changes to the "desc" field.
func (m *IndexMutation) ResetDesc() {
	m.desc = nil
}

// SetAttr sets the "attr" field.
func (m *IndexMutation) SetAttr(i int32) {
	m.attr = &i
	m.addattr = nil
}

// Attr returns the value of the "attr" field in the mutation.
func (m *IndexMutation) Attr() (r int32, exists bool) {
	v := m.attr
	if v == nil {
		return
	}
	return *v, true
}

// OldAttr returns the old "attr" field's value of the Index entity.
// If the Index object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *IndexMutation) OldAttr(ctx context.Context) (v int32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldAttr is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldAttr requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAttr: %w", err)
	}
	return oldValue.Attr, nil
}

// AddAttr adds i to the "attr" field.
func (m *IndexMutation) AddAttr(i int32) {
	if m.addattr != nil {
		*m.addattr += i
	} else {
		m.addattr = &i
	}
}

// AddedAttr returns the value that was added to the "attr" field in this mutation.
func (m *IndexMutation) AddedAttr() (r int32, exists bool) {
	v := m.addattr
	if v == nil {
		return
	}
	return *v, true
}

// ResetAttr resets all changes to the "attr" field.
func (m *IndexMutation) ResetAttr() {
	m.attr = nil
	m.addattr = nil
}

// SetView sets the "view" field.
func (m *IndexMutation) SetView(i int32) {
	m.view = &i
	m.addview = nil
}

// View returns the value of the "view" field in the mutation.
func (m *IndexMutation) View() (r int32, exists bool) {
	v := m.view
	if v == nil {
		return
	}
	return *v, true
}

// OldView returns the old "view" field's value of the Index entity.
// If the Index object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *IndexMutation) OldView(ctx context.Context) (v int32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldView is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldView requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldView: %w", err)
	}
	return oldValue.View, nil
}

// AddView adds i to the "view" field.
func (m *IndexMutation) AddView(i int32) {
	if m.addview != nil {
		*m.addview += i
	} else {
		m.addview = &i
	}
}

// AddedView returns the value that was added to the "view" field in this mutation.
func (m *IndexMutation) AddedView() (r int32, exists bool) {
	v := m.addview
	if v == nil {
		return
	}
	return *v, true
}

// ResetView resets all changes to the "view" field.
func (m *IndexMutation) ResetView() {
	m.view = nil
	m.addview = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *IndexMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *IndexMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Index entity.
// If the Index object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *IndexMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *IndexMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *IndexMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *IndexMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Index entity.
// If the Index object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *IndexMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *IndexMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetState sets the "state" field.
func (m *IndexMutation) SetState(i int8) {
	m.state = &i
	m.addstate = nil
}

// State returns the value of the "state" field in the mutation.
func (m *IndexMutation) State() (r int8, exists bool) {
	v := m.state
	if v == nil {
		return
	}
	return *v, true
}

// OldState returns the old "state" field's value of the Index entity.
// If the Index object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *IndexMutation) OldState(ctx context.Context) (v int8, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldState is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldState requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldState: %w", err)
	}
	return oldValue.State, nil
}

// AddState adds i to the "state" field.
func (m *IndexMutation) AddState(i int8) {
	if m.addstate != nil {
		*m.addstate += i
	} else {
		m.addstate = &i
	}
}

// AddedState returns the value that was added to the "state" field in this mutation.
func (m *IndexMutation) AddedState() (r int8, exists bool) {
	v := m.addstate
	if v == nil {
		return
	}
	return *v, true
}

// ResetState resets all changes to the "state" field.
func (m *IndexMutation) ResetState() {
	m.state = nil
	m.addstate = nil
}

// Where appends a list predicates to the IndexMutation builder.
func (m *IndexMutation) Where(ps ...predicate.Index) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *IndexMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Index).
func (m *IndexMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *IndexMutation) Fields() []string {
	fields := make([]string, 0, 8)
	if m.content_id != nil {
		fields = append(fields, index.FieldContentID)
	}
	if m.title != nil {
		fields = append(fields, index.FieldTitle)
	}
	if m.desc != nil {
		fields = append(fields, index.FieldDesc)
	}
	if m.attr != nil {
		fields = append(fields, index.FieldAttr)
	}
	if m.view != nil {
		fields = append(fields, index.FieldView)
	}
	if m.created_at != nil {
		fields = append(fields, index.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, index.FieldUpdatedAt)
	}
	if m.state != nil {
		fields = append(fields, index.FieldState)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *IndexMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case index.FieldContentID:
		return m.ContentID()
	case index.FieldTitle:
		return m.Title()
	case index.FieldDesc:
		return m.Desc()
	case index.FieldAttr:
		return m.Attr()
	case index.FieldView:
		return m.View()
	case index.FieldCreatedAt:
		return m.CreatedAt()
	case index.FieldUpdatedAt:
		return m.UpdatedAt()
	case index.FieldState:
		return m.State()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *IndexMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case index.FieldContentID:
		return m.OldContentID(ctx)
	case index.FieldTitle:
		return m.OldTitle(ctx)
	case index.FieldDesc:
		return m.OldDesc(ctx)
	case index.FieldAttr:
		return m.OldAttr(ctx)
	case index.FieldView:
		return m.OldView(ctx)
	case index.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case index.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case index.FieldState:
		return m.OldState(ctx)
	}
	return nil, fmt.Errorf("unknown Index field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *IndexMutation) SetField(name string, value ent.Value) error {
	switch name {
	case index.FieldContentID:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetContentID(v)
		return nil
	case index.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case index.FieldDesc:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDesc(v)
		return nil
	case index.FieldAttr:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAttr(v)
		return nil
	case index.FieldView:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetView(v)
		return nil
	case index.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case index.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case index.FieldState:
		v, ok := value.(int8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetState(v)
		return nil
	}
	return fmt.Errorf("unknown Index field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *IndexMutation) AddedFields() []string {
	var fields []string
	if m.addcontent_id != nil {
		fields = append(fields, index.FieldContentID)
	}
	if m.addattr != nil {
		fields = append(fields, index.FieldAttr)
	}
	if m.addview != nil {
		fields = append(fields, index.FieldView)
	}
	if m.addstate != nil {
		fields = append(fields, index.FieldState)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *IndexMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case index.FieldContentID:
		return m.AddedContentID()
	case index.FieldAttr:
		return m.AddedAttr()
	case index.FieldView:
		return m.AddedView()
	case index.FieldState:
		return m.AddedState()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *IndexMutation) AddField(name string, value ent.Value) error {
	switch name {
	case index.FieldContentID:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddContentID(v)
		return nil
	case index.FieldAttr:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAttr(v)
		return nil
	case index.FieldView:
		v, ok := value.(int32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddView(v)
		return nil
	case index.FieldState:
		v, ok := value.(int8)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddState(v)
		return nil
	}
	return fmt.Errorf("unknown Index numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *IndexMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *IndexMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *IndexMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Index nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *IndexMutation) ResetField(name string) error {
	switch name {
	case index.FieldContentID:
		m.ResetContentID()
		return nil
	case index.FieldTitle:
		m.ResetTitle()
		return nil
	case index.FieldDesc:
		m.ResetDesc()
		return nil
	case index.FieldAttr:
		m.ResetAttr()
		return nil
	case index.FieldView:
		m.ResetView()
		return nil
	case index.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case index.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case index.FieldState:
		m.ResetState()
		return nil
	}
	return fmt.Errorf("unknown Index field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *IndexMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *IndexMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *IndexMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *IndexMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *IndexMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *IndexMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *IndexMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Index unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *IndexMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Index edge %s", name)
}
