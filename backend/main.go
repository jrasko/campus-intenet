package main

import (
	"backend/api"
	"backend/model"
	"backend/repository"
	"backend/service"
	"log"
	"net/http"
)

type application struct {
	url        string
	service    api.DhcpdService
	repository api.DhcpdRepository
	router     http.Handler
}

func main() {
	config := model.LoadConfig()

	app, err := newApplication(config)
	if err != nil {
		panic(err)
	}

	app.start()
}

func newApplication(config model.Configuration) (*application, error) {
	repo, err := repository.New(config.DSN())
	if err != nil {
		return nil, err
	}

	srv := service.New(config, repo)
	router := api.NewRouter(config, srv)

	return &application{
		repository: repo,
		service:    srv,
		url:        config.URL,
		router:     router,
	}, nil
}

func (app application) start() {
	log.Println("Listening at " + app.url)
	err := http.ListenAndServe(app.url, app.router)
	panic(err)
}
