package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Notification holds the schema definition for the Notification entity.
type Notification struct {
	ent.Schema
}

// @TODO: it would be niсe to get the user timexone to know when to send notification to them

// Fields of the Notification.
func (Notification) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("checked_in_app").
			Default(false),

		field.Enum("type").
			NamedValues(
				"Offer", "OFFER",
				"Insights", "INSIGHTS",
			).
			Default("INSIGHTS"),

		field.Time("created_at").
			Immutable().
			Default(time.Now).
			Annotations(
				entgql.OrderField("CREATED_AT"),
			),

		// @TODO: in case more information must be captured the aggregated offer can be represented as a separate table
		field.Time("included_in_aggregated_offer_at").
			Immutable().
			Optional(),
		// @TODO: the content and logic of the aggregated report must be prototyped separately
	}
}

// Edges of the Notification.
func (Notification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("notification_recipient", User.Type).
			Required().
			Unique().
			Ref("notifications"),

		edge.To("notification_discount_offer", DiscountOffer.Type).
			Unique(),
	}
}

func (Notification) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
	}
}
