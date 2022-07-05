package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ExtensionClient holds the schema definition for the ExtensionClient entity.
type ExtensionClient struct {
	ent.Schema
}

// Fields of the ExtensionClient.
func (ExtensionClient) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("extension_id").Unique(),
		field.Time("last_access_time"),
	}
}

// Edges of the ExtensionClient.
func (ExtensionClient) Edges() []ent.Edge {
	return nil
}
