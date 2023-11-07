// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TodoColumns holds the columns for the "todo" table.
	TodoColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString, Size: 26},
		{Name: "priority", Type: field.TypeEnum, Enums: []string{"HIGH", "MEDIUM", "LOW", "NONE"}, Default: "NONE"},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"IN_PROGRESS", "COMPLETED"}, Default: "IN_PROGRESS"},
		{Name: "time_completed", Type: field.TypeTime, Nullable: true},
		{Name: "todo_parent", Type: field.TypeString, Nullable: true},
	}
	// TodoTable holds the schema information for the "todo" table.
	TodoTable = &schema.Table{
		Name:       "todo",
		Columns:    TodoColumns,
		PrimaryKey: []*schema.Column{TodoColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "todo_todo_parent",
				Columns:    []*schema.Column{TodoColumns[7]},
				RefColumns: []*schema.Column{TodoColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "todo_title_priority_status",
				Unique:  false,
				Columns: []*schema.Column{TodoColumns[3], TodoColumns[4], TodoColumns[5]},
			},
		},
	}
	// UserColumns holds the columns for the "user" table.
	UserColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "username", Type: field.TypeString, Unique: true, Size: 18},
		{Name: "display_name", Type: field.TypeString},
	}
	// UserTable holds the schema information for the "user" table.
	UserTable = &schema.Table{
		Name:       "user",
		Columns:    UserColumns,
		PrimaryKey: []*schema.Column{UserColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_display_name",
				Unique:  false,
				Columns: []*schema.Column{UserColumns[4]},
			},
			{
				Name:    "user_username",
				Unique:  true,
				Columns: []*schema.Column{UserColumns[3]},
			},
		},
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
				RefColumns: []*schema.Column{UserColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_todos_todo_id",
				Columns:    []*schema.Column{UserTodosColumns[1]},
				RefColumns: []*schema.Column{TodoColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TodoTable,
		UserTable,
		UserTodosTable,
	}
)

func init() {
	TodoTable.ForeignKeys[0].RefTable = TodoTable
	TodoTable.Annotation = &entsql.Annotation{
		Table: "todo",
	}
	UserTable.Annotation = &entsql.Annotation{
		Table: "user",
	}
	UserTodosTable.ForeignKeys[0].RefTable = UserTable
	UserTodosTable.ForeignKeys[1].RefTable = TodoTable
}
