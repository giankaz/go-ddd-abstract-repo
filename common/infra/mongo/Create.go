package mongorepo

import (
	"context"
)

func (repo *MongoRepository[T]) Create(data interface{}) (T, error) {
	var returnT T

	result, err := repo.Collection.InsertOne(context.TODO(), data)

	if err != nil {
		return returnT, err
	}

	insertedID := result.InsertedID.(string)

	createdData, err := repo.FindOne(insertedID)

	if err != nil {
		return returnT, err
	}

	return createdData, nil

}
