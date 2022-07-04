package schema

import "entgo.io/ent"

// Tab holds the schema definition for the Tab entity.
type Tab struct {
	ent.Schema
}

// Fields of the Tab.
func (Tab) Fields() []ent.Field {
	return nil
}

// Edges of the Tab.
func (Tab) Edges() []ent.Edge {
	return nil
}
