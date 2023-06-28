package users

import (
	"testing"

	domain "mongodb.com/users/domain/entities"
)

func TestFindUserUseCase(t *testing.T) {
	user, err := domain.NewUser("mock", 43)

	if err != nil {
		t.Error(err)
	}

	user, err = CreateUserUseCase(user)

	if err != nil {
		t.Error(err)
	}

	user, err = FindUserUseCase(user.Id)

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
