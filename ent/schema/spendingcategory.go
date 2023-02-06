package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// SpendingCategory holds the schema definition for the SpendingCategory entity.
type SpendingCategory struct {
	ent.Schema
}

// Fields of the SpendingCategory.
func (SpendingCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("category_id"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the SpendingCategory.
func (SpendingCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("category_transactions", Transaction.Type).
			Ref("transaction_categories"),
		edge.From("interested_users", User.Type).
			Ref("spending_categories"),
	}
}

func (SpendingCategory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
	}
}
