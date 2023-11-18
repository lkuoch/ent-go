// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"lkuoch/ent-todo/ent/generated/remote"
	"lkuoch/ent-todo/ent/schema/types"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Remote is the model entity for the Remote schema.
type Remote struct {
	config `json:"-"`
	// ID of the ent.
	ID types.ID `json:"id,omitempty"`
	// Data holds the value of the "data" field.
	Data         string `json:"data,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Remote) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case remote.FieldData:
			values[i] = new(sql.NullString)
		case remote.FieldID:
			values[i] = new(types.ID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Remote fields.
func (r *Remote) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case remote.FieldID:
			if value, ok := values[i].(*types.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				r.ID = *value
			}
		case remote.FieldData:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field data", values[i])
			} else if value.Valid {
				r.Data = value.String
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Remote.
// This includes values selected through modifiers, order, etc.
func (r *Remote) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// Update returns a builder for updating this Remote.
// Note that you need to call Remote.Unwrap() before calling this method if this Remote
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Remote) Update() *RemoteUpdateOne {
	return NewRemoteClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Remote entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Remote) Unwrap() *Remote {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("generated: Remote is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Remote) String() string {
	var builder strings.Builder
	builder.WriteString("Remote(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("data=")
	builder.WriteString(r.Data)
	builder.WriteByte(')')
	return builder.String()
}

// Remotes is a parsable slice of Remote.
type Remotes []*Remote