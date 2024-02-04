package domain

import (
	"esim-service/config"
	db "esim-service/db/sqlc"
	"time"
)

// Server serves HTTP requests for our banking service.
type Server interface {
	Start() error
}

// ConcreteServer serves HTTP requests for our banking service.
type ConcreteServer struct {
	Config  config.Env
	Store   db.Store
	Timeout time.Duration
}
