package mongorepo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *MongoRepository[T]) Update(id string, updateData bson.M) error {

	updateObject := bson.M{"$set": updateData}

	_, err := repo.Collection.UpdateOne(context.TODO(), bson.M{"_id": id}, updateObject)

	if err == mongo.ErrNoDocuments {
		return err
	}

	return nil

}
