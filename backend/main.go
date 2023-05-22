package main

import (
	"backend/api"
	"backend/model"
	"backend/repository"
	"backend/service"
	"log"
	"net/http"
	"os"
)

func main() {
	config := loadConfig()
	app, err := newApplication(config)
	if err != nil {
		panic(err)
	}

	app.start(config)
}

func loadConfig() model.Configuration {
	config := model.Configuration{
		Username:   os.Getenv("LOGIN_USER"),
		Password:   os.Getenv("LOGIN_PASSWORD_HASH"),
		HMACSecret: os.Getenv("HMAC_SECRET"),
		Salt:       os.Getenv("SALT"),
		DBDatabase: os.Getenv("POSTGRES_DB"),
		DBHost:     os.Getenv("POSTGRES_HOST"),
		DBUser:     os.Getenv("POSTGRES_USER"),
		DBPassword: os.Getenv("POSTGRES_PASSWORD"),
	}
	if (config == model.Configuration{}) {
		panic("empty config")
	}
	return config
}

type application struct {
	port       string
	service    api.DhcpdService
	repository api.DhcpdRepository
}

func newApplication(cfg model.Configuration) (*application, error) {
	repo, err := repository.New(cfg.DSN())
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

func (app application) start(config model.Configuration) {
	router := api.NewRouter(app.service, config)
	log.Println("Listening at Port " + app.port)
	err := http.ListenAndServe(":"+app.port, router)
	panic(err)
}
