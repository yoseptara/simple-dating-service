package bootstrap

import (
	"database/sql"
	"esim-service/config"
	db "esim-service/db/sqlc"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Application struct {
	Env          config.Env
	PostgresConn *sql.DB
	Store        db.Store
	Timeout      time.Duration
}

func App() Application {
	app := &Application{}

	env, err := config.NewEnv()
	if err != nil {
		log.Fatal("cannot get env variables", err)
	}

	app.Env = env

	timeout := time.Duration(env.ContextTimeout) * time.Second
	app.Timeout = timeout

	conn, err := sql.Open("postgres", env.PostgresConnStr)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	app.PostgresConn = conn

	store := db.NewStore(conn)
	app.Store = store

	return *app
}

func (app *Application) ClosePostgresConn() {
	app.PostgresConn.Close()
}
