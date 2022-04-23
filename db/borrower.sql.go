// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: borrower.sql

package db

import (
	"context"
)

const createBorrower = `-- name: CreateBorrower :one
INSERT INTO borrowers (
  user_id,
  loan_id,
  created_by,
  last_updated_by,
  ip_from,
  user_agent
) VALUES (
  $1, $2,$3,$4,$5,$6
)
RETURNING id, user_id, loan_id, is_active, created_by, created_at, last_updated_by, updated_at, ip_from, user_agent
`

type CreateBorrowerParams struct {
	UserID        int64  `json:"user_id"`
	LoanID        int64  `json:"loan_id"`
	CreatedBy     string `json:"created_by"`
	LastUpdatedBy string `json:"last_updated_by"`
	IpFrom        string `json:"ip_from"`
	UserAgent     string `json:"user_agent"`
}

func (q *Queries) CreateBorrower(ctx context.Context, arg CreateBorrowerParams) (Borrower, error) {
	row := q.db.QueryRowContext(ctx, createBorrower,
		arg.UserID,
		arg.LoanID,
		arg.CreatedBy,
		arg.LastUpdatedBy,
		arg.IpFrom,
		arg.UserAgent,
	)
	var i Borrower
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.LoanID,
		&i.IsActive,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.LastUpdatedBy,
		&i.UpdatedAt,
		&i.IpFrom,
		&i.UserAgent,
	)
	return i, err
}

const deleteBorrower = `-- name: DeleteBorrower :exec
UPDATE borrowers SET 
is_active=false
WHERE id = $1 RETURNING id, user_id, loan_id, is_active, created_by, created_at, last_updated_by, updated_at, ip_from, user_agent
`

func (q *Queries) DeleteBorrower(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteBorrower, id)
	return err
}

const getBorrower = `-- name: GetBorrower :one
SELECT id, user_id, loan_id, is_active, created_by, created_at, last_updated_by, updated_at, ip_from, user_agent FROM borrowers
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetBorrower(ctx context.Context, id int64) (Borrower, error) {
	row := q.db.QueryRowContext(ctx, getBorrower, id)
	var i Borrower
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.LoanID,
		&i.IsActive,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.LastUpdatedBy,
		&i.UpdatedAt,
		&i.IpFrom,
		&i.UserAgent,
	)
	return i, err
}

const getBorrowerByUserIdAndLoanId = `-- name: GetBorrowerByUserIdAndLoanId :one
SELECT id, user_id, loan_id, is_active, created_by, created_at, last_updated_by, updated_at, ip_from, user_agent FROM borrowers
WHERE user_id = $1 AND loan_id=$2  LIMIT 1
`

type GetBorrowerByUserIdAndLoanIdParams struct {
	UserID int64 `json:"user_id"`
	LoanID int64 `json:"loan_id"`
}

func (q *Queries) GetBorrowerByUserIdAndLoanId(ctx context.Context, arg GetBorrowerByUserIdAndLoanIdParams) (Borrower, error) {
	row := q.db.QueryRowContext(ctx, getBorrowerByUserIdAndLoanId, arg.UserID, arg.LoanID)
	var i Borrower
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.LoanID,
		&i.IsActive,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.LastUpdatedBy,
		&i.UpdatedAt,
		&i.IpFrom,
		&i.UserAgent,
	)
	return i, err
}

const listBorrower = `-- name: ListBorrower :many
SELECT id, user_id, loan_id, is_active, created_by, created_at, last_updated_by, updated_at, ip_from, user_agent FROM borrowers 
ORDER BY id 
LIMIT $1 OFFSET $2
`

type ListBorrowerParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListBorrower(ctx context.Context, arg ListBorrowerParams) ([]Borrower, error) {
	rows, err := q.db.QueryContext(ctx, listBorrower, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Borrower{}
	for rows.Next() {
		var i Borrower
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.LoanID,
			&i.IsActive,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.LastUpdatedBy,
			&i.UpdatedAt,
			&i.IpFrom,
			&i.UserAgent,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listDescBorrower = `-- name: ListDescBorrower :many
SELECT id, user_id, loan_id, is_active, created_by, created_at, last_updated_by, updated_at, ip_from, user_agent FROM borrowers 
ORDER BY id  DESC
LIMIT $1 OFFSET $2
`

type ListDescBorrowerParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListDescBorrower(ctx context.Context, arg ListDescBorrowerParams) ([]Borrower, error) {
	rows, err := q.db.QueryContext(ctx, listDescBorrower, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Borrower{}
	for rows.Next() {
		var i Borrower
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.LoanID,
			&i.IsActive,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.LastUpdatedBy,
			&i.UpdatedAt,
			&i.IpFrom,
			&i.UserAgent,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBorrower = `-- name: UpdateBorrower :one
UPDATE borrowers SET 
user_id=$2,
loan_id=$3,
last_updated_by=$4
WHERE id = $1 RETURNING id, user_id, loan_id, is_active, created_by, created_at, last_updated_by, updated_at, ip_from, user_agent
`

type UpdateBorrowerParams struct {
	ID            int64  `json:"id"`
	UserID        int64  `json:"user_id"`
	LoanID        int64  `json:"loan_id"`
	LastUpdatedBy string `json:"last_updated_by"`
}

func (q *Queries) UpdateBorrower(ctx context.Context, arg UpdateBorrowerParams) (Borrower, error) {
	row := q.db.QueryRowContext(ctx, updateBorrower,
		arg.ID,
		arg.UserID,
		arg.LoanID,
		arg.LastUpdatedBy,
	)
	var i Borrower
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.LoanID,
		&i.IsActive,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.LastUpdatedBy,
		&i.UpdatedAt,
		&i.IpFrom,
		&i.UserAgent,
	)
	return i, err
}
