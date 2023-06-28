package users

import (
	domain "mongodb.com/users/domain/entities"
	infra "mongodb.com/users/infra"
)

func CreateUserUseCase(input domain.User) (domain.User, error) {
	createdUser, err := infra.UserRepository.Create(input)

	if err != nil {
		return input, err
	}

	return createdUser, nil

}
