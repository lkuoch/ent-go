// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"io"
	"lkuoch/ent-todo/ent/generated/todo"
	"lkuoch/ent-todo/ent/generated/user"
	"lkuoch/ent-todo/ent/schema/types/pulid"
	"strconv"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Common entgql types.
type (
	Cursor         = entgql.Cursor[pulid.ID]
	PageInfo       = entgql.PageInfo[pulid.ID]
	OrderDirection = entgql.OrderDirection
)

func orderFunc(o OrderDirection, field string) func(*sql.Selector) {
	if o == entgql.OrderDirectionDesc {
		return Desc(field)
	}
	return Asc(field)
}

const errInvalidPagination = "INVALID_PAGINATION"

func validateFirstLast(first, last *int) (err *gqlerror.Error) {
	switch {
	case first != nil && last != nil:
		err = &gqlerror.Error{
			Message: "Passing both `first` and `last` to paginate a connection is not supported.",
		}
	case first != nil && *first < 0:
		err = &gqlerror.Error{
			Message: "`first` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	case last != nil && *last < 0:
		err = &gqlerror.Error{
			Message: "`last` on a connection cannot be less than zero.",
		}
		errcode.Set(err, errInvalidPagination)
	}
	return err
}

func collectedField(ctx context.Context, path ...string) *graphql.CollectedField {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	field := fc.Field
	oc := graphql.GetOperationContext(ctx)
walk:
	for _, name := range path {
		for _, f := range graphql.CollectFields(oc, field.Selections, nil) {
			if f.Alias == name {
				field = f
				continue walk
			}
		}
		return nil
	}
	return &field
}

func hasCollectedField(ctx context.Context, path ...string) bool {
	if graphql.GetFieldContext(ctx) == nil {
		return true
	}
	return collectedField(ctx, path...) != nil
}

const (
	edgesField      = "edges"
	nodeField       = "node"
	pageInfoField   = "pageInfo"
	totalCountField = "totalCount"
)

func paginateLimit(first, last *int) int {
	var limit int
	if first != nil {
		limit = *first + 1
	} else if last != nil {
		limit = *last + 1
	}
	return limit
}

// TodoEdge is the edge representation of Todo.
type TodoEdge struct {
	Node   *Todo  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// TodoConnection is the connection containing edges to Todo.
type TodoConnection struct {
	Edges      []*TodoEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

func (c *TodoConnection) build(nodes []*Todo, pager *todoPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *Todo
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *Todo {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *Todo {
			return nodes[i]
		}
	}
	c.Edges = make([]*TodoEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &TodoEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// TodoPaginateOption enables pagination customization.
type TodoPaginateOption func(*todoPager) error

// WithTodoOrder configures pagination ordering.
func WithTodoOrder(order []*TodoOrder) TodoPaginateOption {
	return func(pager *todoPager) error {
		for _, o := range order {
			if err := o.Direction.Validate(); err != nil {
				return err
			}
		}
		pager.order = append(pager.order, order...)
		return nil
	}
}

// WithTodoFilter configures pagination filter.
func WithTodoFilter(filter func(*TodoQuery) (*TodoQuery, error)) TodoPaginateOption {
	return func(pager *todoPager) error {
		if filter == nil {
			return errors.New("TodoQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type todoPager struct {
	reverse bool
	order   []*TodoOrder
	filter  func(*TodoQuery) (*TodoQuery, error)
}

func newTodoPager(opts []TodoPaginateOption, reverse bool) (*todoPager, error) {
	pager := &todoPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	for i, o := range pager.order {
		if i > 0 && o.Field == pager.order[i-1].Field {
			return nil, fmt.Errorf("duplicate order direction %q", o.Direction)
		}
	}
	return pager, nil
}

func (p *todoPager) applyFilter(query *TodoQuery) (*TodoQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *todoPager) toCursor(t *Todo) Cursor {
	cs := make([]any, 0, len(p.order))
	for _, o := range p.order {
		cs = append(cs, o.Field.toCursor(t).Value)
	}
	return Cursor{ID: t.ID, Value: cs}
}

func (p *todoPager) applyCursors(query *TodoQuery, after, before *Cursor) (*TodoQuery, error) {
	idDirection := entgql.OrderDirectionAsc
	if p.reverse {
		idDirection = entgql.OrderDirectionDesc
	}
	fields, directions := make([]string, 0, len(p.order)), make([]OrderDirection, 0, len(p.order))
	for _, o := range p.order {
		fields = append(fields, o.Field.column)
		direction := o.Direction
		if p.reverse {
			direction = direction.Reverse()
		}
		directions = append(directions, direction)
	}
	predicates, err := entgql.MultiCursorsPredicate(after, before, &entgql.MultiCursorsOptions{
		FieldID:     DefaultTodoOrder.Field.column,
		DirectionID: idDirection,
		Fields:      fields,
		Directions:  directions,
	})
	if err != nil {
		return nil, err
	}
	for _, predicate := range predicates {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *todoPager) applyOrder(query *TodoQuery) *TodoQuery {
	var defaultOrdered bool
	for _, o := range p.order {
		direction := o.Direction
		if p.reverse {
			direction = direction.Reverse()
		}
		query = query.Order(o.Field.toTerm(direction.OrderTermOption()))
		if o.Field.column == DefaultTodoOrder.Field.column {
			defaultOrdered = true
		}
		if len(query.ctx.Fields) > 0 {
			query.ctx.AppendFieldOnce(o.Field.column)
		}
	}
	if !defaultOrdered {
		direction := entgql.OrderDirectionAsc
		if p.reverse {
			direction = direction.Reverse()
		}
		query = query.Order(DefaultTodoOrder.Field.toTerm(direction.OrderTermOption()))
	}
	return query
}

func (p *todoPager) orderExpr(query *TodoQuery) sql.Querier {
	if len(query.ctx.Fields) > 0 {
		for _, o := range p.order {
			query.ctx.AppendFieldOnce(o.Field.column)
		}
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		for _, o := range p.order {
			direction := o.Direction
			if p.reverse {
				direction = direction.Reverse()
			}
			b.Ident(o.Field.column).Pad().WriteString(string(direction))
			b.Comma()
		}
		direction := entgql.OrderDirectionAsc
		if p.reverse {
			direction = direction.Reverse()
		}
		b.Ident(DefaultTodoOrder.Field.column).Pad().WriteString(string(direction))
	})
}

// Paginate executes the query and returns a relay based cursor connection to Todo.
func (t *TodoQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...TodoPaginateOption,
) (*TodoConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newTodoPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if t, err = pager.applyFilter(t); err != nil {
		return nil, err
	}
	conn := &TodoConnection{Edges: []*TodoEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = t.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if t, err = pager.applyCursors(t, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		t.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := t.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	t = pager.applyOrder(t)
	nodes, err := t.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// TodoOrderFieldID orders Todo by id.
	TodoOrderFieldID = &TodoOrderField{
		Value: func(t *Todo) (ent.Value, error) {
			return t.ID, nil
		},
		column: todo.FieldID,
		toTerm: todo.ByID,
		toCursor: func(t *Todo) Cursor {
			return Cursor{
				ID:    t.ID,
				Value: t.ID,
			}
		},
	}
	// TodoOrderFieldCreatedAt orders Todo by created_at.
	TodoOrderFieldCreatedAt = &TodoOrderField{
		Value: func(t *Todo) (ent.Value, error) {
			return t.CreatedAt, nil
		},
		column: todo.FieldCreatedAt,
		toTerm: todo.ByCreatedAt,
		toCursor: func(t *Todo) Cursor {
			return Cursor{
				ID:    t.ID,
				Value: t.CreatedAt,
			}
		},
	}
	// TodoOrderFieldUpdatedAt orders Todo by updated_at.
	TodoOrderFieldUpdatedAt = &TodoOrderField{
		Value: func(t *Todo) (ent.Value, error) {
			return t.UpdatedAt, nil
		},
		column: todo.FieldUpdatedAt,
		toTerm: todo.ByUpdatedAt,
		toCursor: func(t *Todo) Cursor {
			return Cursor{
				ID:    t.ID,
				Value: t.UpdatedAt,
			}
		},
	}
	// TodoOrderFieldTitle orders Todo by title.
	TodoOrderFieldTitle = &TodoOrderField{
		Value: func(t *Todo) (ent.Value, error) {
			return t.Title, nil
		},
		column: todo.FieldTitle,
		toTerm: todo.ByTitle,
		toCursor: func(t *Todo) Cursor {
			return Cursor{
				ID:    t.ID,
				Value: t.Title,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f TodoOrderField) String() string {
	var str string
	switch f.column {
	case TodoOrderFieldID.column:
		str = "ID"
	case TodoOrderFieldCreatedAt.column:
		str = "CREATED_AT"
	case TodoOrderFieldUpdatedAt.column:
		str = "UPDATED_AT"
	case TodoOrderFieldTitle.column:
		str = "TITLE"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f TodoOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *TodoOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("TodoOrderField %T must be a string", v)
	}
	switch str {
	case "ID":
		*f = *TodoOrderFieldID
	case "CREATED_AT":
		*f = *TodoOrderFieldCreatedAt
	case "UPDATED_AT":
		*f = *TodoOrderFieldUpdatedAt
	case "TITLE":
		*f = *TodoOrderFieldTitle
	default:
		return fmt.Errorf("%s is not a valid TodoOrderField", str)
	}
	return nil
}

// TodoOrderField defines the ordering field of Todo.
type TodoOrderField struct {
	// Value extracts the ordering value from the given Todo.
	Value    func(*Todo) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) todo.OrderOption
	toCursor func(*Todo) Cursor
}

// TodoOrder defines the ordering of Todo.
type TodoOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *TodoOrderField `json:"field"`
}

// DefaultTodoOrder is the default ordering of Todo.
var DefaultTodoOrder = &TodoOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &TodoOrderField{
		Value: func(t *Todo) (ent.Value, error) {
			return t.ID, nil
		},
		column: todo.FieldID,
		toTerm: todo.ByID,
		toCursor: func(t *Todo) Cursor {
			return Cursor{ID: t.ID}
		},
	},
}

// ToEdge converts Todo into TodoEdge.
func (t *Todo) ToEdge(order *TodoOrder) *TodoEdge {
	if order == nil {
		order = DefaultTodoOrder
	}
	return &TodoEdge{
		Node:   t,
		Cursor: order.Field.toCursor(t),
	}
}

// UserEdge is the edge representation of User.
type UserEdge struct {
	Node   *User  `json:"node"`
	Cursor Cursor `json:"cursor"`
}

// UserConnection is the connection containing edges to User.
type UserConnection struct {
	Edges      []*UserEdge `json:"edges"`
	PageInfo   PageInfo    `json:"pageInfo"`
	TotalCount int         `json:"totalCount"`
}

func (c *UserConnection) build(nodes []*User, pager *userPager, after *Cursor, first *int, before *Cursor, last *int) {
	c.PageInfo.HasNextPage = before != nil
	c.PageInfo.HasPreviousPage = after != nil
	if first != nil && *first+1 == len(nodes) {
		c.PageInfo.HasNextPage = true
		nodes = nodes[:len(nodes)-1]
	} else if last != nil && *last+1 == len(nodes) {
		c.PageInfo.HasPreviousPage = true
		nodes = nodes[:len(nodes)-1]
	}
	var nodeAt func(int) *User
	if last != nil {
		n := len(nodes) - 1
		nodeAt = func(i int) *User {
			return nodes[n-i]
		}
	} else {
		nodeAt = func(i int) *User {
			return nodes[i]
		}
	}
	c.Edges = make([]*UserEdge, len(nodes))
	for i := range nodes {
		node := nodeAt(i)
		c.Edges[i] = &UserEdge{
			Node:   node,
			Cursor: pager.toCursor(node),
		}
	}
	if l := len(c.Edges); l > 0 {
		c.PageInfo.StartCursor = &c.Edges[0].Cursor
		c.PageInfo.EndCursor = &c.Edges[l-1].Cursor
	}
	if c.TotalCount == 0 {
		c.TotalCount = len(nodes)
	}
}

// UserPaginateOption enables pagination customization.
type UserPaginateOption func(*userPager) error

// WithUserOrder configures pagination ordering.
func WithUserOrder(order []*UserOrder) UserPaginateOption {
	return func(pager *userPager) error {
		for _, o := range order {
			if err := o.Direction.Validate(); err != nil {
				return err
			}
		}
		pager.order = append(pager.order, order...)
		return nil
	}
}

// WithUserFilter configures pagination filter.
func WithUserFilter(filter func(*UserQuery) (*UserQuery, error)) UserPaginateOption {
	return func(pager *userPager) error {
		if filter == nil {
			return errors.New("UserQuery filter cannot be nil")
		}
		pager.filter = filter
		return nil
	}
}

type userPager struct {
	reverse bool
	order   []*UserOrder
	filter  func(*UserQuery) (*UserQuery, error)
}

func newUserPager(opts []UserPaginateOption, reverse bool) (*userPager, error) {
	pager := &userPager{reverse: reverse}
	for _, opt := range opts {
		if err := opt(pager); err != nil {
			return nil, err
		}
	}
	for i, o := range pager.order {
		if i > 0 && o.Field == pager.order[i-1].Field {
			return nil, fmt.Errorf("duplicate order direction %q", o.Direction)
		}
	}
	return pager, nil
}

func (p *userPager) applyFilter(query *UserQuery) (*UserQuery, error) {
	if p.filter != nil {
		return p.filter(query)
	}
	return query, nil
}

func (p *userPager) toCursor(u *User) Cursor {
	cs := make([]any, 0, len(p.order))
	for _, o := range p.order {
		cs = append(cs, o.Field.toCursor(u).Value)
	}
	return Cursor{ID: u.ID, Value: cs}
}

func (p *userPager) applyCursors(query *UserQuery, after, before *Cursor) (*UserQuery, error) {
	idDirection := entgql.OrderDirectionAsc
	if p.reverse {
		idDirection = entgql.OrderDirectionDesc
	}
	fields, directions := make([]string, 0, len(p.order)), make([]OrderDirection, 0, len(p.order))
	for _, o := range p.order {
		fields = append(fields, o.Field.column)
		direction := o.Direction
		if p.reverse {
			direction = direction.Reverse()
		}
		directions = append(directions, direction)
	}
	predicates, err := entgql.MultiCursorsPredicate(after, before, &entgql.MultiCursorsOptions{
		FieldID:     DefaultUserOrder.Field.column,
		DirectionID: idDirection,
		Fields:      fields,
		Directions:  directions,
	})
	if err != nil {
		return nil, err
	}
	for _, predicate := range predicates {
		query = query.Where(predicate)
	}
	return query, nil
}

func (p *userPager) applyOrder(query *UserQuery) *UserQuery {
	var defaultOrdered bool
	for _, o := range p.order {
		direction := o.Direction
		if p.reverse {
			direction = direction.Reverse()
		}
		query = query.Order(o.Field.toTerm(direction.OrderTermOption()))
		if o.Field.column == DefaultUserOrder.Field.column {
			defaultOrdered = true
		}
		if len(query.ctx.Fields) > 0 {
			query.ctx.AppendFieldOnce(o.Field.column)
		}
	}
	if !defaultOrdered {
		direction := entgql.OrderDirectionAsc
		if p.reverse {
			direction = direction.Reverse()
		}
		query = query.Order(DefaultUserOrder.Field.toTerm(direction.OrderTermOption()))
	}
	return query
}

func (p *userPager) orderExpr(query *UserQuery) sql.Querier {
	if len(query.ctx.Fields) > 0 {
		for _, o := range p.order {
			query.ctx.AppendFieldOnce(o.Field.column)
		}
	}
	return sql.ExprFunc(func(b *sql.Builder) {
		for _, o := range p.order {
			direction := o.Direction
			if p.reverse {
				direction = direction.Reverse()
			}
			b.Ident(o.Field.column).Pad().WriteString(string(direction))
			b.Comma()
		}
		direction := entgql.OrderDirectionAsc
		if p.reverse {
			direction = direction.Reverse()
		}
		b.Ident(DefaultUserOrder.Field.column).Pad().WriteString(string(direction))
	})
}

// Paginate executes the query and returns a relay based cursor connection to User.
func (u *UserQuery) Paginate(
	ctx context.Context, after *Cursor, first *int,
	before *Cursor, last *int, opts ...UserPaginateOption,
) (*UserConnection, error) {
	if err := validateFirstLast(first, last); err != nil {
		return nil, err
	}
	pager, err := newUserPager(opts, last != nil)
	if err != nil {
		return nil, err
	}
	if u, err = pager.applyFilter(u); err != nil {
		return nil, err
	}
	conn := &UserConnection{Edges: []*UserEdge{}}
	ignoredEdges := !hasCollectedField(ctx, edgesField)
	if hasCollectedField(ctx, totalCountField) || hasCollectedField(ctx, pageInfoField) {
		hasPagination := after != nil || first != nil || before != nil || last != nil
		if hasPagination || ignoredEdges {
			if conn.TotalCount, err = u.Clone().Count(ctx); err != nil {
				return nil, err
			}
			conn.PageInfo.HasNextPage = first != nil && conn.TotalCount > 0
			conn.PageInfo.HasPreviousPage = last != nil && conn.TotalCount > 0
		}
	}
	if ignoredEdges || (first != nil && *first == 0) || (last != nil && *last == 0) {
		return conn, nil
	}
	if u, err = pager.applyCursors(u, after, before); err != nil {
		return nil, err
	}
	if limit := paginateLimit(first, last); limit != 0 {
		u.Limit(limit)
	}
	if field := collectedField(ctx, edgesField, nodeField); field != nil {
		if err := u.collectField(ctx, graphql.GetOperationContext(ctx), *field, []string{edgesField, nodeField}); err != nil {
			return nil, err
		}
	}
	u = pager.applyOrder(u)
	nodes, err := u.All(ctx)
	if err != nil {
		return nil, err
	}
	conn.build(nodes, pager, after, first, before, last)
	return conn, nil
}

var (
	// UserOrderFieldID orders User by id.
	UserOrderFieldID = &UserOrderField{
		Value: func(u *User) (ent.Value, error) {
			return u.ID, nil
		},
		column: user.FieldID,
		toTerm: user.ByID,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.ID,
			}
		},
	}
	// UserOrderFieldCreatedAt orders User by created_at.
	UserOrderFieldCreatedAt = &UserOrderField{
		Value: func(u *User) (ent.Value, error) {
			return u.CreatedAt, nil
		},
		column: user.FieldCreatedAt,
		toTerm: user.ByCreatedAt,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.CreatedAt,
			}
		},
	}
	// UserOrderFieldUpdatedAt orders User by updated_at.
	UserOrderFieldUpdatedAt = &UserOrderField{
		Value: func(u *User) (ent.Value, error) {
			return u.UpdatedAt, nil
		},
		column: user.FieldUpdatedAt,
		toTerm: user.ByUpdatedAt,
		toCursor: func(u *User) Cursor {
			return Cursor{
				ID:    u.ID,
				Value: u.UpdatedAt,
			}
		},
	}
)

// String implement fmt.Stringer interface.
func (f UserOrderField) String() string {
	var str string
	switch f.column {
	case UserOrderFieldID.column:
		str = "ID"
	case UserOrderFieldCreatedAt.column:
		str = "CREATED_AT"
	case UserOrderFieldUpdatedAt.column:
		str = "UPDATED_AT"
	}
	return str
}

// MarshalGQL implements graphql.Marshaler interface.
func (f UserOrderField) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(f.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (f *UserOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("UserOrderField %T must be a string", v)
	}
	switch str {
	case "ID":
		*f = *UserOrderFieldID
	case "CREATED_AT":
		*f = *UserOrderFieldCreatedAt
	case "UPDATED_AT":
		*f = *UserOrderFieldUpdatedAt
	default:
		return fmt.Errorf("%s is not a valid UserOrderField", str)
	}
	return nil
}

// UserOrderField defines the ordering field of User.
type UserOrderField struct {
	// Value extracts the ordering value from the given User.
	Value    func(*User) (ent.Value, error)
	column   string // field or computed.
	toTerm   func(...sql.OrderTermOption) user.OrderOption
	toCursor func(*User) Cursor
}

// UserOrder defines the ordering of User.
type UserOrder struct {
	Direction OrderDirection  `json:"direction"`
	Field     *UserOrderField `json:"field"`
}

// DefaultUserOrder is the default ordering of User.
var DefaultUserOrder = &UserOrder{
	Direction: entgql.OrderDirectionAsc,
	Field: &UserOrderField{
		Value: func(u *User) (ent.Value, error) {
			return u.ID, nil
		},
		column: user.FieldID,
		toTerm: user.ByID,
		toCursor: func(u *User) Cursor {
			return Cursor{ID: u.ID}
		},
	},
}

// ToEdge converts User into UserEdge.
func (u *User) ToEdge(order *UserOrder) *UserEdge {
	if order == nil {
		order = DefaultUserOrder
	}
	return &UserEdge{
		Node:   u,
		Cursor: order.Field.toCursor(u),
	}
}
