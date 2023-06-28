package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"mongodb.com/users/application/routes"
)

func routes() http.Handler {
	router := mux.NewRouter()

	users.UserRoutes(router)

	return router
}
