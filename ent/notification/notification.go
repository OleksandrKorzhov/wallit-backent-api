// Code generated by ent, DO NOT EDIT.

package notification

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

const (
	// Label holds the string label denoting the notification type in the database.
	Label = "notification"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCheckedInApp holds the string denoting the checked_in_app field in the database.
	FieldCheckedInApp = "checked_in_app"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldIncludedInAggregatedOfferAt holds the string denoting the included_in_aggregated_offer_at field in the database.
	FieldIncludedInAggregatedOfferAt = "included_in_aggregated_offer_at"
	// EdgeNotificationRecipient holds the string denoting the notification_recipient edge name in mutations.
	EdgeNotificationRecipient = "notification_recipient"
	// EdgeNotificationDiscountOffer holds the string denoting the notification_discount_offer edge name in mutations.
	EdgeNotificationDiscountOffer = "notification_discount_offer"
	// Table holds the table name of the notification in the database.
	Table = "notifications"
	// NotificationRecipientTable is the table that holds the notification_recipient relation/edge.
	NotificationRecipientTable = "notifications"
	// NotificationRecipientInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	NotificationRecipientInverseTable = "users"
	// NotificationRecipientColumn is the table column denoting the notification_recipient relation/edge.
	NotificationRecipientColumn = "user_notifications"
	// NotificationDiscountOfferTable is the table that holds the notification_discount_offer relation/edge.
	NotificationDiscountOfferTable = "discount_offers"
	// NotificationDiscountOfferInverseTable is the table name for the DiscountOffer entity.
	// It exists in this package in order to avoid circular dependency with the "discountoffer" package.
	NotificationDiscountOfferInverseTable = "discount_offers"
	// NotificationDiscountOfferColumn is the table column denoting the notification_discount_offer relation/edge.
	NotificationDiscountOfferColumn = "notification_notification_discount_offer"
)

// Columns holds all SQL columns for notification fields.
var Columns = []string{
	FieldID,
	FieldCheckedInApp,
	FieldType,
	FieldCreatedAt,
	FieldIncludedInAggregatedOfferAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "notifications"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_notifications",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCheckedInApp holds the default value on creation for the "checked_in_app" field.
	DefaultCheckedInApp bool
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// Type defines the type for the "type" enum field.
type Type string

// TypeInsights is the default value of the Type enum.
const DefaultType = TypeInsights

// Type values.
const (
	TypeOffer    Type = "OFFER"
	TypeInsights Type = "INSIGHTS"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeOffer, TypeInsights:
		return nil
	default:
		return fmt.Errorf("notification: invalid enum value for type field: %q", _type)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Type) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Type) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Type(str)
	if err := TypeValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Type", str)
	}
	return nil
}
