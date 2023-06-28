package mongorepo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *MongoRepository[T]) Delete(id string) error {

	_, err := repo.Collection.DeleteOne(context.TODO(), bson.M{"_id": id})

	if err == mongo.ErrNoDocuments {
		return err
	}

	return nil

}
