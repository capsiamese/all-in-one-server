package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"notification/internal/entity"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("share_url").Optional(),
		field.JSON("option", entity.GroupOption{}).Optional(),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tabs", Tab.Type),
		edge.From("client", ExtensionClient.Type).Ref("groups").Unique(),
	}
}
