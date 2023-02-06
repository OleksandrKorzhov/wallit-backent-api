// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"wallit/ent/plaiditem"
	"wallit/ent/transactionsync"

	"entgo.io/ent/dialect/sql"
)

// TransactionSync is the model entity for the TransactionSync schema.
type TransactionSync struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Cursor holds the value of the "cursor" field.
	Cursor string `json:"cursor,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TransactionSyncQuery when eager-loading is set.
	Edges                        TransactionSyncEdges `json:"edges"`
	plaid_item_transaction_syncs *int
}

// TransactionSyncEdges holds the relations/edges for other nodes in the graph.
type TransactionSyncEdges struct {
	// Item holds the value of the item edge.
	Item *PlaidItem `json:"item,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// ItemOrErr returns the Item value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransactionSyncEdges) ItemOrErr() (*PlaidItem, error) {
	if e.loadedTypes[0] {
		if e.Item == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: plaiditem.Label}
		}
		return e.Item, nil
	}
	return nil, &NotLoadedError{edge: "item"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TransactionSync) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case transactionsync.FieldID:
			values[i] = new(sql.NullInt64)
		case transactionsync.FieldCursor:
			values[i] = new(sql.NullString)
		case transactionsync.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case transactionsync.ForeignKeys[0]: // plaid_item_transaction_syncs
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TransactionSync", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TransactionSync fields.
func (ts *TransactionSync) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case transactionsync.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ts.ID = int(value.Int64)
		case transactionsync.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ts.CreatedAt = value.Time
			}
		case transactionsync.FieldCursor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cursor", values[i])
			} else if value.Valid {
				ts.Cursor = value.String
			}
		case transactionsync.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field plaid_item_transaction_syncs", value)
			} else if value.Valid {
				ts.plaid_item_transaction_syncs = new(int)
				*ts.plaid_item_transaction_syncs = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryItem queries the "item" edge of the TransactionSync entity.
func (ts *TransactionSync) QueryItem() *PlaidItemQuery {
	return NewTransactionSyncClient(ts.config).QueryItem(ts)
}

// Update returns a builder for updating this TransactionSync.
// Note that you need to call TransactionSync.Unwrap() before calling this method if this TransactionSync
// was returned from a transaction, and the transaction was committed or rolled back.
func (ts *TransactionSync) Update() *TransactionSyncUpdateOne {
	return NewTransactionSyncClient(ts.config).UpdateOne(ts)
}

// Unwrap unwraps the TransactionSync entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ts *TransactionSync) Unwrap() *TransactionSync {
	_tx, ok := ts.config.driver.(*txDriver)
	if !ok {
		panic("ent: TransactionSync is not a transactional entity")
	}
	ts.config.driver = _tx.drv
	return ts
}

// String implements the fmt.Stringer.
func (ts *TransactionSync) String() string {
	var builder strings.Builder
	builder.WriteString("TransactionSync(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ts.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ts.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("cursor=")
	builder.WriteString(ts.Cursor)
	builder.WriteByte(')')
	return builder.String()
}

// TransactionSyncs is a parsable slice of TransactionSync.
type TransactionSyncs []*TransactionSync

func (ts TransactionSyncs) config(cfg config) {
	for _i := range ts {
		ts[_i].config = cfg
	}
}
