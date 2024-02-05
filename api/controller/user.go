package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"simple-dating-app-service/config"
	db "simple-dating-app-service/db/sqlc"
	"simple-dating-app-service/domain"
	"simple-dating-app-service/domain/user"
	"simple-dating-app-service/token"
	"simple-dating-app-service/util"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Usecase    user.Usecase
	Env        *config.Env
	TokenMaker token.Maker
}

func (uc *UserController) Register(c *gin.Context) {
	var req user.RegisterUserReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return
	}

	arg := db.CreateUserParams{
		Username: req.Username,
		Password: hashedPassword,
		Fullname: req.Fullname,
		Email:    req.Email,
	}

	user, err := uc.Usecase.CreateUser(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) Login(ctx *gin.Context) {
	var req user.LoginUserReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	fetchedUser, err := uc.Usecase.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, domain.ErrorResponse{Message: err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = util.CheckPassword(req.Password, fetchedUser.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, err := uc.TokenMaker.CreateToken(
		fetchedUser.Username,
		uc.Env.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	rsp := user.LoginUserResponse{
		AccessToken: accessToken,
		User:        user.NewUserResponse(fetchedUser),
	}
	ctx.JSON(http.StatusOK, rsp)
}
