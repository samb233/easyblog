package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Index holds the schema definition for the Index entity.
type Index struct {
	ent.Schema
}

func (Index) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "blog_index"},
	}
}

// Fields of the Index.
func (Index) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id"),
		field.Int32("content_id").
			Unique(),
		field.String("title").
			MaxLen(100),
		field.String("desc"),
		field.Int32("attr").
			Default(0),
		field.Int32("view").
			Default(0),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
		field.Time("updated_at").
			Default(time.Now).
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
		field.Int8("state").
			Default(1),
	}
}

// Indexes of the Index.
func (Index) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("content_id").
			Unique(),
	}
}

// Edges of the Index.
func (Index) Edges() []ent.Edge {
	return nil
}
