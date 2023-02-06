package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Transaction holds the schema definition for the Transaction entity.
type Transaction struct {
	ent.Schema
}

// Fields of the Transaction.
func (Transaction) Fields() []ent.Field {
	return []ent.Field{
		field.String("financial_account_id").NotEmpty(),
		field.Float("amount"),
		field.String("iso_currency_code").NotEmpty(),
		field.String("unofficial_currency_code").Nillable(),
		//field.String("category").NotEmpty(), // CSV of categories
		field.String("category").Optional(), // CSV of categories
		//field.String("category_id").NotEmpty(),
		field.String("category_id").Optional(),
		field.String("check_number").Nillable(),
		field.String("date").
			NotEmpty().
			Annotations(
				entgql.OrderField("DATE"),
			),
		field.Time("datetime").
			Optional().
			Annotations(
				entgql.OrderField("DATETIME"),
			),
		field.String("authorized_date").Optional(),
		field.Time("authorized_datetime").Optional(),
		field.String("location_address"),
		field.String("location_city"),
		field.String("location_region"),
		field.String("location_postal_code"),
		field.Float("location_lat"),
		field.Float("location_lon"),
		field.String("location_store_number"),
		field.String("name"),
		field.String("merchant_name"),
		field.String("payment_channel"),
		field.Bool("pending"),
		field.String("pending_transaction_id"),
		field.String("account_owner"),
		field.String("transaction_id"),
		field.String("transaction_code"),
		field.Time("created_at").Default(time.Now),
		// @TODO: think of adding relation to the item that was used to fetch transaction data
	}
}

// Edges of the Transaction.
func (Transaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("institution_account", PlaidInstitutionAccount.Type).
			Ref("transactions").
			Unique(),
		edge.To("transaction_categories", SpendingCategory.Type),
	}
}

func (Transaction) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
	}
}
