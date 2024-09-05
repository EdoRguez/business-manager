// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
)

type Querier interface {
	CreateCompany(ctx context.Context, arg CreateCompanyParams) (CompanyCompany, error)
	CreatePayment(ctx context.Context, arg CreatePaymentParams) (CompanyPayment, error)
	DeleteCompany(ctx context.Context, id int64) error
	DeletePayment(ctx context.Context, id int64) error
	GetCompanies(ctx context.Context, arg GetCompaniesParams) ([]CompanyCompany, error)
	GetCompany(ctx context.Context, id int64) (CompanyCompany, error)
	GetPayment(ctx context.Context, id int64) (GetPaymentRow, error)
	GetPaymentType(ctx context.Context, id int64) (GetPaymentTypeRow, error)
	GetPaymentTypes(ctx context.Context, arg GetPaymentTypesParams) ([]GetPaymentTypesRow, error)
	GetPayments(ctx context.Context, arg GetPaymentsParams) ([]GetPaymentsRow, error)
	GetPaymentsTypes(ctx context.Context, companyID int64) ([]GetPaymentsTypesRow, error)
	UpdateCompany(ctx context.Context, arg UpdateCompanyParams) (CompanyCompany, error)
	UpdatePayment(ctx context.Context, arg UpdatePaymentParams) (CompanyPayment, error)
	UpdatePaymentStatus(ctx context.Context, arg UpdatePaymentStatusParams) (CompanyPayment, error)
}

var _ Querier = (*Queries)(nil)
