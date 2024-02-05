package controller

import (
	"database/sql"

	"net/http"
	"simple-dating-app-service/config"
	db "simple-dating-app-service/db/sqlc"
	"simple-dating-app-service/domain"
	"simple-dating-app-service/domain/swipe"

	"github.com/gin-gonic/gin"
)

type SwipeController struct {
	Usecase swipe.Usecase
	Env     *config.Env
}

func (sc *SwipeController) CreateSwipe(c *gin.Context) {
	var req swipe.CreateOrUpdateSwipeReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	arg := db.CreateOrUpdateSwipeParams{
		UserID: sql.NullInt64{
			Valid: true,
			Int64: req.UserId,
		},
		TargetID:  req.TargetId,
		Direction: req.Direction,
	}

	swipe, err := sc.Usecase.CreateOrUpdateSwipe(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, swipe)
}
