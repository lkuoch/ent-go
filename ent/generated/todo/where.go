// Code generated by ent, DO NOT EDIT.

package todo

import (
	"lkuoch/ent-todo/ent/generated/predicate"
	"lkuoch/ent-todo/ent/schema/types/pulid"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id pulid.ID) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id pulid.ID) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id pulid.ID) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...pulid.ID) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...pulid.ID) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id pulid.ID) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id pulid.ID) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id pulid.ID) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id pulid.ID) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldUpdatedAt, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldTitle, v))
}

// TimeCompleted applies equality check predicate on the "time_completed" field. It's identical to TimeCompletedEQ.
func TimeCompleted(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldTimeCompleted, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldUpdatedAt, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Todo {
	return predicate.Todo(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Todo {
	return predicate.Todo(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Todo {
	return predicate.Todo(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Todo {
	return predicate.Todo(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Todo {
	return predicate.Todo(sql.FieldContainsFold(FieldTitle, v))
}

// PriorityEQ applies the EQ predicate on the "priority" field.
func PriorityEQ(v Priority) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldPriority, v))
}

// PriorityNEQ applies the NEQ predicate on the "priority" field.
func PriorityNEQ(v Priority) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldPriority, v))
}

// PriorityIn applies the In predicate on the "priority" field.
func PriorityIn(vs ...Priority) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldPriority, vs...))
}

// PriorityNotIn applies the NotIn predicate on the "priority" field.
func PriorityNotIn(vs ...Priority) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldPriority, vs...))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldStatus, vs...))
}

// TimeCompletedEQ applies the EQ predicate on the "time_completed" field.
func TimeCompletedEQ(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldTimeCompleted, v))
}

// TimeCompletedNEQ applies the NEQ predicate on the "time_completed" field.
func TimeCompletedNEQ(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldTimeCompleted, v))
}

// TimeCompletedIn applies the In predicate on the "time_completed" field.
func TimeCompletedIn(vs ...time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldTimeCompleted, vs...))
}

// TimeCompletedNotIn applies the NotIn predicate on the "time_completed" field.
func TimeCompletedNotIn(vs ...time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldTimeCompleted, vs...))
}

// TimeCompletedGT applies the GT predicate on the "time_completed" field.
func TimeCompletedGT(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldTimeCompleted, v))
}

// TimeCompletedGTE applies the GTE predicate on the "time_completed" field.
func TimeCompletedGTE(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldTimeCompleted, v))
}

// TimeCompletedLT applies the LT predicate on the "time_completed" field.
func TimeCompletedLT(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldTimeCompleted, v))
}

// TimeCompletedLTE applies the LTE predicate on the "time_completed" field.
func TimeCompletedLTE(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldTimeCompleted, v))
}

// TimeCompletedIsNil applies the IsNil predicate on the "time_completed" field.
func TimeCompletedIsNil() predicate.Todo {
	return predicate.Todo(sql.FieldIsNull(FieldTimeCompleted))
}

// TimeCompletedNotNil applies the NotNil predicate on the "time_completed" field.
func TimeCompletedNotNil() predicate.Todo {
	return predicate.Todo(sql.FieldNotNull(FieldTimeCompleted))
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.Todo) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		step := newChildrenStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.Todo) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		step := newParentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UserTable, UserPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Todo) predicate.Todo {
	return predicate.Todo(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Todo) predicate.Todo {
	return predicate.Todo(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Todo) predicate.Todo {
	return predicate.Todo(sql.NotPredicates(p))
}
