// Code generated by ent, DO NOT EDIT.

package discountoffer

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

const (
	// Label holds the string label denoting the discountoffer type in the database.
	Label = "discount_offer"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldCurrency holds the string denoting the currency field in the database.
	FieldCurrency = "currency"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldMerchantSpecificIdentification holds the string denoting the merchant_specific_identification field in the database.
	FieldMerchantSpecificIdentification = "merchant_specific_identification"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldExpiresAt holds the string denoting the expires_at field in the database.
	FieldExpiresAt = "expires_at"
	// EdgeOwnerMerchant holds the string denoting the owner_merchant edge name in mutations.
	EdgeOwnerMerchant = "owner_merchant"
	// EdgeDiscountEligibleUsers holds the string denoting the discount_eligible_users edge name in mutations.
	EdgeDiscountEligibleUsers = "discount_eligible_users"
	// EdgeDiscountOfferNotification holds the string denoting the discount_offer_notification edge name in mutations.
	EdgeDiscountOfferNotification = "discount_offer_notification"
	// Table holds the table name of the discountoffer in the database.
	Table = "discount_offers"
	// OwnerMerchantTable is the table that holds the owner_merchant relation/edge.
	OwnerMerchantTable = "discount_offers"
	// OwnerMerchantInverseTable is the table name for the Merchant entity.
	// It exists in this package in order to avoid circular dependency with the "merchant" package.
	OwnerMerchantInverseTable = "merchants"
	// OwnerMerchantColumn is the table column denoting the owner_merchant relation/edge.
	OwnerMerchantColumn = "merchant_discount_offers"
	// DiscountEligibleUsersTable is the table that holds the discount_eligible_users relation/edge. The primary key declared below.
	DiscountEligibleUsersTable = "user_available_discount_offers"
	// DiscountEligibleUsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	DiscountEligibleUsersInverseTable = "users"
	// DiscountOfferNotificationTable is the table that holds the discount_offer_notification relation/edge.
	DiscountOfferNotificationTable = "discount_offers"
	// DiscountOfferNotificationInverseTable is the table name for the Notification entity.
	// It exists in this package in order to avoid circular dependency with the "notification" package.
	DiscountOfferNotificationInverseTable = "notifications"
	// DiscountOfferNotificationColumn is the table column denoting the discount_offer_notification relation/edge.
	DiscountOfferNotificationColumn = "notification_notification_discount_offer"
)

// Columns holds all SQL columns for discountoffer fields.
var Columns = []string{
	FieldID,
	FieldAmount,
	FieldType,
	FieldCurrency,
	FieldDescription,
	FieldMerchantSpecificIdentification,
	FieldCreatedAt,
	FieldExpiresAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "discount_offers"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"merchant_discount_offers",
	"notification_notification_discount_offer",
}

var (
	// DiscountEligibleUsersPrimaryKey and DiscountEligibleUsersColumn2 are the table columns denoting the
	// primary key for the discount_eligible_users relation (M2M).
	DiscountEligibleUsersPrimaryKey = []string{"user_id", "discount_offer_id"}
)

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
	// AmountValidator is a validator for the "amount" field. It is called by the builders before save.
	AmountValidator func(int) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// MerchantSpecificIdentificationValidator is a validator for the "merchant_specific_identification" field. It is called by the builders before save.
	MerchantSpecificIdentificationValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// Type defines the type for the "type" enum field.
type Type string

// TypePercentFromPrice is the default value of the Type enum.
const DefaultType = TypePercentFromPrice

// Type values.
const (
	TypeFixedAmount      Type = "FIXED_AMOUNT"
	TypePercentFromPrice Type = "PERCENT_FROM_PRICE"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeFixedAmount, TypePercentFromPrice:
		return nil
	default:
		return fmt.Errorf("discountoffer: invalid enum value for type field: %q", _type)
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
