package users

import (
	"go.mongodb.org/mongo-driver/mongo"
	mongorepo "mongodb.com/common/infra/mongo"
	domain "mongodb.com/users/domain/entities"
)

type UserRepositoryMethods struct{}

type UserRepositoryInterface interface {
	SampleMethod(id string)
}

var UserRepository *UserRepositoryStruct

type UserRepositoryStruct struct {
	mongorepo.MongoRepository[domain.User]
	UserRepositoryInterface
}

func (repo *UserRepositoryMethods) SampleMethod(id string) {

}

func NewUserRepository(database *mongo.Database, collectionName string) {
	collection := database.Collection(collectionName)

	repo := &UserRepositoryStruct{
		mongorepo.MongoRepository[domain.User]{
			Collection: collection,
			SearchableFields: []string{
				"name",
			},
		},
		&UserRepositoryMethods{},
	}

	UserRepository = repo
}
