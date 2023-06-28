package users

import (
	"testing"

	domain "mongodb.com/users/domain/entities"
	infra "mongodb.com/users/infra"
)

func TestUpdateUserUseCase(t *testing.T) {
	user, err := domain.NewUser("mock", 43)

	if err != nil {
		t.Error(err)
	}

	user, err = CreateUserUseCase(user)

	if err != nil {
		t.Error(err)
	}

	user.ChangeAge(20)
	user.ChangeName("new name")

	err = UpdateUserUseCase(user.Id, user)

	if err != nil {
		t.Error(err)
	}

	user, err = infra.UserRepository.FindOne(user.Id)

	if err != nil {
		t.Error(err)
	}

	if user.Age != 20 {
		t.Error("wrong age on update")
	}

	if user.Name != "new name" {
		t.Error("wrong name on update")
	}
}
