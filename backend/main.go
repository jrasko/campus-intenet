package main

import (
	"backend/api"
	"backend/repository"
	"backend/service"
	"fmt"
	"net/http"
)

type Config struct {
	DSN string
}

func main() {
	app, err := newApplication(Config{
		DSN: "host=dhcpd_db user=network password=network dbname=network port=5432 sslmode=disable",
	})
	if err != nil {
		panic(err)
	}

	config := api.Configuration{}
	app.start(config)
}

type application struct {
	port       string
	service    api.DhcpdService
	repository api.DhcpdRepository
}

func newApplication(cfg Config) (*application, error) {
	repo, err := repository.New(cfg.DSN)
	if err != nil {
		return nil, err
	}
	srv := service.New(repo)

	return &application{
		repository: repo,
		service:    srv,
		port:       "8080",
	}, nil
}

func (app application) start(config api.Configuration) {
	router := api.NewRouter(app.service, config)
	fmt.Println("Listening at Port " + app.port)
	err := http.ListenAndServe(":"+app.port, router)
	panic(err)
}
