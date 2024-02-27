// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: client.sql

package db

import (
	"context"
	"database/sql"
)

const createClient = `-- name: CreateClient :one
INSERT INTO 
  client.client (
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

type CreateClientParams struct {
	CompanyID            int64          `json:"company_id"`
	FirstName            string         `json:"first_name"`
	LastName             sql.NullString `json:"last_name"`
	Email                sql.NullString `json:"email"`
	Phone                sql.NullString `json:"phone"`
	IdentificationNumber string         `json:"identification_number"`
	IdentificationType   string         `json:"identification_type"`
}

func (q *Queries) CreateClient(ctx context.Context, arg CreateClientParams) (ClientClient, error) {
	row := q.db.QueryRowContext(ctx, createClient,
		arg.CompanyID,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Phone,
		arg.IdentificationNumber,
		arg.IdentificationType,
	)
	var i ClientClient
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

const deleteClient = `-- name: DeleteClient :exec
DELETE FROM 
  client.client
WHERE 
  id = $1
`

func (q *Queries) DeleteClient(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteClient, id)
	return err
}

const getClient = `-- name: GetClient :one
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
  client.client
WHERE 
  id = $1 
LIMIT 1
`

func (q *Queries) GetClient(ctx context.Context, id int64) (ClientClient, error) {
	row := q.db.QueryRowContext(ctx, getClient, id)
	var i ClientClient
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

const getClients = `-- name: GetClients :many
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
  client.client
ORDER BY 
  id
LIMIT 
  $1
OFFSET 
  $2
`

type GetClientsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetClients(ctx context.Context, arg GetClientsParams) ([]ClientClient, error) {
	rows, err := q.db.QueryContext(ctx, getClients, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ClientClient{}
	for rows.Next() {
		var i ClientClient
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

const getClientsByCompanyId = `-- name: GetClientsByCompanyId :many
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
  client.client
WHERE
  company_id = $1
ORDER BY 
  id
LIMIT 
  $2
OFFSET 
  $3
`

type GetClientsByCompanyIdParams struct {
	CompanyID int64 `json:"company_id"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *Queries) GetClientsByCompanyId(ctx context.Context, arg GetClientsByCompanyIdParams) ([]ClientClient, error) {
	rows, err := q.db.QueryContext(ctx, getClientsByCompanyId, arg.CompanyID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ClientClient{}
	for rows.Next() {
		var i ClientClient
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

const updateClient = `-- name: UpdateClient :one
UPDATE 
  client.client
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

type UpdateClientParams struct {
	ID                   int64          `json:"id"`
	FirstName            string         `json:"first_name"`
	LastName             sql.NullString `json:"last_name"`
	Email                sql.NullString `json:"email"`
	Phone                sql.NullString `json:"phone"`
	IdentificationNumber string         `json:"identification_number"`
	IdentificationType   string         `json:"identification_type"`
}

func (q *Queries) UpdateClient(ctx context.Context, arg UpdateClientParams) (ClientClient, error) {
	row := q.db.QueryRowContext(ctx, updateClient,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Phone,
		arg.IdentificationNumber,
		arg.IdentificationType,
	)
	var i ClientClient
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
