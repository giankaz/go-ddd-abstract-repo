package common

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	er "mongodb.com/common/domain/error"
)

var NewValidator = validator.New()

func (e *Entity) Validate(entity interface{}) (bool, error) {
	var errorsSlice []er.IError

	err := NewValidator.Struct(entity)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {

			errorsSlice = append(errorsSlice, er.IError{
				Field:   err.Field(),
				Message: err.Error(),
			})
		}

		response := er.ErrorStruct{
			Errors: errorsSlice,
		}

		jsonErrors, _ := json.Marshal(response)

		err = errors.New(string(jsonErrors))
		fmt.Println(err)
		return false, err
	}

	return true, nil
}
