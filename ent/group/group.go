// Code generated by entc, DO NOT EDIT.

package group

const (
	// Label holds the string label denoting the group type in the database.
	Label = "group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUID holds the string denoting the uid field in the database.
	FieldUID = "uid"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldShareURL holds the string denoting the share_url field in the database.
	FieldShareURL = "share_url"
	// FieldOption holds the string denoting the option field in the database.
	FieldOption = "option"
	// FieldSeq holds the string denoting the seq field in the database.
	FieldSeq = "seq"
	// EdgeTabs holds the string denoting the tabs edge name in mutations.
	EdgeTabs = "tabs"
	// EdgeClient holds the string denoting the client edge name in mutations.
	EdgeClient = "client"
	// Table holds the table name of the group in the database.
	Table = "groups"
	// TabsTable is the table that holds the tabs relation/edge.
	TabsTable = "tabs"
	// TabsInverseTable is the table name for the Tab entity.
	// It exists in this package in order to avoid circular dependency with the "tab" package.
	TabsInverseTable = "tabs"
	// TabsColumn is the table column denoting the tabs relation/edge.
	TabsColumn = "group_tabs"
	// ClientTable is the table that holds the client relation/edge.
	ClientTable = "groups"
	// ClientInverseTable is the table name for the ExtensionClient entity.
	// It exists in this package in order to avoid circular dependency with the "extensionclient" package.
	ClientInverseTable = "extension_clients"
	// ClientColumn is the table column denoting the client relation/edge.
	ClientColumn = "extension_client_groups"
)

// Columns holds all SQL columns for group fields.
var Columns = []string{
	FieldID,
	FieldUID,
	FieldName,
	FieldCreatedAt,
	FieldShareURL,
	FieldOption,
	FieldSeq,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "groups"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"extension_client_groups",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
