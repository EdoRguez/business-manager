// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: whatsapp.sql

package db

import (
	"context"
)

const createBusinessPhone = `-- name: CreateBusinessPhone :one
INSERT INTO
  whatsapp.business_phone (
    company_id,
    phone
  )
VALUES (
  $1, $2
)
RETURNING id, company_id, phone, created_at, modified_at
`

type CreateBusinessPhoneParams struct {
	CompanyID int64  `json:"company_id"`
	Phone     string `json:"phone"`
}

func (q *Queries) CreateBusinessPhone(ctx context.Context, arg CreateBusinessPhoneParams) (WhatsappBusinessPhone, error) {
	row := q.db.QueryRowContext(ctx, createBusinessPhone, arg.CompanyID, arg.Phone)
	var i WhatsappBusinessPhone
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.Phone,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const getBusinessPhoneByCompanyId = `-- name: GetBusinessPhoneByCompanyId :one
SELECT
  id,
  company_id,
  phone,
  created_at,
  modified_at
FROM
  whatsapp.business_phone
WHERE
  company_id = $1
LIMIT 1
`

func (q *Queries) GetBusinessPhoneByCompanyId(ctx context.Context, companyID int64) (WhatsappBusinessPhone, error) {
	row := q.db.QueryRowContext(ctx, getBusinessPhoneByCompanyId, companyID)
	var i WhatsappBusinessPhone
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.Phone,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const updateBusinessPhone = `-- name: UpdateBusinessPhone :one
UPDATE
  whatsapp.business_phone
SET
  phone = $1,
  modified_at = NOW()
WHERE
  company_id = $2
RETURNING id, company_id, phone, created_at, modified_at
`

type UpdateBusinessPhoneParams struct {
	Phone     string `json:"phone"`
	CompanyID int64  `json:"company_id"`
}

func (q *Queries) UpdateBusinessPhone(ctx context.Context, arg UpdateBusinessPhoneParams) (WhatsappBusinessPhone, error) {
	row := q.db.QueryRowContext(ctx, updateBusinessPhone, arg.Phone, arg.CompanyID)
	var i WhatsappBusinessPhone
	err := row.Scan(
		&i.ID,
		&i.CompanyID,
		&i.Phone,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}
