package users

import (
	infra "mongodb.com/users/infra"
)

func DeleteUserUseCase(id string) error {
	err := infra.UserRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil

}
