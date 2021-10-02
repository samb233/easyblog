package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Content holds the schema definition for the Content entity.
type Content struct {
	ent.Schema
}

func (Content) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "blog_content"},
	}
}

// Fields of the Content.
func (Content) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("id"),
		field.String("content").
			SchemaType(map[string]string{
				dialect.MySQL: "text",
			}),
		field.Time("created_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("updated_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
	}
}

// Edges of the Content.
func (Content) Edges() []ent.Edge {
	return nil
}
