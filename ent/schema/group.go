package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	uuid "github.com/satori/go.uuid"
	"notification/internal/entity"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uid", uuid.NewV4()),
		field.String("name"),
		field.Time("created_at"),
		field.String("share_url").Optional(),
		field.JSON("option", entity.GroupOption{}).Optional(),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tabs", Tab.Type),
		edge.From("client", ExtensionClient.Type).
			Unique().
			Ref("groups").
			StructTag(`swaggerignore:"true"`),
	}
}

func (Group) Annotations() []schema.Annotation {
	return nil
	/*
		return []schema.Annotation{
			edge.Annotation{StructTag: `swaggerignore:"true"`},
		}
	*/
}
