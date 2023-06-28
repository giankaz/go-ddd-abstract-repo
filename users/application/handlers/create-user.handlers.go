package users

import (
	"encoding/json"
	"net/http"

	useCases "mongodb.com/users/application/use-cases"
	domain "mongodb.com/users/domain/entities"
)

func (h *UserHandlerStruct) Create(w http.ResponseWriter, r *http.Request) {
	var user domain.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	newUser, err := domain.NewUser(user.Name, user.Age)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

	}

	createdUser, err := useCases.CreateUserUseCase(newUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(createdUser)

}
