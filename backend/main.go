package main

import (
	"backend/api"
	"backend/repository"
	"backend/service"
	"net/http"
)

type Config struct {
	DSN string
}

func main() {
	app, err := newApplication(Config{
		DSN: "host=localhost user=network password=network dbname=network port=5432 sslmode=disable",
	})
	if err != nil {
		panic(err)
	}

	app.start()
}

type application struct {
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
	}, nil
}

func (app application) start() {
	router := api.NewRouter(app.service)
	err := http.ListenAndServe("localhost:8000", router)
	panic(err)
}
