// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"notification/ent/extensionclient"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"

	uuid "github.com/satori/go.uuid"
)

// ExtensionClient is the model entity for the ExtensionClient schema.
type ExtensionClient struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// ExtensionID holds the value of the "extension_id" field.
	ExtensionID string `json:"extension_id,omitempty"`
	// ClientUID holds the value of the "client_uid" field.
	ClientUID uuid.UUID `json:"client_uid,omitempty"`
	// LastAccessTime holds the value of the "last_access_time" field.
	LastAccessTime time.Time `json:"last_access_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ExtensionClientQuery when eager-loading is set.
	Edges ExtensionClientEdges `json:"edges"`
}

// ExtensionClientEdges holds the relations/edges for other nodes in the graph.
type ExtensionClientEdges struct {
	// Histories holds the value of the histories edge.
	Histories []*TabHistory `json:"histories,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// HistoriesOrErr returns the Histories value or an error if the edge
// was not loaded in eager-loading.
func (e ExtensionClientEdges) HistoriesOrErr() ([]*TabHistory, error) {
	if e.loadedTypes[0] {
		return e.Histories, nil
	}
	return nil, &NotLoadedError{edge: "histories"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ExtensionClient) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case extensionclient.FieldID:
			values[i] = new(sql.NullInt64)
		case extensionclient.FieldName, extensionclient.FieldExtensionID:
			values[i] = new(sql.NullString)
		case extensionclient.FieldLastAccessTime:
			values[i] = new(sql.NullTime)
		case extensionclient.FieldClientUID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ExtensionClient", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ExtensionClient fields.
func (ec *ExtensionClient) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case extensionclient.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ec.ID = int(value.Int64)
		case extensionclient.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ec.Name = value.String
			}
		case extensionclient.FieldExtensionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field extension_id", values[i])
			} else if value.Valid {
				ec.ExtensionID = value.String
			}
		case extensionclient.FieldClientUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field client_uid", values[i])
			} else if value != nil {
				ec.ClientUID = *value
			}
		case extensionclient.FieldLastAccessTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_access_time", values[i])
			} else if value.Valid {
				ec.LastAccessTime = value.Time
			}
		}
	}
	return nil
}

// QueryHistories queries the "histories" edge of the ExtensionClient entity.
func (ec *ExtensionClient) QueryHistories() *TabHistoryQuery {
	return (&ExtensionClientClient{config: ec.config}).QueryHistories(ec)
}

// Update returns a builder for updating this ExtensionClient.
// Note that you need to call ExtensionClient.Unwrap() before calling this method if this ExtensionClient
// was returned from a transaction, and the transaction was committed or rolled back.
func (ec *ExtensionClient) Update() *ExtensionClientUpdateOne {
	return (&ExtensionClientClient{config: ec.config}).UpdateOne(ec)
}

// Unwrap unwraps the ExtensionClient entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ec *ExtensionClient) Unwrap() *ExtensionClient {
	tx, ok := ec.config.driver.(*txDriver)
	if !ok {
		panic("ent: ExtensionClient is not a transactional entity")
	}
	ec.config.driver = tx.drv
	return ec
}

// String implements the fmt.Stringer.
func (ec *ExtensionClient) String() string {
	var builder strings.Builder
	builder.WriteString("ExtensionClient(")
	builder.WriteString(fmt.Sprintf("id=%v", ec.ID))
	builder.WriteString(", name=")
	builder.WriteString(ec.Name)
	builder.WriteString(", extension_id=")
	builder.WriteString(ec.ExtensionID)
	builder.WriteString(", client_uid=")
	builder.WriteString(fmt.Sprintf("%v", ec.ClientUID))
	builder.WriteString(", last_access_time=")
	builder.WriteString(ec.LastAccessTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ExtensionClients is a parsable slice of ExtensionClient.
type ExtensionClients []*ExtensionClient

func (ec ExtensionClients) config(cfg config) {
	for _i := range ec {
		ec[_i].config = cfg
	}
}
