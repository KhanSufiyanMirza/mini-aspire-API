package db

import (
	"context"
	"database/sql"
	"fmt"
)

//Store provide all funtions to execute db queries and transactions
type Store interface {
	Querier
	CreateLoanWithBorrower(ctx context.Context, arg CreateLoanParams) (Loan, error)
	CreatePaymentTerms(ctx context.Context, arg CreatePaymentParams) (TransactionDetail, error)
}

//SQLStore provide all funtions to execute SQL queries and transactions
type SQLStore struct {
	db *sql.DB
	*Queries
}

//NewStore Creates a new store
func NewStore(db *sql.DB) Store {

	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// ExecTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
