package users

import (
	"github.com/google/uuid"
	entity "mongodb.com/common/domain/entity"
)

type User struct {
	entity.Entity `bson:"-" json:"-"`
	Id            string `bson:"_id,omitempty" json:"id,omitempty"`
	Name          string `bson:"name,omitempty" json:"name,omitempty" validate:"required,min=3,max=12"`
	Age           int64  `bson:"age,omitempty" json:"age,omitempty" validate:"required,min=0,max=120"`
}

func (user *User) ChangeName(newName string) error {
	user.Name = newName
	return user.Update()

}

func (user *User) ChangeAge(newAge int64) error {
	user.Age = newAge
	return user.Update()
}

func (user *User) Update() error {
	validation, err := user.Validate(user)

	if validation {
		return nil
	}

	return err
}

func NewUser(name string, age int64) (User, error) {
	user := User{
		Id:   uuid.New().String(),
		Name: name,
		Age:  age,
	}

	validation, err := user.Validate(user)

	if validation {
		return user, nil
	}

	return user, err
}
