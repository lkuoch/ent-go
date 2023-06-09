// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"lkuoch/ent-todo/ent/generated/predicate"
	"lkuoch/ent-todo/ent/generated/todo"
	"lkuoch/ent-todo/ent/generated/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TodoUpdate is the builder for updating Todo entities.
type TodoUpdate struct {
	config
	hooks    []Hook
	mutation *TodoMutation
}

// Where appends a list predicates to the TodoUpdate builder.
func (tu *TodoUpdate) Where(ps ...predicate.Todo) *TodoUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TodoUpdate) SetUpdatedAt(t time.Time) *TodoUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetTitle sets the "title" field.
func (tu *TodoUpdate) SetTitle(s string) *TodoUpdate {
	tu.mutation.SetTitle(s)
	return tu
}

// SetPriority sets the "priority" field.
func (tu *TodoUpdate) SetPriority(t todo.Priority) *TodoUpdate {
	tu.mutation.SetPriority(t)
	return tu
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (tu *TodoUpdate) SetNillablePriority(t *todo.Priority) *TodoUpdate {
	if t != nil {
		tu.SetPriority(*t)
	}
	return tu
}

// SetStatus sets the "status" field.
func (tu *TodoUpdate) SetStatus(t todo.Status) *TodoUpdate {
	tu.mutation.SetStatus(t)
	return tu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tu *TodoUpdate) SetNillableStatus(t *todo.Status) *TodoUpdate {
	if t != nil {
		tu.SetStatus(*t)
	}
	return tu
}

// SetTimeCompleted sets the "time_completed" field.
func (tu *TodoUpdate) SetTimeCompleted(t time.Time) *TodoUpdate {
	tu.mutation.SetTimeCompleted(t)
	return tu
}

// SetNillableTimeCompleted sets the "time_completed" field if the given value is not nil.
func (tu *TodoUpdate) SetNillableTimeCompleted(t *time.Time) *TodoUpdate {
	if t != nil {
		tu.SetTimeCompleted(*t)
	}
	return tu
}

// ClearTimeCompleted clears the value of the "time_completed" field.
func (tu *TodoUpdate) ClearTimeCompleted() *TodoUpdate {
	tu.mutation.ClearTimeCompleted()
	return tu
}

// AddChildIDs adds the "children" edge to the Todo entity by IDs.
func (tu *TodoUpdate) AddChildIDs(ids ...string) *TodoUpdate {
	tu.mutation.AddChildIDs(ids...)
	return tu
}

// AddChildren adds the "children" edges to the Todo entity.
func (tu *TodoUpdate) AddChildren(t ...*Todo) *TodoUpdate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddChildIDs(ids...)
}

// SetParentID sets the "parent" edge to the Todo entity by ID.
func (tu *TodoUpdate) SetParentID(id string) *TodoUpdate {
	tu.mutation.SetParentID(id)
	return tu
}

// SetNillableParentID sets the "parent" edge to the Todo entity by ID if the given value is not nil.
func (tu *TodoUpdate) SetNillableParentID(id *string) *TodoUpdate {
	if id != nil {
		tu = tu.SetParentID(*id)
	}
	return tu
}

// SetParent sets the "parent" edge to the Todo entity.
func (tu *TodoUpdate) SetParent(t *Todo) *TodoUpdate {
	return tu.SetParentID(t.ID)
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (tu *TodoUpdate) AddUserIDs(ids ...string) *TodoUpdate {
	tu.mutation.AddUserIDs(ids...)
	return tu
}

// AddUser adds the "user" edges to the User entity.
func (tu *TodoUpdate) AddUser(u ...*User) *TodoUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tu.AddUserIDs(ids...)
}

// Mutation returns the TodoMutation object of the builder.
func (tu *TodoUpdate) Mutation() *TodoMutation {
	return tu.mutation
}

// ClearChildren clears all "children" edges to the Todo entity.
func (tu *TodoUpdate) ClearChildren() *TodoUpdate {
	tu.mutation.ClearChildren()
	return tu
}

// RemoveChildIDs removes the "children" edge to Todo entities by IDs.
func (tu *TodoUpdate) RemoveChildIDs(ids ...string) *TodoUpdate {
	tu.mutation.RemoveChildIDs(ids...)
	return tu
}

// RemoveChildren removes "children" edges to Todo entities.
func (tu *TodoUpdate) RemoveChildren(t ...*Todo) *TodoUpdate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveChildIDs(ids...)
}

// ClearParent clears the "parent" edge to the Todo entity.
func (tu *TodoUpdate) ClearParent() *TodoUpdate {
	tu.mutation.ClearParent()
	return tu
}

// ClearUser clears all "user" edges to the User entity.
func (tu *TodoUpdate) ClearUser() *TodoUpdate {
	tu.mutation.ClearUser()
	return tu
}

// RemoveUserIDs removes the "user" edge to User entities by IDs.
func (tu *TodoUpdate) RemoveUserIDs(ids ...string) *TodoUpdate {
	tu.mutation.RemoveUserIDs(ids...)
	return tu
}

// RemoveUser removes "user" edges to User entities.
func (tu *TodoUpdate) RemoveUser(u ...*User) *TodoUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tu.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TodoUpdate) Save(ctx context.Context) (int, error) {
	tu.defaults()
	return withHooks[int, TodoMutation](ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TodoUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TodoUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TodoUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TodoUpdate) defaults() {
	if _, ok := tu.mutation.UpdatedAt(); !ok {
		v := todo.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TodoUpdate) check() error {
	if v, ok := tu.mutation.Title(); ok {
		if err := todo.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`generated: validator failed for field "Todo.title": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Priority(); ok {
		if err := todo.PriorityValidator(v); err != nil {
			return &ValidationError{Name: "priority", err: fmt.Errorf(`generated: validator failed for field "Todo.priority": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Status(); ok {
		if err := todo.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`generated: validator failed for field "Todo.status": %w`, err)}
		}
	}
	return nil
}

func (tu *TodoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(todo.Table, todo.Columns, sqlgraph.NewFieldSpec(todo.FieldID, field.TypeString))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(todo.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.Title(); ok {
		_spec.SetField(todo.FieldTitle, field.TypeString, value)
	}
	if value, ok := tu.mutation.Priority(); ok {
		_spec.SetField(todo.FieldPriority, field.TypeEnum, value)
	}
	if value, ok := tu.mutation.Status(); ok {
		_spec.SetField(todo.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := tu.mutation.TimeCompleted(); ok {
		_spec.SetField(todo.FieldTimeCompleted, field.TypeTime, value)
	}
	if tu.mutation.TimeCompletedCleared() {
		_spec.ClearField(todo.FieldTimeCompleted, field.TypeTime)
	}
	if tu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   todo.ChildrenTable,
			Columns: []string{todo.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !tu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   todo.ChildrenTable,
			Columns: []string{todo.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   todo.ChildrenTable,
			Columns: []string{todo.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   todo.ParentTable,
			Columns: []string{todo.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   todo.ParentTable,
			Columns: []string{todo.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   todo.UserTable,
			Columns: todo.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedUserIDs(); len(nodes) > 0 && !tu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   todo.UserTable,
			Columns: todo.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   todo.UserTable,
			Columns: todo.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{todo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TodoUpdateOne is the builder for updating a single Todo entity.
type TodoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TodoMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TodoUpdateOne) SetUpdatedAt(t time.Time) *TodoUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetTitle sets the "title" field.
func (tuo *TodoUpdateOne) SetTitle(s string) *TodoUpdateOne {
	tuo.mutation.SetTitle(s)
	return tuo
}

// SetPriority sets the "priority" field.
func (tuo *TodoUpdateOne) SetPriority(t todo.Priority) *TodoUpdateOne {
	tuo.mutation.SetPriority(t)
	return tuo
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillablePriority(t *todo.Priority) *TodoUpdateOne {
	if t != nil {
		tuo.SetPriority(*t)
	}
	return tuo
}

// SetStatus sets the "status" field.
func (tuo *TodoUpdateOne) SetStatus(t todo.Status) *TodoUpdateOne {
	tuo.mutation.SetStatus(t)
	return tuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableStatus(t *todo.Status) *TodoUpdateOne {
	if t != nil {
		tuo.SetStatus(*t)
	}
	return tuo
}

// SetTimeCompleted sets the "time_completed" field.
func (tuo *TodoUpdateOne) SetTimeCompleted(t time.Time) *TodoUpdateOne {
	tuo.mutation.SetTimeCompleted(t)
	return tuo
}

// SetNillableTimeCompleted sets the "time_completed" field if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableTimeCompleted(t *time.Time) *TodoUpdateOne {
	if t != nil {
		tuo.SetTimeCompleted(*t)
	}
	return tuo
}

// ClearTimeCompleted clears the value of the "time_completed" field.
func (tuo *TodoUpdateOne) ClearTimeCompleted() *TodoUpdateOne {
	tuo.mutation.ClearTimeCompleted()
	return tuo
}

// AddChildIDs adds the "children" edge to the Todo entity by IDs.
func (tuo *TodoUpdateOne) AddChildIDs(ids ...string) *TodoUpdateOne {
	tuo.mutation.AddChildIDs(ids...)
	return tuo
}

// AddChildren adds the "children" edges to the Todo entity.
func (tuo *TodoUpdateOne) AddChildren(t ...*Todo) *TodoUpdateOne {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddChildIDs(ids...)
}

// SetParentID sets the "parent" edge to the Todo entity by ID.
func (tuo *TodoUpdateOne) SetParentID(id string) *TodoUpdateOne {
	tuo.mutation.SetParentID(id)
	return tuo
}

// SetNillableParentID sets the "parent" edge to the Todo entity by ID if the given value is not nil.
func (tuo *TodoUpdateOne) SetNillableParentID(id *string) *TodoUpdateOne {
	if id != nil {
		tuo = tuo.SetParentID(*id)
	}
	return tuo
}

// SetParent sets the "parent" edge to the Todo entity.
func (tuo *TodoUpdateOne) SetParent(t *Todo) *TodoUpdateOne {
	return tuo.SetParentID(t.ID)
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (tuo *TodoUpdateOne) AddUserIDs(ids ...string) *TodoUpdateOne {
	tuo.mutation.AddUserIDs(ids...)
	return tuo
}

// AddUser adds the "user" edges to the User entity.
func (tuo *TodoUpdateOne) AddUser(u ...*User) *TodoUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tuo.AddUserIDs(ids...)
}

// Mutation returns the TodoMutation object of the builder.
func (tuo *TodoUpdateOne) Mutation() *TodoMutation {
	return tuo.mutation
}

// ClearChildren clears all "children" edges to the Todo entity.
func (tuo *TodoUpdateOne) ClearChildren() *TodoUpdateOne {
	tuo.mutation.ClearChildren()
	return tuo
}

// RemoveChildIDs removes the "children" edge to Todo entities by IDs.
func (tuo *TodoUpdateOne) RemoveChildIDs(ids ...string) *TodoUpdateOne {
	tuo.mutation.RemoveChildIDs(ids...)
	return tuo
}

// RemoveChildren removes "children" edges to Todo entities.
func (tuo *TodoUpdateOne) RemoveChildren(t ...*Todo) *TodoUpdateOne {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveChildIDs(ids...)
}

// ClearParent clears the "parent" edge to the Todo entity.
func (tuo *TodoUpdateOne) ClearParent() *TodoUpdateOne {
	tuo.mutation.ClearParent()
	return tuo
}

// ClearUser clears all "user" edges to the User entity.
func (tuo *TodoUpdateOne) ClearUser() *TodoUpdateOne {
	tuo.mutation.ClearUser()
	return tuo
}

// RemoveUserIDs removes the "user" edge to User entities by IDs.
func (tuo *TodoUpdateOne) RemoveUserIDs(ids ...string) *TodoUpdateOne {
	tuo.mutation.RemoveUserIDs(ids...)
	return tuo
}

// RemoveUser removes "user" edges to User entities.
func (tuo *TodoUpdateOne) RemoveUser(u ...*User) *TodoUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tuo.RemoveUserIDs(ids...)
}

// Where appends a list predicates to the TodoUpdate builder.
func (tuo *TodoUpdateOne) Where(ps ...predicate.Todo) *TodoUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TodoUpdateOne) Select(field string, fields ...string) *TodoUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Todo entity.
func (tuo *TodoUpdateOne) Save(ctx context.Context) (*Todo, error) {
	tuo.defaults()
	return withHooks[*Todo, TodoMutation](ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TodoUpdateOne) SaveX(ctx context.Context) *Todo {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TodoUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TodoUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TodoUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdatedAt(); !ok {
		v := todo.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TodoUpdateOne) check() error {
	if v, ok := tuo.mutation.Title(); ok {
		if err := todo.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`generated: validator failed for field "Todo.title": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Priority(); ok {
		if err := todo.PriorityValidator(v); err != nil {
			return &ValidationError{Name: "priority", err: fmt.Errorf(`generated: validator failed for field "Todo.priority": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Status(); ok {
		if err := todo.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`generated: validator failed for field "Todo.status": %w`, err)}
		}
	}
	return nil
}

func (tuo *TodoUpdateOne) sqlSave(ctx context.Context) (_node *Todo, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(todo.Table, todo.Columns, sqlgraph.NewFieldSpec(todo.FieldID, field.TypeString))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Todo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, todo.FieldID)
		for _, f := range fields {
			if !todo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != todo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(todo.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.Title(); ok {
		_spec.SetField(todo.FieldTitle, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Priority(); ok {
		_spec.SetField(todo.FieldPriority, field.TypeEnum, value)
	}
	if value, ok := tuo.mutation.Status(); ok {
		_spec.SetField(todo.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := tuo.mutation.TimeCompleted(); ok {
		_spec.SetField(todo.FieldTimeCompleted, field.TypeTime, value)
	}
	if tuo.mutation.TimeCompletedCleared() {
		_spec.ClearField(todo.FieldTimeCompleted, field.TypeTime)
	}
	if tuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   todo.ChildrenTable,
			Columns: []string{todo.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !tuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   todo.ChildrenTable,
			Columns: []string{todo.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   todo.ChildrenTable,
			Columns: []string{todo.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   todo.ParentTable,
			Columns: []string{todo.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   todo.ParentTable,
			Columns: []string{todo.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(todo.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   todo.UserTable,
			Columns: todo.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedUserIDs(); len(nodes) > 0 && !tuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   todo.UserTable,
			Columns: todo.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   todo.UserTable,
			Columns: todo.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Todo{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{todo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
