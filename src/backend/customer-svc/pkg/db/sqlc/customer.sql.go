// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: customer.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createCustomer = `-- name: CreateCustomer :one
INSERT INTO
  customer.customer (
    company_id,
    first_name,
    last_name,
    email,
    phone,
    identification_number,
    identification_type
  )
VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING id, company_id, first_name, last_name, email, phone, identification_number, identification_type, created_at, modified_at
`

type CreateCustomerParams struct {
	CompanyID            int64          `json:"company_id"`
	FirstName            string         `json:"first_name"`
	LastName             sql.NullString `json:"last_name"`
	Email                sql.NullString `json:"email"`
	Phone                sql.NullString `json:"phone"`
	IdentificationNumber string         `json:"identification_number"`
	IdentificationType   string         `json:"identification_type"`
}

func (q *Queries) CreateCustomer(ctx context.Context, arg CreateCustomerParams) (CustomerCustomer, error) {
	row := q.db.QueryRowContext(ctx, createCustomer,
		arg.CompanyID,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Phone,
		arg.IdentificationNumber,
		arg.IdentificationType,
	)
	var i CustomerCustomer
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Phone,
		&i.IdentificationNumber,
		&i.IdentificationType,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const deleteCustomer = `-- name: DeleteCustomer :exec
DELETE FROM
  customer.customer
WHERE
  id = $1
`

func (q *Queries) DeleteCustomer(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCustomer, id)
	return err
}

const getCustomer = `-- name: GetCustomer :one
SELECT
  id,
  company_id,
  first_name,
  last_name,
  email,
  phone,
  identification_number,
  identification_type,
  created_at,
  modified_at
FROM
  customer.customer
WHERE
  id = $1
LIMIT 1
`

func (q *Queries) GetCustomer(ctx context.Context, id int64) (CustomerCustomer, error) {
	row := q.db.QueryRowContext(ctx, getCustomer, id)
	var i CustomerCustomer
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Phone,
		&i.IdentificationNumber,
		&i.IdentificationType,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const getCustomerByIdentification = `-- name: GetCustomerByIdentification :one
SELECT
  id,
  company_id,
  first_name,
  last_name,
  email,
  phone,
  identification_number,
  identification_type,
  created_at,
  modified_at
FROM
  customer.customer
WHERE
  identification_number = $1 AND identification_type = $2
LIMIT 1
`

type GetCustomerByIdentificationParams struct {
	IdentificationNumber string `json:"identification_number"`
	IdentificationType   string `json:"identification_type"`
}

func (q *Queries) GetCustomerByIdentification(ctx context.Context, arg GetCustomerByIdentificationParams) (CustomerCustomer, error) {
	row := q.db.QueryRowContext(ctx, getCustomerByIdentification, arg.IdentificationNumber, arg.IdentificationType)
	var i CustomerCustomer
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Phone,
		&i.IdentificationNumber,
		&i.IdentificationType,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const getCustomers = `-- name: GetCustomers :many
SELECT
  id,
  company_id,
  first_name,
  last_name,
  email,
  phone,
  identification_number,
  identification_type,
  created_at,
  modified_at
FROM
  customer.customer
WHERE
  ($3 = 0 OR company_id = $3) AND
  ($4::text = '' OR first_name LIKE CONCAT('%', $4::text, '%')) AND
  ($5::text = '' OR last_name LIKE CONCAT('%', $5::text, '%')) AND
  ($6::text = '' OR identification_number LIKE CONCAT('%', $6::text, '%'))
ORDER BY
  id
DESC
LIMIT
  $1
OFFSET
  $2
`

type GetCustomersParams struct {
	Limit                int32       `json:"limit"`
	Offset               int32       `json:"offset"`
	CompanyID            interface{} `json:"company_id"`
	FirstName            string      `json:"first_name"`
	LastName             string      `json:"last_name"`
	IdentificationNumber string      `json:"identification_number"`
}

func (q *Queries) GetCustomers(ctx context.Context, arg GetCustomersParams) ([]CustomerCustomer, error) {
	rows, err := q.db.QueryContext(ctx, getCustomers,
		arg.Limit,
		arg.Offset,
		arg.CompanyID,
		arg.FirstName,
		arg.LastName,
		arg.IdentificationNumber,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CustomerCustomer{}
	for rows.Next() {
		var i CustomerCustomer
		if err := rows.Scan(
			&i.ID,
			&i.CompanyID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Phone,
			&i.IdentificationNumber,
			&i.IdentificationType,
			&i.CreatedAt,
			&i.ModifiedAt,
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

const getCustomersByMonths = `-- name: GetCustomersByMonths :many
SELECT
    created_at
FROM
    customer.customer
WHERE
  ($1 = 0 OR company_id = $1) AND
  (created_at >  CURRENT_DATE - ($2::int * INTERVAL '1 MONTH'))
ORDER BY
  created_at
`

type GetCustomersByMonthsParams struct {
	CompanyID interface{} `json:"company_id"`
	Months    int32       `json:"months"`
}

func (q *Queries) GetCustomersByMonths(ctx context.Context, arg GetCustomersByMonthsParams) ([]time.Time, error) {
	rows, err := q.db.QueryContext(ctx, getCustomersByMonths, arg.CompanyID, arg.Months)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []time.Time{}
	for rows.Next() {
		var created_at time.Time
		if err := rows.Scan(&created_at); err != nil {
			return nil, err
		}
		items = append(items, created_at)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCustomer = `-- name: UpdateCustomer :one
UPDATE
  customer.customer
SET
  first_name = $2,
  last_name = $3,
  email = $4,
  phone = $5,
  identification_number = $6,
  identification_type = $7,
  modified_at = NOW()
WHERE
  id = $1
RETURNING id, company_id, first_name, last_name, email, phone, identification_number, identification_type, created_at, modified_at
`

type UpdateCustomerParams struct {
	ID                   int64          `json:"id"`
	FirstName            string         `json:"first_name"`
	LastName             sql.NullString `json:"last_name"`
	Email                sql.NullString `json:"email"`
	Phone                sql.NullString `json:"phone"`
	IdentificationNumber string         `json:"identification_number"`
	IdentificationType   string         `json:"identification_type"`
}

func (q *Queries) UpdateCustomer(ctx context.Context, arg UpdateCustomerParams) (CustomerCustomer, error) {
	row := q.db.QueryRowContext(ctx, updateCustomer,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Phone,
		arg.IdentificationNumber,
		arg.IdentificationType,
	)
	var i CustomerCustomer
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.Phone,
		&i.IdentificationNumber,
		&i.IdentificationType,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}
