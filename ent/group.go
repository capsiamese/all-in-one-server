// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"notification/ent/extensionclient"
	"notification/ent/group"
	"notification/internal/entity"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"

	uuid "github.com/satori/go.uuid"
)

// Group is the model entity for the Group schema.
type Group struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UID holds the value of the "uid" field.
	UID uuid.UUID `json:"uid,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// ShareURL holds the value of the "share_url" field.
	ShareURL string `json:"share_url,omitempty"`
	// Option holds the value of the "option" field.
	Option entity.GroupOption `json:"option,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupQuery when eager-loading is set.
	Edges                   GroupEdges `json:"edges" swaggerignore:"true"`
	extension_client_groups *int
}

// GroupEdges holds the relations/edges for other nodes in the graph.
type GroupEdges struct {
	// Tabs holds the value of the tabs edge.
	Tabs []*Tab `json:"tabs,omitempty"`
	// Client holds the value of the client edge.
	Client *ExtensionClient `json:"client,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TabsOrErr returns the Tabs value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) TabsOrErr() ([]*Tab, error) {
	if e.loadedTypes[0] {
		return e.Tabs, nil
	}
	return nil, &NotLoadedError{edge: "tabs"}
}

// ClientOrErr returns the Client value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupEdges) ClientOrErr() (*ExtensionClient, error) {
	if e.loadedTypes[1] {
		if e.Client == nil {
			// The edge client was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: extensionclient.Label}
		}
		return e.Client, nil
	}
	return nil, &NotLoadedError{edge: "client"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Group) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case group.FieldOption:
			values[i] = new([]byte)
		case group.FieldID:
			values[i] = new(sql.NullInt64)
		case group.FieldName, group.FieldShareURL:
			values[i] = new(sql.NullString)
		case group.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case group.FieldUID:
			values[i] = new(uuid.UUID)
		case group.ForeignKeys[0]: // extension_client_groups
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Group", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Group fields.
func (gr *Group) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case group.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gr.ID = int(value.Int64)
		case group.FieldUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field uid", values[i])
			} else if value != nil {
				gr.UID = *value
			}
		case group.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				gr.Name = value.String
			}
		case group.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gr.CreatedAt = value.Time
			}
		case group.FieldShareURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field share_url", values[i])
			} else if value.Valid {
				gr.ShareURL = value.String
			}
		case group.FieldOption:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field option", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &gr.Option); err != nil {
					return fmt.Errorf("unmarshal field option: %w", err)
				}
			}
		case group.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field extension_client_groups", value)
			} else if value.Valid {
				gr.extension_client_groups = new(int)
				*gr.extension_client_groups = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryTabs queries the "tabs" edge of the Group entity.
func (gr *Group) QueryTabs() *TabQuery {
	return (&GroupClient{config: gr.config}).QueryTabs(gr)
}

// QueryClient queries the "client" edge of the Group entity.
func (gr *Group) QueryClient() *ExtensionClientQuery {
	return (&GroupClient{config: gr.config}).QueryClient(gr)
}

// Update returns a builder for updating this Group.
// Note that you need to call Group.Unwrap() before calling this method if this Group
// was returned from a transaction, and the transaction was committed or rolled back.
func (gr *Group) Update() *GroupUpdateOne {
	return (&GroupClient{config: gr.config}).UpdateOne(gr)
}

// Unwrap unwraps the Group entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gr *Group) Unwrap() *Group {
	tx, ok := gr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Group is not a transactional entity")
	}
	gr.config.driver = tx.drv
	return gr
}

// String implements the fmt.Stringer.
func (gr *Group) String() string {
	var builder strings.Builder
	builder.WriteString("Group(")
	builder.WriteString(fmt.Sprintf("id=%v", gr.ID))
	builder.WriteString(", uid=")
	builder.WriteString(fmt.Sprintf("%v", gr.UID))
	builder.WriteString(", name=")
	builder.WriteString(gr.Name)
	builder.WriteString(", created_at=")
	builder.WriteString(gr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", share_url=")
	builder.WriteString(gr.ShareURL)
	builder.WriteString(", option=")
	builder.WriteString(fmt.Sprintf("%v", gr.Option))
	builder.WriteByte(')')
	return builder.String()
}

// Groups is a parsable slice of Group.
type Groups []*Group

func (gr Groups) config(cfg config) {
	for _i := range gr {
		gr[_i].config = cfg
	}
}
