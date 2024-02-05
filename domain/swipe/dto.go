package esim

import db "simple-dating-app-service/db/sqlc"

type CreateOrUpdateSwipeReq struct {
	UserId    string       `json:"user_id" binding:"required"`
	TargetId  string       `json:"target_id" binding:"required"`
	Direction db.Direction `json:"direction" binding:"required"`
}
