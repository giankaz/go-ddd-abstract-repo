package users

import (
	"net/http"

	"github.com/gorilla/mux"
	useCases "mongodb.com/users/application/use-cases"
)

func (h *UserHandlerStruct) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := useCases.DeleteUserUseCase(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	_, err = useCases.FindUserUseCase(id)

	if err == nil {
		http.Error(w, "error deleting user", http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)

}
