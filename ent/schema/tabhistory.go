package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TabHistory holds the schema definition for the TabHistory entity.
type TabHistory struct {
	ent.Schema
}

// Fields of the TabHistory.
func (TabHistory) Fields() []ent.Field {
	return []ent.Field{
		field.String("url"),
		field.String("name").Optional(),
		field.String("icon").Optional(),
	}
}

// Edges of the TabHistory.
func (TabHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("history_list", ExtensionClient.Type).Ref("histories").Unique(),
	}
}
