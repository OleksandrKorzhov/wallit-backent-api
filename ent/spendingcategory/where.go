// Code generated by ent, DO NOT EDIT.

package spendingcategory

import (
	"time"
	"wallit/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldEQ(FieldName, v))
}

// CategoryID applies equality check predicate on the "category_id" field. It's identical to CategoryIDEQ.
func CategoryID(v string) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldEQ(FieldCategoryID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldEQ(FieldCreatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldContainsFold(FieldName, v))
}

// CategoryIDEQ applies the EQ predicate on the "category_id" field.
func CategoryIDEQ(v string) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldEQ(FieldCategoryID, v))
}

// CategoryIDNEQ applies the NEQ predicate on the "category_id" field.
func CategoryIDNEQ(v string) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldNEQ(FieldCategoryID, v))
}

// CategoryIDIn applies the In predicate on the "category_id" field.
func CategoryIDIn(vs ...string) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldIn(FieldCategoryID, vs...))
}

// CategoryIDNotIn applies the NotIn predicate on the "category_id" field.
func CategoryIDNotIn(vs ...string) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldNotIn(FieldCategoryID, vs...))
}

// CategoryIDGT applies the GT predicate on the "category_id" field.
func CategoryIDGT(v string) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldGT(FieldCategoryID, v))
}

// CategoryIDGTE applies the GTE predicate on the "category_id" field.
func CategoryIDGTE(v string) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldGTE(FieldCategoryID, v))
}

// CategoryIDLT applies the LT predicate on the "category_id" field.
func CategoryIDLT(v string) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldLT(FieldCategoryID, v))
}

// CategoryIDLTE applies the LTE predicate on the "category_id" field.
func CategoryIDLTE(v string) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldLTE(FieldCategoryID, v))
}

// CategoryIDContains applies the Contains predicate on the "category_id" field.
func CategoryIDContains(v string) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldContains(FieldCategoryID, v))
}

// CategoryIDHasPrefix applies the HasPrefix predicate on the "category_id" field.
func CategoryIDHasPrefix(v string) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldHasPrefix(FieldCategoryID, v))
}

// CategoryIDHasSuffix applies the HasSuffix predicate on the "category_id" field.
func CategoryIDHasSuffix(v string) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldHasSuffix(FieldCategoryID, v))
}

// CategoryIDEqualFold applies the EqualFold predicate on the "category_id" field.
func CategoryIDEqualFold(v string) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldEqualFold(FieldCategoryID, v))
}

// CategoryIDContainsFold applies the ContainsFold predicate on the "category_id" field.
func CategoryIDContainsFold(v string) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldContainsFold(FieldCategoryID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.SpendingCategory {
	return predicate.SpendingCategory(sql.FieldLTE(FieldCreatedAt, v))
}

// HasCategoryTransactions applies the HasEdge predicate on the "category_transactions" edge.
func HasCategoryTransactions() predicate.SpendingCategory {
	return predicate.SpendingCategory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, CategoryTransactionsTable, CategoryTransactionsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCategoryTransactionsWith applies the HasEdge predicate on the "category_transactions" edge with a given conditions (other predicates).
func HasCategoryTransactionsWith(preds ...predicate.Transaction) predicate.SpendingCategory {
	return predicate.SpendingCategory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CategoryTransactionsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, CategoryTransactionsTable, CategoryTransactionsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasInterestedUsers applies the HasEdge predicate on the "interested_users" edge.
func HasInterestedUsers() predicate.SpendingCategory {
	return predicate.SpendingCategory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, InterestedUsersTable, InterestedUsersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInterestedUsersWith applies the HasEdge predicate on the "interested_users" edge with a given conditions (other predicates).
func HasInterestedUsersWith(preds ...predicate.User) predicate.SpendingCategory {
	return predicate.SpendingCategory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InterestedUsersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, InterestedUsersTable, InterestedUsersPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.SpendingCategory) predicate.SpendingCategory {
	return predicate.SpendingCategory(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.SpendingCategory) predicate.SpendingCategory {
	return predicate.SpendingCategory(func(s *sql.Selector) {
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
func Not(p predicate.SpendingCategory) predicate.SpendingCategory {
	return predicate.SpendingCategory(func(s *sql.Selector) {
		p(s.Not())
	})
}