package usecase

import (
	"context"
	db "simple-dating-app-service/db/sqlc"
	"simple-dating-app-service/domain"
	"simple-dating-app-service/domain/swipe"
)

type swipeUsecase struct {
	server *domain.ConcreteServer
}

func NewSwipeUsecase(server *domain.ConcreteServer) swipe.Usecase {
	return &swipeUsecase{
		server: server,
	}
}

func (su *swipeUsecase) CreateOrUpdateSwipe(c context.Context, params db.CreateOrUpdateSwipeParams) (*db.Swipe, error) {
	ctx, cancel := context.WithTimeout(c, su.server.Timeout)
	defer cancel()

	res, err := su.server.Store.CreateOrUpdateSwipe(ctx, params)

	return &res, err
}
