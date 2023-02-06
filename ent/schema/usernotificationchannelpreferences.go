package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserNotificationChannelPreferences holds the schema definition for the UserNotificationChannelPreferences entity.
type UserNotificationChannelPreferences struct {
	ent.Schema
}

// Fields of the UserNotificationChannelPreferences.
func (UserNotificationChannelPreferences) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("chanel").NamedValues(
			"Email", "EMAIL",
			"Sms", "SMS",
			"Push", "PUSH",
		),
	}
}

// Edges of the UserNotificationChannelPreferences.
func (UserNotificationChannelPreferences) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("chanel_users", User.Type).
			Unique().
			Ref("notification_channels").
			Annotations(
				entgql.Skip(entgql.SkipMutationUpdateInput),
			),
	}
}

func (UserNotificationChannelPreferences) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationUpdate()),
	}
}
