// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: payment_type.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const getPaymentType = `-- name: GetPaymentType :one
SELECT
  id,
  name,
  image_path,
  created_at,
  modified_at
FROM
  company.payment_type
WHERE
  id = $1
LIMIT 1
`

type GetPaymentTypeRow struct {
	ID         int64          `json:"id"`
	Name       string         `json:"name"`
	ImagePath  sql.NullString `json:"image_path"`
	CreatedAt  time.Time      `json:"created_at"`
	ModifiedAt time.Time      `json:"modified_at"`
}

func (q *Queries) GetPaymentType(ctx context.Context, id int64) (GetPaymentTypeRow, error) {
	row := q.db.QueryRowContext(ctx, getPaymentType, id)
	var i GetPaymentTypeRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.ImagePath,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const getPaymentTypes = `-- name: GetPaymentTypes :many
SELECT
  id,
  name,
  image_path,
  created_at,
  modified_at
FROM
  company.payment_type
LIMIT
  $1
OFFSET
  $2
`

type GetPaymentTypesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetPaymentTypesRow struct {
	ID         int64          `json:"id"`
	Name       string         `json:"name"`
	ImagePath  sql.NullString `json:"image_path"`
	CreatedAt  time.Time      `json:"created_at"`
	ModifiedAt time.Time      `json:"modified_at"`
}

func (q *Queries) GetPaymentTypes(ctx context.Context, arg GetPaymentTypesParams) ([]GetPaymentTypesRow, error) {
	rows, err := q.db.QueryContext(ctx, getPaymentTypes, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPaymentTypesRow{}
	for rows.Next() {
		var i GetPaymentTypesRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.ImagePath,
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
