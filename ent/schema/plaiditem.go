package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// PlaidItem holds the schema definition for the PlaidItem entity.
type PlaidItem struct {
	ent.Schema
}

// Fields of the PlaidItem.
func (PlaidItem) Fields() []ent.Field {
	// @TODO: add unique constraint
	return []ent.Field{
		field.String("item_id").NotEmpty(),
		field.String("access_token").NotEmpty(),
		field.Time("created_at").
			Default(time.Now).
			Annotations(
				entgql.OrderField("CREATED_AT"),
			),
		field.Time("updated_at").
			Default(time.Now).
			Annotations(
				entgql.OrderField("UPDATED_AT"),
			),
	}
}

// Edges of the PlaidItem.
func (PlaidItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("plaid_items").
			Unique(),
		edge.To("transaction_syncs", TransactionSync.Type),
		edge.To("institution", PlaidInstitution.Type).Unique(),
	}
}

func (PlaidItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
	}
}
