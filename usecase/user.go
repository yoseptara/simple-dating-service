package usecase

import (
	"context"
	db "simple-dating-app-service/db/sqlc"
	"simple-dating-app-service/domain"
	"simple-dating-app-service/domain/user"
)

type userUsecase struct {
	server *domain.ConcreteServer
}

func NewUserUsecase(server *domain.ConcreteServer) user.Usecase {
	return &userUsecase{
		server: server,
	}
}

func (uu *userUsecase) CreateUser(c context.Context, params db.CreateUserParams) (db.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.server.Timeout)
	defer cancel()

	user, err := uu.server.Store.CreateUser(ctx, params)

	return user, err
}

func (uu *userUsecase) GetUser(c context.Context, username string) (db.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.server.Timeout)
	defer cancel()

	user, err := uu.server.Store.GetUser(ctx, username)

	return user, err
}

func (uu *userUsecase) ListSwipableProfiles(c context.Context, params db.ListSwipableProfilesParams) (*[]db.ListSwipableProfilesRow, error) {
	ctx, cancel := context.WithTimeout(c, uu.server.Timeout)
	defer cancel()

	countries, err := uu.server.Store.ListSwipableProfiles(ctx, params)

	return &countries, err
}
