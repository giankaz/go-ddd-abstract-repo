package mongorepo

import (
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestDelete(t *testing.T) {
	const name = "NAME"

	me := newMockEntity(name)

	_, err := mockRepo.Create(me)

	if err != nil {
		t.Error(err)
	}

	err = mockRepo.Delete(me.ID)

	if err != nil {
		t.Error(err)
	}

	_, err = mockRepo.FindOne(me.ID)

	fmt.Printf("%v", err == mongo.ErrNoDocuments)

	if err != mongo.ErrNoDocuments {
		t.Error("Found when deleted was expected")
	}

}
