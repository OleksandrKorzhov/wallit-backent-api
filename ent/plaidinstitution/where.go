// Code generated by ent, DO NOT EDIT.

package plaidinstitution

import (
	"wallit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldLTE(FieldID, id))
}

// InstitutionID applies equality check predicate on the "institution_id" field. It's identical to InstitutionIDEQ.
func InstitutionID(v string) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldEQ(FieldInstitutionID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldEQ(FieldName, v))
}

// InstitutionIDEQ applies the EQ predicate on the "institution_id" field.
func InstitutionIDEQ(v string) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldEQ(FieldInstitutionID, v))
}

// InstitutionIDNEQ applies the NEQ predicate on the "institution_id" field.
func InstitutionIDNEQ(v string) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldNEQ(FieldInstitutionID, v))
}

// InstitutionIDIn applies the In predicate on the "institution_id" field.
func InstitutionIDIn(vs ...string) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldIn(FieldInstitutionID, vs...))
}

// InstitutionIDNotIn applies the NotIn predicate on the "institution_id" field.
func InstitutionIDNotIn(vs ...string) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldNotIn(FieldInstitutionID, vs...))
}

// InstitutionIDGT applies the GT predicate on the "institution_id" field.
func InstitutionIDGT(v string) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldGT(FieldInstitutionID, v))
}

// InstitutionIDGTE applies the GTE predicate on the "institution_id" field.
func InstitutionIDGTE(v string) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldGTE(FieldInstitutionID, v))
}

// InstitutionIDLT applies the LT predicate on the "institution_id" field.
func InstitutionIDLT(v string) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldLT(FieldInstitutionID, v))
}

// InstitutionIDLTE applies the LTE predicate on the "institution_id" field.
func InstitutionIDLTE(v string) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldLTE(FieldInstitutionID, v))
}

// InstitutionIDContains applies the Contains predicate on the "institution_id" field.
func InstitutionIDContains(v string) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldContains(FieldInstitutionID, v))
}

// InstitutionIDHasPrefix applies the HasPrefix predicate on the "institution_id" field.
func InstitutionIDHasPrefix(v string) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldHasPrefix(FieldInstitutionID, v))
}

// InstitutionIDHasSuffix applies the HasSuffix predicate on the "institution_id" field.
func InstitutionIDHasSuffix(v string) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldHasSuffix(FieldInstitutionID, v))
}

// InstitutionIDEqualFold applies the EqualFold predicate on the "institution_id" field.
func InstitutionIDEqualFold(v string) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldEqualFold(FieldInstitutionID, v))
}

// InstitutionIDContainsFold applies the ContainsFold predicate on the "institution_id" field.
func InstitutionIDContainsFold(v string) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldContainsFold(FieldInstitutionID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(sql.FieldContainsFold(FieldName, v))
}

// HasPlaidItem applies the HasEdge predicate on the "plaid_item" edge.
func HasPlaidItem() predicate.PlaidInstitution {
	return predicate.PlaidInstitution(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, PlaidItemTable, PlaidItemColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlaidItemWith applies the HasEdge predicate on the "plaid_item" edge with a given conditions (other predicates).
func HasPlaidItemWith(preds ...predicate.PlaidItem) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlaidItemInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, PlaidItemTable, PlaidItemColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAccounts applies the HasEdge predicate on the "accounts" edge.
func HasAccounts() predicate.PlaidInstitution {
	return predicate.PlaidInstitution(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AccountsTable, AccountsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAccountsWith applies the HasEdge predicate on the "accounts" edge with a given conditions (other predicates).
func HasAccountsWith(preds ...predicate.PlaidInstitutionAccount) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AccountsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AccountsTable, AccountsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PlaidInstitution) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PlaidInstitution) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PlaidInstitution) predicate.PlaidInstitution {
	return predicate.PlaidInstitution(func(s *sql.Selector) {
		p(s.Not())
	})
}
