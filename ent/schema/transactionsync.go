package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// TransactionSync holds the schema definition for the TransactionSync entity.
type TransactionSync struct {
	ent.Schema
}

// Fields of the TransactionSync.
func (TransactionSync) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now),
		field.String("cursor").NotEmpty(),
	}
}

// Edges of the TransactionSync.
func (TransactionSync) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("item", PlaidItem.Type).
			Ref("transaction_syncs").
			Unique(),
	}
}
