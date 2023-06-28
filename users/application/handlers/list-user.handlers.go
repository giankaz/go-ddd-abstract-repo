package users

import (
	"encoding/json"
	"net/http"

	useCases "mongodb.com/users/application/use-cases"
)

func (h *UserHandlerStruct) List(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	listResponse, err := useCases.ListUserUseCase(query)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	json.NewEncoder(w).Encode(listResponse)

}
