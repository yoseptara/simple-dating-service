package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store defines all function to execute db queries and transaction
type Store interface {
	Querier
	ExecTx(ctx context.Context, fn func(*Queries) error) error
}

// Store provides all functions to execute SQL queries and transactions
type SQLStore struct {
	db *sql.DB
	*Queries
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SQLStore) ExecTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbError := tx.Rollback(); rbError != nil {
			return fmt.Errorf("tx err: %v, rb err : %v", err, rbError)
		}
		return err
	}

	return tx.Commit()
}
