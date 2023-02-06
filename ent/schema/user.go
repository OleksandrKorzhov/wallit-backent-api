package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("identity_provider_id").
			Optional().
			Annotations(
				entgql.Skip(entgql.SkipMutationUpdateInput),
			),
		field.Enum("offer_frequency").
			NamedValues(
				"NoOffers", "NO_OFFERS",
				"AssSoonAsPossible", "ASS_SOON_AS_POSSIBLE",
				"Daily", "DAILY",
				"Weekly", "WEEKLY",
				"BiWeekly", "BI_WEEKLY",
				"Monthly", "MONTHLY",
			).
			Default("NO_OFFERS"),
		field.String("home_country").Optional(),
		field.String("home_state").Optional(),
		field.String("home_city").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("plaid_items", PlaidItem.Type).
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
			),

		edge.To("spending_categories", SpendingCategory.Type).
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
			),

		edge.To("notification_channels", UserNotificationChannelPreferences.Type).
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
			),

		edge.To("notifications", Notification.Type).
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
			),

		edge.To("available_discount_offers", DiscountOffer.Type).
			Annotations(
				entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
			),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}
