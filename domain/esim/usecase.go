package esim

import (
	"context"
	db "esim-service/db/sqlc"
)

type Usecase interface {
	Create(c context.Context, params db.CreateEsimParams) (*db.Esim, error)
	ListByCountry(c context.Context, countryCode string) (*[]db.ListEsimsByCountryRow, error)
	// UpdateStock(c context.Context, params db.UpdateEsimStockParams) (*db.Esim, error)
	Delete(c context.Context, id int64) error
}
