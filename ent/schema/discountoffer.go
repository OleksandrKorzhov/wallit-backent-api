package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// DiscountOffer holds the schema definition for the DiscountOffer entity.
type DiscountOffer struct {
	ent.Schema
}

// @TODO: the problem with such implementation is that if a new user is registered and is formally eligible to get the discount they will not be included
// In can be a next step to implement a kind of rules engine to automatically and in real time select users for the discount

// @TODO: also, in the current implementation it is hard to track what exactly users consumed discount

// Fields of the DiscountOffer.
func (DiscountOffer) Fields() []ent.Field {
	return []ent.Field{
		field.Int("amount").NonNegative(),

		field.Enum("type").
			NamedValues(
				"FixedAmount", "FIXED_AMOUNT",
				"PercentFromPrice", "PERCENT_FROM_PRICE",
			).
			Default("PERCENT_FROM_PRICE"),

		field.String("currency").Optional(),

		field.String("description").NotEmpty(),

		field.String("merchant_specific_identification").
			NotEmpty(),

		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Annotations(
				entgql.OrderField("CREATED_AT"),
				entgql.Skip(entgql.SkipMutationCreateInput|entgql.SkipMutationUpdateInput),
			),

		field.Time("expires_at"),
	}
}

// Edges of the DiscountOffer.
func (DiscountOffer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner_merchant", Merchant.Type).
			Ref("discount_offers").
			Required().
			Unique(),

		edge.From("discount_eligible_users", User.Type).
			Ref("available_discount_offers").
			Required(),

		edge.From("discount_offer_notification", Notification.Type).
			Unique().
			Ref("notification_discount_offer"),
	}
}

func (DiscountOffer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.Mutations(
			entgql.MutationCreate(),
		),
	}
}
