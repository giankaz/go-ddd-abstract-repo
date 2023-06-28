package users

import (
	domain "mongodb.com/users/domain/entities"
	infra "mongodb.com/users/infra"
)

func FindUserUseCase(id string) (domain.User, error) {
	user, err := infra.UserRepository.FindOne(id)

	if err != nil {
		return domain.User{}, err
	}

	return user, nil

}
