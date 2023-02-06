package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PlaidInstitution holds the schema definition for the PlaidInstitution entity.
type PlaidInstitution struct {
	ent.Schema
}

// Fields of the PlaidInstitution.
func (PlaidInstitution) Fields() []ent.Field {
	return []ent.Field{
		field.String("institution_id"),
		field.String("name"),
	}
}

// Edges of the PlaidInstitution.
func (PlaidInstitution) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("plaid_item", PlaidItem.Type).
			Unique().
			Ref("institution"),
		edge.To("accounts", PlaidInstitutionAccount.Type),
	}
}

func (PlaidInstitution) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
	}
}
