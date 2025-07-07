package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Addresses holds the schema definition for the Addresses entity.
type Addresses struct {
	ent.Schema
}

// Fields of the Addresses.
func (Addresses) Fields() []ent.Field {
	return []ent.Field{
		field.String("street_name"),
		field.String("number"),
		field.String("floor").Optional(),
		field.String("postal_code"),
		field.String("city"),
		field.Bool("invoice_address").Default(false),
		field.Bool("service_address").Default(false),
		field.Time("created_at").Default(time.Now),
	}
}

func (Addresses) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("contact", Contacts.Type).
			Ref("addresses").
			Required().
			Unique(),
	}
}
