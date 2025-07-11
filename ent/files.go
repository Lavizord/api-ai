// Code generated by ent, DO NOT EDIT.

package ent

import (
	"api-ai/ent/files"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Files is the model entity for the Files schema.
type Files struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// FileSource holds the value of the "file_source" field.
	FileSource string `json:"file_source,omitempty"`
	// FileName holds the value of the "file_name" field.
	FileName string `json:"file_name,omitempty"`
	// FileURL holds the value of the "file_url" field.
	FileURL string `json:"file_url,omitempty"`
	// FileData holds the value of the "file_data" field.
	FileData []byte `json:"file_data,omitempty"`
	// PromptUsed holds the value of the "prompt_used" field.
	PromptUsed string `json:"prompt_used,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FilesQuery when eager-loading is set.
	Edges        FilesEdges `json:"edges"`
	selectValues sql.SelectValues
}

// FilesEdges holds the relations/edges for other nodes in the graph.
type FilesEdges struct {
	// Contacts holds the value of the contacts edge.
	Contacts []*Contacts `json:"contacts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ContactsOrErr returns the Contacts value or an error if the edge
// was not loaded in eager-loading.
func (e FilesEdges) ContactsOrErr() ([]*Contacts, error) {
	if e.loadedTypes[0] {
		return e.Contacts, nil
	}
	return nil, &NotLoadedError{edge: "contacts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Files) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case files.FieldFileData:
			values[i] = new([]byte)
		case files.FieldID:
			values[i] = new(sql.NullInt64)
		case files.FieldFileSource, files.FieldFileName, files.FieldFileURL, files.FieldPromptUsed, files.FieldType:
			values[i] = new(sql.NullString)
		case files.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Files fields.
func (f *Files) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case files.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = int(value.Int64)
		case files.FieldFileSource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field file_source", values[i])
			} else if value.Valid {
				f.FileSource = value.String
			}
		case files.FieldFileName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field file_name", values[i])
			} else if value.Valid {
				f.FileName = value.String
			}
		case files.FieldFileURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field file_url", values[i])
			} else if value.Valid {
				f.FileURL = value.String
			}
		case files.FieldFileData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field file_data", values[i])
			} else if value != nil {
				f.FileData = *value
			}
		case files.FieldPromptUsed:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field prompt_used", values[i])
			} else if value.Valid {
				f.PromptUsed = value.String
			}
		case files.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				f.CreatedAt = value.Time
			}
		case files.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				f.Type = value.String
			}
		default:
			f.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Files.
// This includes values selected through modifiers, order, etc.
func (f *Files) Value(name string) (ent.Value, error) {
	return f.selectValues.Get(name)
}

// QueryContacts queries the "contacts" edge of the Files entity.
func (f *Files) QueryContacts() *ContactsQuery {
	return NewFilesClient(f.config).QueryContacts(f)
}

// Update returns a builder for updating this Files.
// Note that you need to call Files.Unwrap() before calling this method if this Files
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Files) Update() *FilesUpdateOne {
	return NewFilesClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the Files entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Files) Unwrap() *Files {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Files is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Files) String() string {
	var builder strings.Builder
	builder.WriteString("Files(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("file_source=")
	builder.WriteString(f.FileSource)
	builder.WriteString(", ")
	builder.WriteString("file_name=")
	builder.WriteString(f.FileName)
	builder.WriteString(", ")
	builder.WriteString("file_url=")
	builder.WriteString(f.FileURL)
	builder.WriteString(", ")
	builder.WriteString("file_data=")
	builder.WriteString(fmt.Sprintf("%v", f.FileData))
	builder.WriteString(", ")
	builder.WriteString("prompt_used=")
	builder.WriteString(f.PromptUsed)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(f.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(f.Type)
	builder.WriteByte(')')
	return builder.String()
}

// FilesSlice is a parsable slice of Files.
type FilesSlice []*Files
