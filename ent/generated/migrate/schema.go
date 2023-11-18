// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TaskColumns holds the columns for the "task" table.
	TaskColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "title", Type: field.TypeString},
		{Name: "item_status", Type: field.TypeEnum, Enums: []string{"IN_PROGRESS", "COMPLETED"}, Default: "IN_PROGRESS"},
		{Name: "todo_tasks", Type: field.TypeString},
	}
	// TaskTable holds the schema information for the "task" table.
	TaskTable = &schema.Table{
		Name:       "task",
		Columns:    TaskColumns,
		PrimaryKey: []*schema.Column{TaskColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "task_todo_tasks",
				Columns:    []*schema.Column{TaskColumns[3]},
				RefColumns: []*schema.Column{TodoColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "task_id",
				Unique:  true,
				Columns: []*schema.Column{TaskColumns[0]},
			},
			{
				Name:    "task_title_item_status",
				Unique:  false,
				Columns: []*schema.Column{TaskColumns[1], TaskColumns[2]},
			},
		},
	}
	// TodoColumns holds the columns for the "todo" table.
	TodoColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString, Size: 26},
		{Name: "body", Type: field.TypeString},
		{Name: "item_priority", Type: field.TypeEnum, Enums: []string{"HIGH", "MEDIUM", "LOW", "NONE"}, Default: "NONE"},
		{Name: "item_status", Type: field.TypeEnum, Enums: []string{"IN_PROGRESS", "COMPLETED"}, Default: "IN_PROGRESS"},
		{Name: "time_completed", Type: field.TypeTime, Nullable: true},
		{Name: "user_todos", Type: field.TypeString},
	}
	// TodoTable holds the schema information for the "todo" table.
	TodoTable = &schema.Table{
		Name:       "todo",
		Columns:    TodoColumns,
		PrimaryKey: []*schema.Column{TodoColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "todo_user_todos",
				Columns:    []*schema.Column{TodoColumns[8]},
				RefColumns: []*schema.Column{UserColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "todo_id",
				Unique:  true,
				Columns: []*schema.Column{TodoColumns[0]},
			},
			{
				Name:    "todo_title_item_priority_item_status",
				Unique:  false,
				Columns: []*schema.Column{TodoColumns[3], TodoColumns[5], TodoColumns[6]},
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
				Name:    "user_id",
				Unique:  true,
				Columns: []*schema.Column{UserColumns[0]},
			},
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
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TaskTable,
		TodoTable,
		UserTable,
	}
)

func init() {
	TaskTable.ForeignKeys[0].RefTable = TodoTable
	TaskTable.Annotation = &entsql.Annotation{
		Table: "task",
	}
	TodoTable.ForeignKeys[0].RefTable = UserTable
	TodoTable.Annotation = &entsql.Annotation{
		Table: "todo",
	}
	UserTable.Annotation = &entsql.Annotation{
		Table: "user",
	}
}
