package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ObjectIdToString(value interface{}) string {
	objectId, ok := value.(primitive.ObjectID)
	if ok {
		return objectId.Hex()
	}
	return ""
}

func StringToObjectId(value string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(value)
}
