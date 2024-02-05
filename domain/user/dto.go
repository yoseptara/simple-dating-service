package user

import (
	db "simple-dating-app-service/db/sqlc"
)

type LoginUserReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterUserReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
}

type ListSwipableProfilesReq struct {
	UserId string `form:"user_id" json:"user_id" binding:"required"`
	Limit  int    `form:"limit" json:"limit" binding:"required"`
}

type UserResponse struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type LoginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        UserResponse `json:"user"`
}

func NewUserResponse(user db.User) UserResponse {
	return UserResponse{
		Username: user.Username,
		FullName: user.Fullname,
		Email:    user.Email,
	}
}
