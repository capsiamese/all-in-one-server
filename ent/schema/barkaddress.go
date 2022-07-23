package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BarkAddress holds the schema definition for the BarkAddress entity.
type BarkAddress struct {
	ent.Schema
}

// Fields of the BarkAddress.
func (BarkAddress) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("target").NotEmpty(),
		field.Int64("index"),
	}
}

// Edges of the BarkAddress.
func (BarkAddress) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("client", ExtensionClient.Type).
			Ref("addresses").
			StructTag(`swaggerignore:"true"`),
	}
}
