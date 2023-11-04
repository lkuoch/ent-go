// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gql

import (
	"bytes"
	"context"
	"errors"
	"lkuoch/ent-todo/ent/generated"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		schema:     cfg.Schema,
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Schema     *ast.Schema
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Mutation() MutationResolver
	Query() QueryResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	Mutation struct {
		CreateTodo func(childComplexity int, input generated.CreateTodoInput) int
		CreateUser func(childComplexity int, input generated.CreateUserInput) int
	}

	PageInfo struct {
		EndCursor       func(childComplexity int) int
		HasNextPage     func(childComplexity int) int
		HasPreviousPage func(childComplexity int) int
		StartCursor     func(childComplexity int) int
	}

	Query struct {
		Node  func(childComplexity int, id string) int
		Nodes func(childComplexity int, ids []string) int
		Todos func(childComplexity int) int
		Users func(childComplexity int) int
	}

	Todo struct {
		Children      func(childComplexity int) int
		CreatedAt     func(childComplexity int) int
		ID            func(childComplexity int) int
		Parent        func(childComplexity int) int
		Priority      func(childComplexity int) int
		Status        func(childComplexity int) int
		TimeCompleted func(childComplexity int) int
		Title         func(childComplexity int) int
		UpdatedAt     func(childComplexity int) int
		User          func(childComplexity int) int
	}

	User struct {
		CreatedAt   func(childComplexity int) int
		DisplayName func(childComplexity int) int
		ID          func(childComplexity int) int
		Todos       func(childComplexity int) int
		UpdatedAt   func(childComplexity int) int
		Username    func(childComplexity int) int
	}
}

type executableSchema struct {
	schema     *ast.Schema
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	if e.schema != nil {
		return e.schema
	}
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e, 0, 0, nil}
	_ = ec
	switch typeName + "." + field {

	case "Mutation.createTodo":
		if e.complexity.Mutation.CreateTodo == nil {
			break
		}

		args, err := ec.field_Mutation_createTodo_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateTodo(childComplexity, args["input"].(generated.CreateTodoInput)), true

	case "Mutation.createUser":
		if e.complexity.Mutation.CreateUser == nil {
			break
		}

		args, err := ec.field_Mutation_createUser_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateUser(childComplexity, args["input"].(generated.CreateUserInput)), true

	case "PageInfo.endCursor":
		if e.complexity.PageInfo.EndCursor == nil {
			break
		}

		return e.complexity.PageInfo.EndCursor(childComplexity), true

	case "PageInfo.hasNextPage":
		if e.complexity.PageInfo.HasNextPage == nil {
			break
		}

		return e.complexity.PageInfo.HasNextPage(childComplexity), true

	case "PageInfo.hasPreviousPage":
		if e.complexity.PageInfo.HasPreviousPage == nil {
			break
		}

		return e.complexity.PageInfo.HasPreviousPage(childComplexity), true

	case "PageInfo.startCursor":
		if e.complexity.PageInfo.StartCursor == nil {
			break
		}

		return e.complexity.PageInfo.StartCursor(childComplexity), true

	case "Query.node":
		if e.complexity.Query.Node == nil {
			break
		}

		args, err := ec.field_Query_node_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Node(childComplexity, args["id"].(string)), true

	case "Query.nodes":
		if e.complexity.Query.Nodes == nil {
			break
		}

		args, err := ec.field_Query_nodes_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Nodes(childComplexity, args["ids"].([]string)), true

	case "Query.todos":
		if e.complexity.Query.Todos == nil {
			break
		}

		return e.complexity.Query.Todos(childComplexity), true

	case "Query.users":
		if e.complexity.Query.Users == nil {
			break
		}

		return e.complexity.Query.Users(childComplexity), true

	case "Todo.children":
		if e.complexity.Todo.Children == nil {
			break
		}

		return e.complexity.Todo.Children(childComplexity), true

	case "Todo.createdAt":
		if e.complexity.Todo.CreatedAt == nil {
			break
		}

		return e.complexity.Todo.CreatedAt(childComplexity), true

	case "Todo.id":
		if e.complexity.Todo.ID == nil {
			break
		}

		return e.complexity.Todo.ID(childComplexity), true

	case "Todo.parent":
		if e.complexity.Todo.Parent == nil {
			break
		}

		return e.complexity.Todo.Parent(childComplexity), true

	case "Todo.priority":
		if e.complexity.Todo.Priority == nil {
			break
		}

		return e.complexity.Todo.Priority(childComplexity), true

	case "Todo.status":
		if e.complexity.Todo.Status == nil {
			break
		}

		return e.complexity.Todo.Status(childComplexity), true

	case "Todo.timeCompleted":
		if e.complexity.Todo.TimeCompleted == nil {
			break
		}

		return e.complexity.Todo.TimeCompleted(childComplexity), true

	case "Todo.title":
		if e.complexity.Todo.Title == nil {
			break
		}

		return e.complexity.Todo.Title(childComplexity), true

	case "Todo.updatedAt":
		if e.complexity.Todo.UpdatedAt == nil {
			break
		}

		return e.complexity.Todo.UpdatedAt(childComplexity), true

	case "Todo.user":
		if e.complexity.Todo.User == nil {
			break
		}

		return e.complexity.Todo.User(childComplexity), true

	case "User.createdAt":
		if e.complexity.User.CreatedAt == nil {
			break
		}

		return e.complexity.User.CreatedAt(childComplexity), true

	case "User.displayName":
		if e.complexity.User.DisplayName == nil {
			break
		}

		return e.complexity.User.DisplayName(childComplexity), true

	case "User.id":
		if e.complexity.User.ID == nil {
			break
		}

		return e.complexity.User.ID(childComplexity), true

	case "User.todos":
		if e.complexity.User.Todos == nil {
			break
		}

		return e.complexity.User.Todos(childComplexity), true

	case "User.updatedAt":
		if e.complexity.User.UpdatedAt == nil {
			break
		}

		return e.complexity.User.UpdatedAt(childComplexity), true

	case "User.username":
		if e.complexity.User.Username == nil {
			break
		}

		return e.complexity.User.Username(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e, 0, 0, make(chan graphql.DeferredResult)}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputCreateTodoInput,
		ec.unmarshalInputCreateUserInput,
	)
	first := true

	switch rc.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			var response graphql.Response
			var data graphql.Marshaler
			if first {
				first = false
				ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
				data = ec._Query(ctx, rc.Operation.SelectionSet)
			} else {
				if atomic.LoadInt32(&ec.pendingDeferred) > 0 {
					result := <-ec.deferredResults
					atomic.AddInt32(&ec.pendingDeferred, -1)
					data = result.Result
					response.Path = result.Path
					response.Label = result.Label
					response.Errors = result.Errors
				} else {
					return nil
				}
			}
			var buf bytes.Buffer
			data.MarshalGQL(&buf)
			response.Data = buf.Bytes()
			if atomic.LoadInt32(&ec.deferred) > 0 {
				hasNext := atomic.LoadInt32(&ec.pendingDeferred) > 0
				response.HasNext = &hasNext
			}

			return &response
		}
	case ast.Mutation:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Mutation(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
	deferred        int32
	pendingDeferred int32
	deferredResults chan graphql.DeferredResult
}

func (ec *executionContext) processDeferredGroup(dg graphql.DeferredGroup) {
	atomic.AddInt32(&ec.pendingDeferred, 1)
	go func() {
		ctx := graphql.WithFreshResponseContext(dg.Context)
		dg.FieldSet.Dispatch(ctx)
		ds := graphql.DeferredResult{
			Path:   dg.Path,
			Label:  dg.Label,
			Result: dg.FieldSet,
			Errors: graphql.GetErrors(ctx),
		}
		// null fields should bubble up
		if dg.FieldSet.Invalids > 0 {
			ds.Result = graphql.Null
		}
		ec.deferredResults <- ds
	}()
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(ec.Schema()), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(ec.Schema(), ec.Schema().Types[name]), nil
}

var sources = []*ast.Source{
	{Name: "../../../ent-schema.graphql", Input: `directive @goField(forceResolver: Boolean, name: String) on FIELD_DEFINITION | INPUT_FIELD_DEFINITION
directive @goModel(model: String, models: [String!]) on OBJECT | INPUT_OBJECT | SCALAR | ENUM | INTERFACE | UNION
"""
CreateTodoInput is used for create Todo object.
Input was generated by ent.
"""
input CreateTodoInput {
  createdAt: Time
  updatedAt: Time
  title: String!
  priority: TodoPriority
  status: TodoStatus
  timeCompleted: Time
  childIDs: [ID!]
  parentID: ID
  userIDs: [ID!]
}
"""
CreateUserInput is used for create User object.
Input was generated by ent.
"""
input CreateUserInput {
  createdAt: Time
  updatedAt: Time
  username: String!
  displayName: String!
  todoIDs: [ID!]
}
"""
Define a Relay Cursor type:
https://relay.dev/graphql/connections.htm#sec-Cursor
"""
scalar Cursor
"""
An object with an ID.
Follows the [Relay Global Object Identification Specification](https://relay.dev/graphql/objectidentification.htm)
"""
interface Node @goModel(model: "lkuoch/ent-todo/ent/generated.Noder") {
  """The id of the object."""
  id: ID!
}
"""Possible directions in which to order a list of items when provided an ` + "`" + `orderBy` + "`" + ` argument."""
enum OrderDirection {
  """Specifies an ascending order for a given ` + "`" + `orderBy` + "`" + ` argument."""
  ASC
  """Specifies a descending order for a given ` + "`" + `orderBy` + "`" + ` argument."""
  DESC
}
"""
Information about pagination in a connection.
https://relay.dev/graphql/connections.htm#sec-undefined.PageInfo
"""
type PageInfo {
  """When paginating forwards, are there more items?"""
  hasNextPage: Boolean!
  """When paginating backwards, are there more items?"""
  hasPreviousPage: Boolean!
  """When paginating backwards, the cursor to continue."""
  startCursor: Cursor
  """When paginating forwards, the cursor to continue."""
  endCursor: Cursor
}
type Query {
  """Fetches an object given its ID."""
  node(
    """ID of the object."""
    id: ID!
  ): Node
  """Lookup nodes by a list of IDs."""
  nodes(
    """The list of node IDs."""
    ids: [ID!]!
  ): [Node]!
  todos: [Todo!]!
  users: [User!]!
}
"""The builtin Time type"""
scalar Time
type Todo implements Node {
  id: ID!
  createdAt: Time!
  updatedAt: Time!
  title: String!
  priority: TodoPriority!
  status: TodoStatus!
  timeCompleted: Time
  children: [Todo!]
  parent: Todo
  user: [User!]
}
"""TodoPriority is enum for the field priority"""
enum TodoPriority @goModel(model: "lkuoch/ent-todo/ent/generated/todo.Priority") {
  HIGH
  MEDIUM
  LOW
  NONE
}
"""TodoStatus is enum for the field status"""
enum TodoStatus @goModel(model: "lkuoch/ent-todo/ent/generated/todo.Status") {
  IN_PROGRESS
  COMPLETED
}
type User implements Node {
  id: ID!
  createdAt: Time!
  updatedAt: Time!
  username: String!
  displayName: String!
  todos: [Todo!]
}
`, BuiltIn: false},
	{Name: "../../../typedef.graphql", Input: `# All custom graphql definitions go here.
# Types from ` + "`" + `ent-schema.graphql` + "`" + ` are referenced here

type Mutation {
  createTodo(input: CreateTodoInput!): Todo
  createUser(input: CreateUserInput!): User
}
`, BuiltIn: false},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
