// Code generated by ent, DO NOT EDIT.

package daemon

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/hvturingga/ya/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Daemon {
	return predicate.Daemon(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Daemon {
	return predicate.Daemon(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Daemon {
	return predicate.Daemon(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Daemon {
	return predicate.Daemon(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Daemon {
	return predicate.Daemon(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Daemon {
	return predicate.Daemon(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Daemon {
	return predicate.Daemon(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Daemon {
	return predicate.Daemon(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Daemon {
	return predicate.Daemon(sql.FieldLTE(FieldID, id))
}

// Path applies equality check predicate on the "path" field. It's identical to PathEQ.
func Path(v string) predicate.Daemon {
	return predicate.Daemon(sql.FieldEQ(FieldPath, v))
}

// Enable applies equality check predicate on the "enable" field. It's identical to EnableEQ.
func Enable(v bool) predicate.Daemon {
	return predicate.Daemon(sql.FieldEQ(FieldEnable, v))
}

// PathEQ applies the EQ predicate on the "path" field.
func PathEQ(v string) predicate.Daemon {
	return predicate.Daemon(sql.FieldEQ(FieldPath, v))
}

// PathNEQ applies the NEQ predicate on the "path" field.
func PathNEQ(v string) predicate.Daemon {
	return predicate.Daemon(sql.FieldNEQ(FieldPath, v))
}

// PathIn applies the In predicate on the "path" field.
func PathIn(vs ...string) predicate.Daemon {
	return predicate.Daemon(sql.FieldIn(FieldPath, vs...))
}

// PathNotIn applies the NotIn predicate on the "path" field.
func PathNotIn(vs ...string) predicate.Daemon {
	return predicate.Daemon(sql.FieldNotIn(FieldPath, vs...))
}

// PathGT applies the GT predicate on the "path" field.
func PathGT(v string) predicate.Daemon {
	return predicate.Daemon(sql.FieldGT(FieldPath, v))
}

// PathGTE applies the GTE predicate on the "path" field.
func PathGTE(v string) predicate.Daemon {
	return predicate.Daemon(sql.FieldGTE(FieldPath, v))
}

// PathLT applies the LT predicate on the "path" field.
func PathLT(v string) predicate.Daemon {
	return predicate.Daemon(sql.FieldLT(FieldPath, v))
}

// PathLTE applies the LTE predicate on the "path" field.
func PathLTE(v string) predicate.Daemon {
	return predicate.Daemon(sql.FieldLTE(FieldPath, v))
}

// PathContains applies the Contains predicate on the "path" field.
func PathContains(v string) predicate.Daemon {
	return predicate.Daemon(sql.FieldContains(FieldPath, v))
}

// PathHasPrefix applies the HasPrefix predicate on the "path" field.
func PathHasPrefix(v string) predicate.Daemon {
	return predicate.Daemon(sql.FieldHasPrefix(FieldPath, v))
}

// PathHasSuffix applies the HasSuffix predicate on the "path" field.
func PathHasSuffix(v string) predicate.Daemon {
	return predicate.Daemon(sql.FieldHasSuffix(FieldPath, v))
}

// PathEqualFold applies the EqualFold predicate on the "path" field.
func PathEqualFold(v string) predicate.Daemon {
	return predicate.Daemon(sql.FieldEqualFold(FieldPath, v))
}

// PathContainsFold applies the ContainsFold predicate on the "path" field.
func PathContainsFold(v string) predicate.Daemon {
	return predicate.Daemon(sql.FieldContainsFold(FieldPath, v))
}

// EnableEQ applies the EQ predicate on the "enable" field.
func EnableEQ(v bool) predicate.Daemon {
	return predicate.Daemon(sql.FieldEQ(FieldEnable, v))
}

// EnableNEQ applies the NEQ predicate on the "enable" field.
func EnableNEQ(v bool) predicate.Daemon {
	return predicate.Daemon(sql.FieldNEQ(FieldEnable, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Daemon {
	return predicate.Daemon(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Daemon {
	return predicate.Daemon(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Daemon) predicate.Daemon {
	return predicate.Daemon(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Daemon) predicate.Daemon {
	return predicate.Daemon(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Daemon) predicate.Daemon {
	return predicate.Daemon(sql.NotPredicates(p))
}
