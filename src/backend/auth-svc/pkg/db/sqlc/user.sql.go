// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO 
  auth.user (
    company_id,
    role_id,
    email,
    password_hash
  ) 
VALUES (
  $1, $2, $3, $4
) 
RETURNING id, company_id, role_id, email, password_hash, created_at, modified_at
`

type CreateUserParams struct {
	CompanyID    int64  `json:"company_id"`
	RoleID       int64  `json:"role_id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (AuthUser, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.CompanyID,
		arg.RoleID,
		arg.Email,
		arg.PasswordHash,
	)
	var i AuthUser
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.RoleID,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM 
  auth.user
WHERE 
  id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT 
  id,
  company_id,
  role_id,
  email,
  password_hash,
  created_at,
  modified_at
FROM 
  auth.user
WHERE 
  id = $1 
LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (AuthUser, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i AuthUser
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.RoleID,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT 
  id,
  company_id,
  role_id,
  email,
  password_hash,
  created_at,
  modified_at
FROM 
  auth.user
WHERE 
  email = $1 
LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (AuthUser, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i AuthUser
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.RoleID,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT 
  id,
  company_id,
  role_id,
  email,
  password_hash,
  created_at,
  modified_at
FROM 
  auth.user
WHERE
  (company_id = $1) OR $1 = 0
ORDER BY 
  id
LIMIT 
  $2
OFFSET 
  $3
`

type GetUsersParams struct {
	CompanyID int64 `json:"company_id"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

func (q *Queries) GetUsers(ctx context.Context, arg GetUsersParams) ([]AuthUser, error) {
	rows, err := q.db.QueryContext(ctx, getUsers, arg.CompanyID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AuthUser{}
	for rows.Next() {
		var i AuthUser
		if err := rows.Scan(
			&i.ID,
			&i.CompanyID,
			&i.RoleID,
			&i.Email,
			&i.PasswordHash,
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

const updateUser = `-- name: UpdateUser :one
UPDATE 
  auth.user
SET 
  role_id = $2,
  email = $3,
  password_hash = $4,
  modified_at = NOW()
WHERE 
  id = $1
RETURNING id, company_id, role_id, email, password_hash, created_at, modified_at
`

type UpdateUserParams struct {
	ID           int64  `json:"id"`
	RoleID       int64  `json:"role_id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (AuthUser, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.ID,
		arg.RoleID,
		arg.Email,
		arg.PasswordHash,
	)
	var i AuthUser
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.RoleID,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}
