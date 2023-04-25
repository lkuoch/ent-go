// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TodosColumns holds the columns for the "todos" table.
	TodosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString, Size: 26},
		{Name: "priority", Type: field.TypeEnum, Enums: []string{"HIGH", "MEDIUM", "LOW", "NONE"}, Default: "NONE"},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"IN_PROGRESS", "COMPLETED"}, Default: "IN_PROGRESS"},
		{Name: "time_completed", Type: field.TypeTime, Nullable: true},
		{Name: "todo_parent", Type: field.TypeString, Nullable: true},
	}
	// TodosTable holds the schema information for the "todos" table.
	TodosTable = &schema.Table{
		Name:       "todos",
		Columns:    TodosColumns,
		PrimaryKey: []*schema.Column{TodosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "todos_todos_parent",
				Columns:    []*schema.Column{TodosColumns[7]},
				RefColumns: []*schema.Column{TodosColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "username", Type: field.TypeString, Unique: true, Size: 18},
		{Name: "display_name", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserTodosColumns holds the columns for the "user_todos" table.
	UserTodosColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeString},
		{Name: "todo_id", Type: field.TypeString},
	}
	// UserTodosTable holds the schema information for the "user_todos" table.
	UserTodosTable = &schema.Table{
		Name:       "user_todos",
		Columns:    UserTodosColumns,
		PrimaryKey: []*schema.Column{UserTodosColumns[0], UserTodosColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_todos_user_id",
				Columns:    []*schema.Column{UserTodosColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_todos_todo_id",
				Columns:    []*schema.Column{UserTodosColumns[1]},
				RefColumns: []*schema.Column{TodosColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TodosTable,
		UsersTable,
		UserTodosTable,
	}
)

func init() {
	TodosTable.ForeignKeys[0].RefTable = TodosTable
	UserTodosTable.ForeignKeys[0].RefTable = UsersTable
	UserTodosTable.ForeignKeys[1].RefTable = TodosTable
}
