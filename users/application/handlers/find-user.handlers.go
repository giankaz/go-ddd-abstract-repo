package users

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	common "mongodb.com/common/application"
	commonError "mongodb.com/common/domain/error"
	useCases "mongodb.com/users/application/use-cases"
)

func (h *UserHandlerStruct) FindOne(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	findResponse, err := useCases.FindUserUseCase(id)

	if err != nil {
		common.ErrorLog.Println(err)
		commonError.HttpErrorHandler(w, []error{err}, http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(findResponse)

}
