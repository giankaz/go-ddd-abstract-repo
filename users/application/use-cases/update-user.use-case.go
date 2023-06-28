package users

import (
	"go.mongodb.org/mongo-driver/bson"
	domain "mongodb.com/users/domain/entities"
	infra "mongodb.com/users/infra"
)

func UpdateUserUseCase(id string, input domain.User) error {
	update := bson.M{}

	if input.Name != "" {
		update["name"] = input.Name
	}

	if input.Age != 0 {
		update["age"] = input.Age
	}

	err := infra.UserRepository.Update(id, update)

	if err != nil {
		return err
	}

	return nil

}
