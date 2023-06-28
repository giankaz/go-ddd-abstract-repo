package common

import "testing"

type entityToValidate struct {
	Entity `bson:"-" json:"-"`
	ID     string `validate:"required,min=3,max=12" json:"id,omitempty"`
	Name   string `validate:"required,min=3,max=12" json:"name,omitempty"`
}

func TestValidator(t *testing.T) {
	newEntity := entityToValidate{
		Name: "f",
		ID:   "fake-id",
	}

	_, err := newEntity.Validate(newEntity)

	if err == nil {
		t.Error("validation failed on name")
	}

	newEntity = entityToValidate{
		Name: "fake-name",
		ID:   "f",
	}

	_, err = newEntity.Validate(newEntity)

	if err == nil {
		t.Error("validation failed on id")
	}

}
