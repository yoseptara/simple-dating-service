package usecase

import (
	"context"
	db "esim-service/db/sqlc"
	"esim-service/domain"
	"esim-service/domain/esim"
)

type esimUsecase struct {
	server *domain.ConcreteServer
}

func NewEsimUsecase(server *domain.ConcreteServer) esim.Usecase {
	return &esimUsecase{
		server: server,
	}
}

func (eu *esimUsecase) Create(c context.Context, params db.CreateEsimParams) (*db.Esim, error) {
	ctx, cancel := context.WithTimeout(c, eu.server.Timeout)
	defer cancel()

	res, err := eu.server.Store.CreateEsim(ctx, params)

	return &res, err
}

func (eu *esimUsecase) ListByCountry(c context.Context, countryCode string) (*[]db.ListEsimsByCountryRow, error) {
	ctx, cancel := context.WithTimeout(c, eu.server.Timeout)
	defer cancel()

	res, err := eu.server.Store.ListEsimsByCountry(ctx, countryCode)

	return &res, err
}

// func (eu *esimUsecase) UpdateStock(c context.Context, params db.UpdateEsimStockParams) (*db.Esim, error) {
// 	ctx, cancel := context.WithTimeout(c, eu.contextTimeout)
// 	defer cancel()

// 	res, err := eu.server.Store.UpdateEsimStock(ctx, params)

// 	return &res, err
// }

func (eu *esimUsecase) Delete(c context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(c, eu.server.Timeout)
	defer cancel()

	err := eu.server.Store.DeleteEsim(ctx, id)

	return err
}
