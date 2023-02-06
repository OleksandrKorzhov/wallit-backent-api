// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
	"wallit/ent/discountoffer"
	"wallit/ent/merchant"
	"wallit/ent/notification"
	"wallit/ent/plaiditem"
	"wallit/ent/schema"
	"wallit/ent/spendingcategory"
	"wallit/ent/transaction"
	"wallit/ent/transactionsync"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	discountofferFields := schema.DiscountOffer{}.Fields()
	_ = discountofferFields
	// discountofferDescAmount is the schema descriptor for amount field.
	discountofferDescAmount := discountofferFields[0].Descriptor()
	// discountoffer.AmountValidator is a validator for the "amount" field. It is called by the builders before save.
	discountoffer.AmountValidator = discountofferDescAmount.Validators[0].(func(int) error)
	// discountofferDescDescription is the schema descriptor for description field.
	discountofferDescDescription := discountofferFields[3].Descriptor()
	// discountoffer.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	discountoffer.DescriptionValidator = discountofferDescDescription.Validators[0].(func(string) error)
	// discountofferDescMerchantSpecificIdentification is the schema descriptor for merchant_specific_identification field.
	discountofferDescMerchantSpecificIdentification := discountofferFields[4].Descriptor()
	// discountoffer.MerchantSpecificIdentificationValidator is a validator for the "merchant_specific_identification" field. It is called by the builders before save.
	discountoffer.MerchantSpecificIdentificationValidator = discountofferDescMerchantSpecificIdentification.Validators[0].(func(string) error)
	// discountofferDescCreatedAt is the schema descriptor for created_at field.
	discountofferDescCreatedAt := discountofferFields[5].Descriptor()
	// discountoffer.DefaultCreatedAt holds the default value on creation for the created_at field.
	discountoffer.DefaultCreatedAt = discountofferDescCreatedAt.Default.(func() time.Time)
	merchantFields := schema.Merchant{}.Fields()
	_ = merchantFields
	// merchantDescName is the schema descriptor for name field.
	merchantDescName := merchantFields[0].Descriptor()
	// merchant.NameValidator is a validator for the "name" field. It is called by the builders before save.
	merchant.NameValidator = merchantDescName.Validators[0].(func(string) error)
	notificationFields := schema.Notification{}.Fields()
	_ = notificationFields
	// notificationDescCheckedInApp is the schema descriptor for checked_in_app field.
	notificationDescCheckedInApp := notificationFields[0].Descriptor()
	// notification.DefaultCheckedInApp holds the default value on creation for the checked_in_app field.
	notification.DefaultCheckedInApp = notificationDescCheckedInApp.Default.(bool)
	// notificationDescCreatedAt is the schema descriptor for created_at field.
	notificationDescCreatedAt := notificationFields[2].Descriptor()
	// notification.DefaultCreatedAt holds the default value on creation for the created_at field.
	notification.DefaultCreatedAt = notificationDescCreatedAt.Default.(func() time.Time)
	plaiditemFields := schema.PlaidItem{}.Fields()
	_ = plaiditemFields
	// plaiditemDescItemID is the schema descriptor for item_id field.
	plaiditemDescItemID := plaiditemFields[0].Descriptor()
	// plaiditem.ItemIDValidator is a validator for the "item_id" field. It is called by the builders before save.
	plaiditem.ItemIDValidator = plaiditemDescItemID.Validators[0].(func(string) error)
	// plaiditemDescAccessToken is the schema descriptor for access_token field.
	plaiditemDescAccessToken := plaiditemFields[1].Descriptor()
	// plaiditem.AccessTokenValidator is a validator for the "access_token" field. It is called by the builders before save.
	plaiditem.AccessTokenValidator = plaiditemDescAccessToken.Validators[0].(func(string) error)
	// plaiditemDescCreatedAt is the schema descriptor for created_at field.
	plaiditemDescCreatedAt := plaiditemFields[2].Descriptor()
	// plaiditem.DefaultCreatedAt holds the default value on creation for the created_at field.
	plaiditem.DefaultCreatedAt = plaiditemDescCreatedAt.Default.(func() time.Time)
	// plaiditemDescUpdatedAt is the schema descriptor for updated_at field.
	plaiditemDescUpdatedAt := plaiditemFields[3].Descriptor()
	// plaiditem.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	plaiditem.DefaultUpdatedAt = plaiditemDescUpdatedAt.Default.(func() time.Time)
	spendingcategoryFields := schema.SpendingCategory{}.Fields()
	_ = spendingcategoryFields
	// spendingcategoryDescCreatedAt is the schema descriptor for created_at field.
	spendingcategoryDescCreatedAt := spendingcategoryFields[2].Descriptor()
	// spendingcategory.DefaultCreatedAt holds the default value on creation for the created_at field.
	spendingcategory.DefaultCreatedAt = spendingcategoryDescCreatedAt.Default.(func() time.Time)
	transactionFields := schema.Transaction{}.Fields()
	_ = transactionFields
	// transactionDescFinancialAccountID is the schema descriptor for financial_account_id field.
	transactionDescFinancialAccountID := transactionFields[0].Descriptor()
	// transaction.FinancialAccountIDValidator is a validator for the "financial_account_id" field. It is called by the builders before save.
	transaction.FinancialAccountIDValidator = transactionDescFinancialAccountID.Validators[0].(func(string) error)
	// transactionDescIsoCurrencyCode is the schema descriptor for iso_currency_code field.
	transactionDescIsoCurrencyCode := transactionFields[2].Descriptor()
	// transaction.IsoCurrencyCodeValidator is a validator for the "iso_currency_code" field. It is called by the builders before save.
	transaction.IsoCurrencyCodeValidator = transactionDescIsoCurrencyCode.Validators[0].(func(string) error)
	// transactionDescDate is the schema descriptor for date field.
	transactionDescDate := transactionFields[7].Descriptor()
	// transaction.DateValidator is a validator for the "date" field. It is called by the builders before save.
	transaction.DateValidator = transactionDescDate.Validators[0].(func(string) error)
	// transactionDescCreatedAt is the schema descriptor for created_at field.
	transactionDescCreatedAt := transactionFields[26].Descriptor()
	// transaction.DefaultCreatedAt holds the default value on creation for the created_at field.
	transaction.DefaultCreatedAt = transactionDescCreatedAt.Default.(func() time.Time)
	transactionsyncFields := schema.TransactionSync{}.Fields()
	_ = transactionsyncFields
	// transactionsyncDescCreatedAt is the schema descriptor for created_at field.
	transactionsyncDescCreatedAt := transactionsyncFields[0].Descriptor()
	// transactionsync.DefaultCreatedAt holds the default value on creation for the created_at field.
	transactionsync.DefaultCreatedAt = transactionsyncDescCreatedAt.Default.(func() time.Time)
	// transactionsyncDescCursor is the schema descriptor for cursor field.
	transactionsyncDescCursor := transactionsyncFields[1].Descriptor()
	// transactionsync.CursorValidator is a validator for the "cursor" field. It is called by the builders before save.
	transactionsync.CursorValidator = transactionsyncDescCursor.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
}
