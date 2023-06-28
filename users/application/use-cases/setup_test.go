package users

import (
	"context"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	mongorepo "mongodb.com/common/infra/mongo"
	users "mongodb.com/users/infra"
)

func TestMain(m *testing.M) {
	collection, client, database := mongorepo.MongoTestsPreparation()

	users.NewUserRepository(database, "testing-users")

	collection.DeleteMany(context.TODO(), bson.D{})

	code := m.Run()

	database.Drop(context.TODO())

	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}

	os.Exit(code)
}
