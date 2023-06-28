package users

import (
	"net/url"
	"testing"

	domain "mongodb.com/users/domain/entities"
)

func TestListUserUseCase(t *testing.T) {
	user, err := domain.NewUser("mock", 43)
	user2, _ := domain.NewUser("mock", 43)

	if err != nil {
		t.Error(err)
	}

	user, err = CreateUserUseCase(user)

	if err != nil {
		t.Error(err)
	}

	_, err = CreateUserUseCase(user2)

	if err != nil {
		t.Error(err)
	}

	params := url.Values{}

	listResult, err := ListUserUseCase(params)

	if len(listResult.Items) != 2 {
		t.Error("expected 2 items, instead got: ", len(listResult.Items))
	}

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
