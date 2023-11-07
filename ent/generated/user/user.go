// Code generated by ent, DO NOT EDIT.

package user

import (
	"lkuoch/ent-todo/ent/schema/types/pulid"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldDisplayName holds the string denoting the display_name field in the database.
	FieldDisplayName = "display_name"
	// EdgeTodos holds the string denoting the todos edge name in mutations.
	EdgeTodos = "todos"
	// Table holds the table name of the user in the database.
	Table = "user"
	// TodosTable is the table that holds the todos relation/edge. The primary key declared below.
	TodosTable = "user_todos"
	// TodosInverseTable is the table name for the Todo entity.
	// It exists in this package in order to avoid circular dependency with the "todo" package.
	TodosInverseTable = "todo"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldUsername,
	FieldDisplayName,
}

var (
	// TodosPrimaryKey and TodosColumn2 are the table columns denoting the
	// primary key for the todos relation (M2M).
	TodosPrimaryKey = []string{"user_id", "todo_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// DisplayNameValidator is a validator for the "display_name" field. It is called by the builders before save.
	DisplayNameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() pulid.ID
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByDisplayName orders the results by the display_name field.
func ByDisplayName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDisplayName, opts...).ToFunc()
}

// ByTodosCount orders the results by todos count.
func ByTodosCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTodosStep(), opts...)
	}
}

// ByTodos orders the results by todos terms.
func ByTodos(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTodosStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTodosStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TodosInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, TodosTable, TodosPrimaryKey...),
	)
}
