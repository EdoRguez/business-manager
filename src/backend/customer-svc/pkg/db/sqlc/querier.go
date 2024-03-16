// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
)

type Querier interface {
	CreateCustomer(ctx context.Context, arg CreateCustomerParams) (CustomerCustomer, error)
	DeleteCustomer(ctx context.Context, id int64) error
	GetCustomer(ctx context.Context, id int64) (CustomerCustomer, error)
	GetCustomers(ctx context.Context, arg GetCustomersParams) ([]CustomerCustomer, error)
	UpdateCustomer(ctx context.Context, arg UpdateCustomerParams) (CustomerCustomer, error)
}

var _ Querier = (*Queries)(nil)
