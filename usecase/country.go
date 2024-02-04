package usecase

import (
	"context"
	db "esim-service/db/sqlc"
	"esim-service/domain"
	"esim-service/domain/country"
)

type countryUsecase struct {
	server *domain.ConcreteServer
}

func NewCountryUsecase(server *domain.ConcreteServer) country.Usecase {
	return &countryUsecase{
		server: server,
	}
}

func (cu *countryUsecase) Insert(c context.Context, params db.CreateCountryParams) error {
	ctx, cancel := context.WithTimeout(c, cu.server.Timeout)
	defer cancel()

	err := cu.server.Store.CreateCountry(ctx, params)

	return err
}

func (cu *countryUsecase) List(c context.Context) (*[]db.ListCountriesWithPriceRow, error) {
	ctx, cancel := context.WithTimeout(c, cu.server.Timeout)
	defer cancel()

	countries, err := cu.server.Store.ListCountriesWithPrice(ctx)

	return &countries, err
}
