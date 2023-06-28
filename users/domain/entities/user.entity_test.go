package users

import "testing"

func TestUserEntity(t *testing.T) {
	user := User{
		Id:   "fake",
		Name: "name",
		Age:  43,
	}

	err := user.ChangeAge(10)

	if err != nil {
		t.Error(err)
	}

	err = user.ChangeName("new Name")

	if err != nil {
		t.Error(err)
	}

	if user.Name != "new Name" {
		t.Error("error while changing name")
	}

	if user.Age != 10 {
		t.Error("error while changing age")
	}

	err = user.ChangeName("f")

	if err == nil {
		t.Error("expected error on user validation")
	}
}
