package common

import (
	"encoding/json"
	"net/http"
)

func HttpErrorHandler(w http.ResponseWriter, errs []error, status ...int) {
	var treatedErrors []IError

	for _, value := range errs {
		treatedErrors = append(treatedErrors, IError{
			Message: value.Error(),
		})
	}

	result := ErrorStruct{
		Errors: treatedErrors,
	}

	if len(status) == 1 {
		w.WriteHeader(status[0])

	} else {

		w.WriteHeader(http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(result)
}
