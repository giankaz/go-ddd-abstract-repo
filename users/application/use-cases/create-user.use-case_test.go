package users

import (
	"testing"

	users "mongodb.com/users/domain/entities"
)

func TestCreateUserUseCase(t *testing.T) {
	user, err := users.NewUser("mock", 43)

	if err != nil {
		t.Error(err)
	}

	user, err = CreateUserUseCase(user)

	if err != nil {
		t.Error(err)
	}

	if user.Name != "mock" {
		t.Error("wrong name was saved")
	}

	if user.Age != 43 {
		t.Error("wrong age was saved")
	}
}
