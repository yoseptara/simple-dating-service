package main

import (
	"log"
	"simple-dating-app-service/api"
	"simple-dating-app-service/bootstrap"
)

func main() {
	app := bootstrap.App()
	defer app.ClosePostgresConn()

	server, err := api.NewServer(app.Env, app.Store, app.Timeout)

	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = server.Start()
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
