package main

import (
	"backend/api"
	"backend/service"
	"net/http"
)

func main() {
	s := service.New()
	router := api.NewRouter(s)

	err := http.ListenAndServe("localhost:8000", router)
	panic(err)
}
