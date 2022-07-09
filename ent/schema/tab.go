package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Tab holds the schema definition for the Tab entity.
type Tab struct {
	ent.Schema
}

// Fields of the Tab.
func (Tab) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("url"),
		field.Int32("seq").Comment("sequence"),
		field.String("favicon").Optional(),
	}
}

// Edges of the Tab.
func (Tab) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group", Group.Type).Ref("tabs").Unique(),
	}
}

func (Tab) Annotations() []schema.Annotation {
	return []schema.Annotation{
		edge.Annotation{StructTag: `swaggerignore:"true"`},
	}
}
