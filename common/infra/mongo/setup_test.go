package mongorepo

import (
	"context"
	"os"
	"testing"

	"github.com/google/uuid"
	entity "mongodb.com/common/domain/entity"
)

type mockRepositoryMethods struct{}

type mockEntity struct {
	entity.Entity `bson:"-" json:"-"`
	ID            string `bson:"_id,omitempty" json:"id,omitempty"`
	Name          string `bson:"name,omitempty" json:"name,omitempty" validate:"required,min=3,max=12"`
}

type mockRepositoryInterface[T interface{}] interface {
	FindByName(name string) T
}

type mockRepositoryStruct struct {
	MongoRepository[mockEntity]
	mockRepositoryInterface[mockEntity]
}

func newMockEntity(name string) mockEntity {
	return mockEntity{
		ID:   uuid.New().String(),
		Name: name,
	}
}

func (repo *mockRepositoryMethods) FindByName(name string) mockEntity {

	return mockEntity{
		Name: "Test",
	}
}

var mockRepo *mockRepositoryStruct

func TestMain(m *testing.M) {
	collection, client, _ := MongoTestsPreparation()

	repo := &mockRepositoryStruct{
		MongoRepository[mockEntity]{
			Collection: collection,
			SearchableFields: []string{
				"name",
			},
		},
		&mockRepositoryMethods{},
	}

	mockRepo = repo

	code := m.Run()

	collection.Database().Drop(context.TODO())

	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}

	os.Exit(code)
}
