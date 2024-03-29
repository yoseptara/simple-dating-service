// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

type Direction string

const (
	DirectionLeft  Direction = "left"
	DirectionRight Direction = "right"
)

func (e *Direction) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Direction(s)
	case string:
		*e = Direction(s)
	default:
		return fmt.Errorf("unsupported scan type for Direction: %T", src)
	}
	return nil
}

type NullDirection struct {
	Direction Direction `json:"Direction"`
	Valid     bool      `json:"valid"` // Valid is true if Direction is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullDirection) Scan(value interface{}) error {
	if value == nil {
		ns.Direction, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Direction.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullDirection) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Direction), nil
}

type Swipe struct {
	ID        int64         `json:"id"`
	UserID    sql.NullInt64 `json:"user_id"`
	TargetID  int64         `json:"target_id"`
	Direction interface{}   `json:"direction"`
	SwipeDate sql.NullTime  `json:"swipe_date"`
	Unique    interface{}   `json:"unique"`
}

type User struct {
	ID            int64         `json:"id"`
	Username      string        `json:"username"`
	Password      string        `json:"password"`
	Fullname      string        `json:"fullname"`
	Email         string        `json:"email"`
	TotalSwipes   sql.NullInt32 `json:"total_swipes"`
	LastSwipeDate sql.NullTime  `json:"last_swipe_date"`
	SwipeCount    sql.NullInt32 `json:"swipe_count"`
	IsPremium     sql.NullBool  `json:"is_premium"`
}
