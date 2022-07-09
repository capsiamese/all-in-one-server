// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ExtensionClientsColumns holds the columns for the "extension_clients" table.
	ExtensionClientsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "extension_id", Type: field.TypeString},
		{Name: "client_uid", Type: field.TypeUUID, Unique: true},
		{Name: "last_access_time", Type: field.TypeTime},
	}
	// ExtensionClientsTable holds the schema information for the "extension_clients" table.
	ExtensionClientsTable = &schema.Table{
		Name:       "extension_clients",
		Columns:    ExtensionClientsColumns,
		PrimaryKey: []*schema.Column{ExtensionClientsColumns[0]},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uid", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "share_url", Type: field.TypeString, Nullable: true},
		{Name: "option", Type: field.TypeJSON, Nullable: true},
		{Name: "extension_client_groups", Type: field.TypeInt, Nullable: true},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "groups_extension_clients_groups",
				Columns:    []*schema.Column{GroupsColumns[6]},
				RefColumns: []*schema.Column{ExtensionClientsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TabsColumns holds the columns for the "tabs" table.
	TabsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "url", Type: field.TypeString},
		{Name: "seq", Type: field.TypeInt32},
		{Name: "favicon", Type: field.TypeString, Nullable: true},
		{Name: "group_tabs", Type: field.TypeInt, Nullable: true},
	}
	// TabsTable holds the schema information for the "tabs" table.
	TabsTable = &schema.Table{
		Name:       "tabs",
		Columns:    TabsColumns,
		PrimaryKey: []*schema.Column{TabsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tabs_groups_tabs",
				Columns:    []*schema.Column{TabsColumns[5]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ExtensionClientsTable,
		GroupsTable,
		TabsTable,
	}
)

func init() {
	GroupsTable.ForeignKeys[0].RefTable = ExtensionClientsTable
	TabsTable.ForeignKeys[0].RefTable = GroupsTable
}
