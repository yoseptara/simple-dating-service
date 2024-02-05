package swipe

import (
	"context"
	db "simple-dating-app-service/db/sqlc"
)

type Usecase interface {
	CreateOrUpdateSwipe(c context.Context, params db.CreateOrUpdateSwipeParams) (*db.Swipe, error)
}
