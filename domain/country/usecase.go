package country

import (
	"context"
	db "esim-service/db/sqlc"
)

type Usecase interface {
	Insert(c context.Context, params db.CreateCountryParams) error
	List(c context.Context) (*[]db.ListCountriesWithPriceRow, error)
}
