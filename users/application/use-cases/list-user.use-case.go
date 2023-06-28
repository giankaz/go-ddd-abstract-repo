package users

import (
	"net/url"

	mongorepo "mongodb.com/common/infra/mongo"
	domain "mongodb.com/users/domain/entities"
	infra "mongodb.com/users/infra"
)

func ListUserUseCase(query url.Values) (mongorepo.ListResponse[domain.User], error) {
	parsedQuery, err := infra.UserRepository.QueryParser(query)

	if err != nil {
		return mongorepo.ListResponse[domain.User]{}, err
	}

	users := infra.UserRepository.List(parsedQuery)

	return users, nil

}
