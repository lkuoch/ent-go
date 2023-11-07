// Code generated by ent, DO NOT EDIT.

package generated

import (
	"lkuoch/ent-todo/ent/generated/todo"
	"lkuoch/ent-todo/ent/schema/types/pulid"
	"time"
)

// CreateTodoInput represents a mutation input for creating todos.
type CreateTodoInput struct {
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	Title         string
	Priority      *todo.Priority
	Status        *todo.Status
	TimeCompleted *time.Time
	ChildIDs      []pulid.ID
	ParentID      *pulid.ID
	UserIDs       []pulid.ID
}

// Mutate applies the CreateTodoInput on the TodoMutation builder.
func (i *CreateTodoInput) Mutate(m *TodoMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetTitle(i.Title)
	if v := i.Priority; v != nil {
		m.SetPriority(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.TimeCompleted; v != nil {
		m.SetTimeCompleted(*v)
	}
	if v := i.ChildIDs; len(v) > 0 {
		m.AddChildIDs(v...)
	}
	if v := i.ParentID; v != nil {
		m.SetParentID(*v)
	}
	if v := i.UserIDs; len(v) > 0 {
		m.AddUserIDs(v...)
	}
}

// SetInput applies the change-set in the CreateTodoInput on the TodoCreate builder.
func (c *TodoCreate) SetInput(i CreateTodoInput) *TodoCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	Username    string
	DisplayName string
	TodoIDs     []pulid.ID
}

// Mutate applies the CreateUserInput on the UserMutation builder.
func (i *CreateUserInput) Mutate(m *UserMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	m.SetUsername(i.Username)
	m.SetDisplayName(i.DisplayName)
	if v := i.TodoIDs; len(v) > 0 {
		m.AddTodoIDs(v...)
	}
}

// SetInput applies the change-set in the CreateUserInput on the UserCreate builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c.Mutation())
	return c
}
