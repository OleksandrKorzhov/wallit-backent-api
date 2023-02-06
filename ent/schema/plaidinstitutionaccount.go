package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PlaidInstitutionAccount holds the schema definition for the PlaidInstitutionAccount entity.
type PlaidInstitutionAccount struct {
	ent.Schema
}

// Fields of the PlaidInstitutionAccount.
func (PlaidInstitutionAccount) Fields() []ent.Field {
	return []ent.Field{
		field.String("account_id"),
		field.Float("balance_available"),
		field.Float("balance_current"),
		field.String("balance_iso_currency_code"),
		field.String("mask"),
		field.String("name"),
		field.String("official_name").Optional(),
		field.String("type"),
	}
}

// Edges of the PlaidInstitutionAccount.
func (PlaidInstitutionAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("parent_institution", PlaidInstitution.Type).
			Unique().
			Ref("accounts"),
		edge.To("transactions", Transaction.Type).
			Annotations(
				entgql.RelayConnection(),
			),
	}
}

func (PlaidInstitutionAccount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
	}
}
