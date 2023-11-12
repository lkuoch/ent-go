// Code generated by ent, DO NOT EDIT.

package task

import (
	"fmt"
	"lkuoch/ent-todo/ent/schema/types"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/99designs/gqlgen/graphql"
)

const (
	// Label holds the string label denoting the task type in the database.
	Label = "task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldItemStatus holds the string denoting the item_status field in the database.
	FieldItemStatus = "item_status"
	// EdgeTodo holds the string denoting the todo edge name in mutations.
	EdgeTodo = "todo"
	// Table holds the table name of the task in the database.
	Table = "task"
	// TodoTable is the table that holds the todo relation/edge.
	TodoTable = "task"
	// TodoInverseTable is the table name for the Todo entity.
	// It exists in this package in order to avoid circular dependency with the "todo" package.
	TodoInverseTable = "todo"
	// TodoColumn is the table column denoting the todo relation/edge.
	TodoColumn = "todo_tasks"
)

// Columns holds all SQL columns for task fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldItemStatus,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "task"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"todo_tasks",
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

var (
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() types.ID
)

const DefaultItemStatus types.ItemStatus = "IN_PROGRESS"

// ItemStatusValidator is a validator for the "item_status" field enum values. It is called by the builders before save.
func ItemStatusValidator(is types.ItemStatus) error {
	switch is {
	case "IN_PROGRESS", "COMPLETED":
		return nil
	default:
		return fmt.Errorf("task: invalid enum value for item_status field: %q", is)
	}
}

// OrderOption defines the ordering options for the Task queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByItemStatus orders the results by the item_status field.
func ByItemStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldItemStatus, opts...).ToFunc()
}

// ByTodoField orders the results by todo field.
func ByTodoField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTodoStep(), sql.OrderByField(field, opts...))
	}
}
func newTodoStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TodoInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, TodoTable, TodoColumn),
	)
}

var (
	// types.ItemStatus must implement graphql.Marshaler.
	_ graphql.Marshaler = (*types.ItemStatus)(nil)
	// types.ItemStatus must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*types.ItemStatus)(nil)
)
