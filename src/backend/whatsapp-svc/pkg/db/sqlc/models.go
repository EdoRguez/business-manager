// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"time"
)

type WhatsappBusinessPhone struct {
	ID         int64     `json:"id"`
	CompanyID  int64     `json:"company_id"`
	Phone      string    `json:"phone"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}
