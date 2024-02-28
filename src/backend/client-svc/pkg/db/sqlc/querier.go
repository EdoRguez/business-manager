// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
)

type Querier interface {
	CreateClient(ctx context.Context, arg CreateClientParams) (ClientClient, error)
	DeleteClient(ctx context.Context, id int64) error
	GetClient(ctx context.Context, id int64) (ClientClient, error)
	GetClients(ctx context.Context, arg GetClientsParams) ([]ClientClient, error)
	GetClientsByCompanyId(ctx context.Context, arg GetClientsByCompanyIdParams) ([]ClientClient, error)
	UpdateClient(ctx context.Context, arg UpdateClientParams) (ClientClient, error)
}

var _ Querier = (*Queries)(nil)