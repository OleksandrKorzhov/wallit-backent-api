// Code generated by ent, DO NOT EDIT.

package spendingcategory

import (
	"time"
)

const (
	// Label holds the string label denoting the spendingcategory type in the database.
	Label = "spending_category"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCategoryID holds the string denoting the category_id field in the database.
	FieldCategoryID = "category_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeCategoryTransactions holds the string denoting the category_transactions edge name in mutations.
	EdgeCategoryTransactions = "category_transactions"
	// EdgeInterestedUsers holds the string denoting the interested_users edge name in mutations.
	EdgeInterestedUsers = "interested_users"
	// Table holds the table name of the spendingcategory in the database.
	Table = "spending_categories"
	// CategoryTransactionsTable is the table that holds the category_transactions relation/edge. The primary key declared below.
	CategoryTransactionsTable = "transaction_transaction_categories"
	// CategoryTransactionsInverseTable is the table name for the Transaction entity.
	// It exists in this package in order to avoid circular dependency with the "transaction" package.
	CategoryTransactionsInverseTable = "transactions"
	// InterestedUsersTable is the table that holds the interested_users relation/edge. The primary key declared below.
	InterestedUsersTable = "user_spending_categories"
	// InterestedUsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	InterestedUsersInverseTable = "users"
)

// Columns holds all SQL columns for spendingcategory fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCategoryID,
	FieldCreatedAt,
}

var (
	// CategoryTransactionsPrimaryKey and CategoryTransactionsColumn2 are the table columns denoting the
	// primary key for the category_transactions relation (M2M).
	CategoryTransactionsPrimaryKey = []string{"transaction_id", "spending_category_id"}
	// InterestedUsersPrimaryKey and InterestedUsersColumn2 are the table columns denoting the
	// primary key for the interested_users relation (M2M).
	InterestedUsersPrimaryKey = []string{"user_id", "spending_category_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)