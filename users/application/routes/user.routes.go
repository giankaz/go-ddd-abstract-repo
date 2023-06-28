package users

import (
	"github.com/gorilla/mux"
	handlers "mongodb.com/users/application/handlers"
)

func UserRoutes(router *mux.Router) {
	router.HandleFunc("/users", handlers.UserHandler.Create).Methods("POST")
	router.HandleFunc("/users", handlers.UserHandler.List).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.UserHandler.Update).Methods("PATCH")
	router.HandleFunc("/users/{id}", handlers.UserHandler.FindOne).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.UserHandler.Delete).Methods("DELETE")
}
