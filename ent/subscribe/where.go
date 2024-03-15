// Code generated by ent, DO NOT EDIT.

package subscribe

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/hvturingga/ya/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldEQ(FieldName, v))
}

// Link applies equality check predicate on the "link" field. It's identical to LinkEQ.
func Link(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldEQ(FieldLink, v))
}

// Conf applies equality check predicate on the "conf" field. It's identical to ConfEQ.
func Conf(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldEQ(FieldConf, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldContainsFold(FieldName, v))
}

// LinkEQ applies the EQ predicate on the "link" field.
func LinkEQ(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldEQ(FieldLink, v))
}

// LinkNEQ applies the NEQ predicate on the "link" field.
func LinkNEQ(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldNEQ(FieldLink, v))
}

// LinkIn applies the In predicate on the "link" field.
func LinkIn(vs ...string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldIn(FieldLink, vs...))
}

// LinkNotIn applies the NotIn predicate on the "link" field.
func LinkNotIn(vs ...string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldNotIn(FieldLink, vs...))
}

// LinkGT applies the GT predicate on the "link" field.
func LinkGT(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldGT(FieldLink, v))
}

// LinkGTE applies the GTE predicate on the "link" field.
func LinkGTE(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldGTE(FieldLink, v))
}

// LinkLT applies the LT predicate on the "link" field.
func LinkLT(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldLT(FieldLink, v))
}

// LinkLTE applies the LTE predicate on the "link" field.
func LinkLTE(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldLTE(FieldLink, v))
}

// LinkContains applies the Contains predicate on the "link" field.
func LinkContains(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldContains(FieldLink, v))
}

// LinkHasPrefix applies the HasPrefix predicate on the "link" field.
func LinkHasPrefix(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldHasPrefix(FieldLink, v))
}

// LinkHasSuffix applies the HasSuffix predicate on the "link" field.
func LinkHasSuffix(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldHasSuffix(FieldLink, v))
}

// LinkIsNil applies the IsNil predicate on the "link" field.
func LinkIsNil() predicate.Subscribe {
	return predicate.Subscribe(sql.FieldIsNull(FieldLink))
}

// LinkNotNil applies the NotNil predicate on the "link" field.
func LinkNotNil() predicate.Subscribe {
	return predicate.Subscribe(sql.FieldNotNull(FieldLink))
}

// LinkEqualFold applies the EqualFold predicate on the "link" field.
func LinkEqualFold(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldEqualFold(FieldLink, v))
}

// LinkContainsFold applies the ContainsFold predicate on the "link" field.
func LinkContainsFold(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldContainsFold(FieldLink, v))
}

// ConfEQ applies the EQ predicate on the "conf" field.
func ConfEQ(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldEQ(FieldConf, v))
}

// ConfNEQ applies the NEQ predicate on the "conf" field.
func ConfNEQ(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldNEQ(FieldConf, v))
}

// ConfIn applies the In predicate on the "conf" field.
func ConfIn(vs ...string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldIn(FieldConf, vs...))
}

// ConfNotIn applies the NotIn predicate on the "conf" field.
func ConfNotIn(vs ...string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldNotIn(FieldConf, vs...))
}

// ConfGT applies the GT predicate on the "conf" field.
func ConfGT(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldGT(FieldConf, v))
}

// ConfGTE applies the GTE predicate on the "conf" field.
func ConfGTE(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldGTE(FieldConf, v))
}

// ConfLT applies the LT predicate on the "conf" field.
func ConfLT(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldLT(FieldConf, v))
}

// ConfLTE applies the LTE predicate on the "conf" field.
func ConfLTE(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldLTE(FieldConf, v))
}

// ConfContains applies the Contains predicate on the "conf" field.
func ConfContains(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldContains(FieldConf, v))
}

// ConfHasPrefix applies the HasPrefix predicate on the "conf" field.
func ConfHasPrefix(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldHasPrefix(FieldConf, v))
}

// ConfHasSuffix applies the HasSuffix predicate on the "conf" field.
func ConfHasSuffix(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldHasSuffix(FieldConf, v))
}

// ConfEqualFold applies the EqualFold predicate on the "conf" field.
func ConfEqualFold(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldEqualFold(FieldConf, v))
}

// ConfContainsFold applies the ContainsFold predicate on the "conf" field.
func ConfContainsFold(v string) predicate.Subscribe {
	return predicate.Subscribe(sql.FieldContainsFold(FieldConf, v))
}

// HasProvider applies the HasEdge predicate on the "provider" edge.
func HasProvider() predicate.Subscribe {
	return predicate.Subscribe(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProviderTable, ProviderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProviderWith applies the HasEdge predicate on the "provider" edge with a given conditions (other predicates).
func HasProviderWith(preds ...predicate.Provider) predicate.Subscribe {
	return predicate.Subscribe(func(s *sql.Selector) {
		step := newProviderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Subscribe {
	return predicate.Subscribe(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Subscribe {
	return predicate.Subscribe(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNodes applies the HasEdge predicate on the "nodes" edge.
func HasNodes() predicate.Subscribe {
	return predicate.Subscribe(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, NodesTable, NodesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNodesWith applies the HasEdge predicate on the "nodes" edge with a given conditions (other predicates).
func HasNodesWith(preds ...predicate.Node) predicate.Subscribe {
	return predicate.Subscribe(func(s *sql.Selector) {
		step := newNodesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Subscribe) predicate.Subscribe {
	return predicate.Subscribe(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Subscribe) predicate.Subscribe {
	return predicate.Subscribe(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Subscribe) predicate.Subscribe {
	return predicate.Subscribe(sql.NotPredicates(p))
}
