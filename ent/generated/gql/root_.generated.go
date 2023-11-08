// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gql

import (
	"bytes"
	"context"
	"errors"
	"lkuoch/ent-todo/ent/generated"
	"lkuoch/ent-todo/ent/schema/types/pulid"
	"sync/atomic"

	"entgo.io/contrib/entgql"
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
		Node  func(childComplexity int, id pulid.ID) int
		Nodes func(childComplexity int, ids []pulid.ID) int
		Todos func(childComplexity int, after *entgql.Cursor[pulid.ID], first *int, before *entgql.Cursor[pulid.ID], last *int, orderBy []*generated.TodoOrder, where *generated.TodoWhereInput) int
		Users func(childComplexity int, after *entgql.Cursor[pulid.ID], first *int, before *entgql.Cursor[pulid.ID], last *int, orderBy []*generated.UserOrder, where *generated.UserWhereInput) int
	}

	Todo struct {
		CreatedAt     func(childComplexity int) int
		ID            func(childComplexity int) int
		Priority      func(childComplexity int) int
		Status        func(childComplexity int) int
		TimeCompleted func(childComplexity int) int
		Title         func(childComplexity int) int
		UpdatedAt     func(childComplexity int) int
		User          func(childComplexity int) int
	}

	TodoConnection struct {
		Edges      func(childComplexity int) int
		PageInfo   func(childComplexity int) int
		TotalCount func(childComplexity int) int
	}

	TodoEdge struct {
		Cursor func(childComplexity int) int
		Node   func(childComplexity int) int
	}

	User struct {
		CreatedAt   func(childComplexity int) int
		DisplayName func(childComplexity int) int
		ID          func(childComplexity int) int
		Todos       func(childComplexity int) int
		UpdatedAt   func(childComplexity int) int
		Username    func(childComplexity int) int
	}

	UserConnection struct {
		Edges      func(childComplexity int) int
		PageInfo   func(childComplexity int) int
		TotalCount func(childComplexity int) int
	}

	UserEdge struct {
		Cursor func(childComplexity int) int
		Node   func(childComplexity int) int
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

		return e.complexity.Query.Node(childComplexity, args["id"].(pulid.ID)), true

	case "Query.nodes":
		if e.complexity.Query.Nodes == nil {
			break
		}

		args, err := ec.field_Query_nodes_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Nodes(childComplexity, args["ids"].([]pulid.ID)), true

	case "Query.todos":
		if e.complexity.Query.Todos == nil {
			break
		}

		args, err := ec.field_Query_todos_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Todos(childComplexity, args["after"].(*entgql.Cursor[pulid.ID]), args["first"].(*int), args["before"].(*entgql.Cursor[pulid.ID]), args["last"].(*int), args["orderBy"].([]*generated.TodoOrder), args["where"].(*generated.TodoWhereInput)), true

	case "Query.users":
		if e.complexity.Query.Users == nil {
			break
		}

		args, err := ec.field_Query_users_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Users(childComplexity, args["after"].(*entgql.Cursor[pulid.ID]), args["first"].(*int), args["before"].(*entgql.Cursor[pulid.ID]), args["last"].(*int), args["orderBy"].([]*generated.UserOrder), args["where"].(*generated.UserWhereInput)), true

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

	case "TodoConnection.edges":
		if e.complexity.TodoConnection.Edges == nil {
			break
		}

		return e.complexity.TodoConnection.Edges(childComplexity), true

	case "TodoConnection.pageInfo":
		if e.complexity.TodoConnection.PageInfo == nil {
			break
		}

		return e.complexity.TodoConnection.PageInfo(childComplexity), true

	case "TodoConnection.totalCount":
		if e.complexity.TodoConnection.TotalCount == nil {
			break
		}

		return e.complexity.TodoConnection.TotalCount(childComplexity), true

	case "TodoEdge.cursor":
		if e.complexity.TodoEdge.Cursor == nil {
			break
		}

		return e.complexity.TodoEdge.Cursor(childComplexity), true

	case "TodoEdge.node":
		if e.complexity.TodoEdge.Node == nil {
			break
		}

		return e.complexity.TodoEdge.Node(childComplexity), true

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

	case "UserConnection.edges":
		if e.complexity.UserConnection.Edges == nil {
			break
		}

		return e.complexity.UserConnection.Edges(childComplexity), true

	case "UserConnection.pageInfo":
		if e.complexity.UserConnection.PageInfo == nil {
			break
		}

		return e.complexity.UserConnection.PageInfo(childComplexity), true

	case "UserConnection.totalCount":
		if e.complexity.UserConnection.TotalCount == nil {
			break
		}

		return e.complexity.UserConnection.TotalCount(childComplexity), true

	case "UserEdge.cursor":
		if e.complexity.UserEdge.Cursor == nil {
			break
		}

		return e.complexity.UserEdge.Cursor(childComplexity), true

	case "UserEdge.node":
		if e.complexity.UserEdge.Node == nil {
			break
		}

		return e.complexity.UserEdge.Node(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e, 0, 0, make(chan graphql.DeferredResult)}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputCreateTodoInput,
		ec.unmarshalInputCreateUserInput,
		ec.unmarshalInputTodoOrder,
		ec.unmarshalInputTodoWhereInput,
		ec.unmarshalInputUpdateTodoInput,
		ec.unmarshalInputUpdateUserInput,
		ec.unmarshalInputUserOrder,
		ec.unmarshalInputUserWhereInput,
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
	{Name: "../../../ent.graphql", Input: `directive @goField(forceResolver: Boolean, name: String) on FIELD_DEFINITION | INPUT_FIELD_DEFINITION
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
  userID: ID
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
  todos(
    """Returns the elements in the list that come after the specified cursor."""
    after: Cursor

    """Returns the first _n_ elements from the list."""
    first: Int

    """Returns the elements in the list that come before the specified cursor."""
    before: Cursor

    """Returns the last _n_ elements from the list."""
    last: Int

    """Ordering options for Todos returned from the connection."""
    orderBy: [TodoOrder!]

    """Filtering options for Todos returned from the connection."""
    where: TodoWhereInput
  ): TodoConnection!
  users(
    """Returns the elements in the list that come after the specified cursor."""
    after: Cursor

    """Returns the first _n_ elements from the list."""
    first: Int

    """Returns the elements in the list that come before the specified cursor."""
    before: Cursor

    """Returns the last _n_ elements from the list."""
    last: Int

    """Ordering options for Users returned from the connection."""
    orderBy: [UserOrder!]

    """Filtering options for Users returned from the connection."""
    where: UserWhereInput
  ): UserConnection!
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
  user: User
}
"""A connection to a list of items."""
type TodoConnection {
  """A list of edges."""
  edges: [TodoEdge]
  """Information to aid in pagination."""
  pageInfo: PageInfo!
  """Identifies the total count of items in the connection."""
  totalCount: Int!
}
"""An edge in a connection."""
type TodoEdge {
  """The item at the end of the edge."""
  node: Todo
  """A cursor for use in pagination."""
  cursor: Cursor!
}
"""Ordering options for Todo connections"""
input TodoOrder {
  """The ordering direction."""
  direction: OrderDirection! = ASC
  """The field by which to order Todos."""
  field: TodoOrderField!
}
"""Properties by which Todo connections can be ordered."""
enum TodoOrderField {
  ID
  CREATED_AT
  UPDATED_AT
  TITLE
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
"""
TodoWhereInput is used for filtering Todo objects.
Input was generated by ent.
"""
input TodoWhereInput {
  not: TodoWhereInput
  and: [TodoWhereInput!]
  or: [TodoWhereInput!]
  """id field predicates"""
  id: ID
  idNEQ: ID
  idIn: [ID!]
  idNotIn: [ID!]
  idGT: ID
  idGTE: ID
  idLT: ID
  idLTE: ID
  """created_at field predicates"""
  createdAt: Time
  createdAtNEQ: Time
  createdAtIn: [Time!]
  createdAtNotIn: [Time!]
  createdAtGT: Time
  createdAtGTE: Time
  createdAtLT: Time
  createdAtLTE: Time
  """updated_at field predicates"""
  updatedAt: Time
  updatedAtNEQ: Time
  updatedAtIn: [Time!]
  updatedAtNotIn: [Time!]
  updatedAtGT: Time
  updatedAtGTE: Time
  updatedAtLT: Time
  updatedAtLTE: Time
  """title field predicates"""
  title: String
  titleNEQ: String
  titleIn: [String!]
  titleNotIn: [String!]
  titleGT: String
  titleGTE: String
  titleLT: String
  titleLTE: String
  titleContains: String
  titleHasPrefix: String
  titleHasSuffix: String
  titleEqualFold: String
  titleContainsFold: String
  """priority field predicates"""
  priority: TodoPriority
  priorityNEQ: TodoPriority
  priorityIn: [TodoPriority!]
  priorityNotIn: [TodoPriority!]
  """status field predicates"""
  status: TodoStatus
  statusNEQ: TodoStatus
  statusIn: [TodoStatus!]
  statusNotIn: [TodoStatus!]
  """time_completed field predicates"""
  timeCompleted: Time
  timeCompletedNEQ: Time
  timeCompletedIn: [Time!]
  timeCompletedNotIn: [Time!]
  timeCompletedGT: Time
  timeCompletedGTE: Time
  timeCompletedLT: Time
  timeCompletedLTE: Time
  timeCompletedIsNil: Boolean
  timeCompletedNotNil: Boolean
  """user edge predicates"""
  hasUser: Boolean
  hasUserWith: [UserWhereInput!]
}
"""
UpdateTodoInput is used for update Todo object.
Input was generated by ent.
"""
input UpdateTodoInput {
  updatedAt: Time
  title: String
  priority: TodoPriority
  status: TodoStatus
  timeCompleted: Time
  clearTimeCompleted: Boolean
  userID: ID
  clearUser: Boolean
}
"""
UpdateUserInput is used for update User object.
Input was generated by ent.
"""
input UpdateUserInput {
  updatedAt: Time
  username: String
  displayName: String
  addTodoIDs: [ID!]
  removeTodoIDs: [ID!]
  clearTodos: Boolean
}
type User implements Node {
  id: ID!
  createdAt: Time!
  updatedAt: Time!
  username: String!
  displayName: String!
  todos: [Todo!]
}
"""A connection to a list of items."""
type UserConnection {
  """A list of edges."""
  edges: [UserEdge]
  """Information to aid in pagination."""
  pageInfo: PageInfo!
  """Identifies the total count of items in the connection."""
  totalCount: Int!
}
"""An edge in a connection."""
type UserEdge {
  """The item at the end of the edge."""
  node: User
  """A cursor for use in pagination."""
  cursor: Cursor!
}
"""Ordering options for User connections"""
input UserOrder {
  """The ordering direction."""
  direction: OrderDirection! = ASC
  """The field by which to order Users."""
  field: UserOrderField!
}
"""Properties by which User connections can be ordered."""
enum UserOrderField {
  ID
  CREATED_AT
  UPDATED_AT
}
"""
UserWhereInput is used for filtering User objects.
Input was generated by ent.
"""
input UserWhereInput {
  not: UserWhereInput
  and: [UserWhereInput!]
  or: [UserWhereInput!]
  """id field predicates"""
  id: ID
  idNEQ: ID
  idIn: [ID!]
  idNotIn: [ID!]
  idGT: ID
  idGTE: ID
  idLT: ID
  idLTE: ID
  """created_at field predicates"""
  createdAt: Time
  createdAtNEQ: Time
  createdAtIn: [Time!]
  createdAtNotIn: [Time!]
  createdAtGT: Time
  createdAtGTE: Time
  createdAtLT: Time
  createdAtLTE: Time
  """updated_at field predicates"""
  updatedAt: Time
  updatedAtNEQ: Time
  updatedAtIn: [Time!]
  updatedAtNotIn: [Time!]
  updatedAtGT: Time
  updatedAtGTE: Time
  updatedAtLT: Time
  updatedAtLTE: Time
  """username field predicates"""
  username: String
  usernameNEQ: String
  usernameIn: [String!]
  usernameNotIn: [String!]
  usernameGT: String
  usernameGTE: String
  usernameLT: String
  usernameLTE: String
  usernameContains: String
  usernameHasPrefix: String
  usernameHasSuffix: String
  usernameEqualFold: String
  usernameContainsFold: String
  """display_name field predicates"""
  displayName: String
  displayNameNEQ: String
  displayNameIn: [String!]
  displayNameNotIn: [String!]
  displayNameGT: String
  displayNameGTE: String
  displayNameLT: String
  displayNameLTE: String
  displayNameContains: String
  displayNameHasPrefix: String
  displayNameHasSuffix: String
  displayNameEqualFold: String
  displayNameContainsFold: String
  """todos edge predicates"""
  hasTodos: Boolean
  hasTodosWith: [TodoWhereInput!]
}
`, BuiltIn: false},
	{Name: "../../../typedef.graphql", Input: `# All custom graphql definitions go here.
# Types from ` + "`" + `ent.graphql` + "`" + ` are referenced here

type Mutation {
  createTodo(input: CreateTodoInput!): Todo
  createUser(input: CreateUserInput!): User
}
`, BuiltIn: false},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
