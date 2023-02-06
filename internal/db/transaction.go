package db

import (
	"context"
	"log"
	"wallit/ent"
)

func WithTransaction(ctx context.Context, db *ent.Client, fn func(db *ent.Client) error) error {
	tx, err := db.Tx(ctx)
	if err != nil {
		log.Printf("failed starting transaction: %v", err)
	}

	if err := fn(tx.Client()); err != nil {
		log.Printf("failed executing operation within transaction: %v", err)
		return tx.Rollback()
	} else {
		return tx.Commit()
	}
}
