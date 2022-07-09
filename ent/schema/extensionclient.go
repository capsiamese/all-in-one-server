package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	uuid "github.com/satori/go.uuid"
)

// ExtensionClient holds the schema definition for the ExtensionClient entity.
type ExtensionClient struct {
	ent.Schema
}

// Fields of the ExtensionClient.
func (ExtensionClient) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("extension_id").NotEmpty(),
		field.UUID("client_uid", uuid.NewV4()).Unique(),
		field.Time("last_access_time"),
	}
}

// Edges of the ExtensionClient.
func (ExtensionClient) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("histories", TabHistory.Type),
	}
}
