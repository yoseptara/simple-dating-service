package country

import (
	"context"
	db "simple-dating-app-service/db/sqlc"
)

type Usecase interface {
	Create(c context.Context, params db.CreateUserParams) error
	ListSwipableProfiles(c context.Context) (*[]db.ListSwipableProfilesRow, error)
}
