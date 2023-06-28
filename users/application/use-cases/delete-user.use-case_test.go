package users

import (
	"testing"

	domain "mongodb.com/users/domain/entities"
	infra "mongodb.com/users/infra"
)

func TestDeleteUserUseCase(t *testing.T) {
	user, err := domain.NewUser("mock", 43)

	if err != nil {
		t.Error(err)
	}

	user, err = CreateUserUseCase(user)

	if err != nil {
		t.Error(err)
	}

	err = DeleteUserUseCase(user.Id)

	if err != nil {
		t.Error(err)
	}

	_, err = infra.UserRepository.FindOne(user.Id)

	if err == nil {
		t.Error("found user when it should be deleted")
	}
}
