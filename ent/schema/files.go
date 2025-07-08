package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Files holds the schema definition for the Files entity.
type Files struct {
	ent.Schema
}

// Fields of the Files.
func (Files) Fields() []ent.Field {
	return []ent.Field{
		field.String("file_source"),                // source of the file
		field.String("file_name"),                  // file name
		field.String("file_url").Optional(),        // URL to access the file
		field.Bytes("file_data").Optional(),        // Also allows the file to be stored.
		field.String("prompt_used"),                // promp used for the chatpdf
		field.Time("created_at").Default(time.Now), // creation timestamp
		field.String("type"),                       // file type (e.g., pdf, jpg)
	}
}

// Files schema
func (Files) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("contacts", Contacts.Type),
	}
}
