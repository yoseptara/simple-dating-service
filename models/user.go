package models

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Swipe struct {
	UserID    uint   `json:"user_id"`
	TargetID  uint   `json:"target_id"`
	Direction string `json:"direction"` // "left" or "right"
}

type PremiumFeature struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}
