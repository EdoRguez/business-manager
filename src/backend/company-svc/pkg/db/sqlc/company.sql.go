// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: company.sql

package db

import (
	"context"
	"database/sql"
)

const createCompany = `-- name: CreateCompany :one
INSERT INTO 
  company.company (
    name,
    image_url
  ) 
VALUES (
  $1, $2 
) 
RETURNING id, name, image_url, created_at, modified_at
`

type CreateCompanyParams struct {
	Name     string         `json:"name"`
	ImageUrl sql.NullString `json:"image_url"`
}

func (q *Queries) CreateCompany(ctx context.Context, arg CreateCompanyParams) (CompanyCompany, error) {
	row := q.db.QueryRowContext(ctx, createCompany, arg.Name, arg.ImageUrl)
	var i CompanyCompany
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ImageUrl,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const deleteCompany = `-- name: DeleteCompany :exec
DELETE FROM 
  company.company
WHERE 
  id = $1
`

func (q *Queries) DeleteCompany(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCompany, id)
	return err
}

const getCompanies = `-- name: GetCompanies :many
SELECT 
  id,
  name,
  image_url,
  created_at,
  modified_at
FROM 
  company.company
ORDER BY 
  id
LIMIT 
  $1
OFFSET 
  $2
`

type GetCompaniesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetCompanies(ctx context.Context, arg GetCompaniesParams) ([]CompanyCompany, error) {
	rows, err := q.db.QueryContext(ctx, getCompanies, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CompanyCompany{}
	for rows.Next() {
		var i CompanyCompany
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.ImageUrl,
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

const getCompany = `-- name: GetCompany :one
SELECT 
  id,
  name,
  image_url,
  created_at,
  modified_at
FROM 
  company.company
WHERE 
  id = $1 
LIMIT 1
`

func (q *Queries) GetCompany(ctx context.Context, id int64) (CompanyCompany, error) {
	row := q.db.QueryRowContext(ctx, getCompany, id)
	var i CompanyCompany
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ImageUrl,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const updateCompany = `-- name: UpdateCompany :one
UPDATE 
  company.company
SET 
  name = $2,
  image_url = $3,
  modified_at = NOW()
WHERE 
  id = $1
RETURNING id, name, image_url, created_at, modified_at
`

type UpdateCompanyParams struct {
	ID       int64          `json:"id"`
	Name     string         `json:"name"`
	ImageUrl sql.NullString `json:"image_url"`
}

func (q *Queries) UpdateCompany(ctx context.Context, arg UpdateCompanyParams) (CompanyCompany, error) {
	row := q.db.QueryRowContext(ctx, updateCompany, arg.ID, arg.Name, arg.ImageUrl)
	var i CompanyCompany
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ImageUrl,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}
