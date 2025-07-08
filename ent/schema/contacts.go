package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Contacts holds the schema definition for the Contacts entity.
type Contacts struct {
	ent.Schema
}

// Fields of the Contacts.
func (Contacts) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("telefone"),
		field.String("email").Unique(),
		field.String("vat").Unique(),
		field.String("type"),
		field.Time("created_at").Default(time.Now),
	}
}

func (Contacts) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("addresses", Addresses.Type),
		edge.From("files", Files.Type).
			Ref("contacts"),
	}
}
