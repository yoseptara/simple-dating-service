package country

type InsertUserReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
}

type ListSwipableProfilesReq struct {
	UserId string `form:"user_id" json:"user_id" binding:"required"`
	Limit  int    `form:"limit" json:"limit" binding:"required"`
}
