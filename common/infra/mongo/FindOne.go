package mongorepo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *MongoRepository[T]) FindOne(id string) (T, error) {
	var result T

	err := repo.Collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		return result, err
	}

	return result, nil

}
