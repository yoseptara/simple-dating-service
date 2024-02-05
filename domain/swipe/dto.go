package swipe

import db "simple-dating-app-service/db/sqlc"

type CreateOrUpdateSwipeReq struct {
	UserId    int64        `json:"user_id" binding:"required"`
	TargetId  int64        `json:"target_id" binding:"required"`
	Direction db.Direction `json:"direction" binding:"required"`
}
