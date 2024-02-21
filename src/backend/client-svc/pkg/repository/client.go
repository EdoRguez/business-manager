package repository

import (
	"context"

	db "github.com/EdoRguez/business-manager/client-svc/pkg/db/sqlc"
)

type ClientRepo struct {
	SQLStorage *db.SQLStorage
}

func NewClientRepo(sql *db.SQLStorage) *ClientRepo {
	return &ClientRepo{
		SQLStorage: sql,
	}
}

func (clientRepo *ClientRepo) CreateClient(ctx context.Context, arg db.CreateClientParams) (db.ClientClient, error) {
	var result db.ClientClient

	err := clientRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.CreateClient(ctx, arg)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}
