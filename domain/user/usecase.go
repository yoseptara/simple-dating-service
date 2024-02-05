package user

import (
	"context"
	db "simple-dating-app-service/db/sqlc"
)

type Usecase interface {
	CreateUser(c context.Context, params db.CreateUserParams) (db.User, error)
	GetUser(c context.Context, username string) (db.User, error)
	ListSwipableProfiles(c context.Context) (*[]db.ListSwipableProfilesRow, error)
}
