package users

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	useCases "mongodb.com/users/application/use-cases"
	domain "mongodb.com/users/domain/entities"
)

func (h *UserHandlerStruct) Update(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	id := mux.Vars(r)["id"]

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = useCases.UpdateUserUseCase(id, user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	user, err = useCases.FindUserUseCase(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(user)

}
