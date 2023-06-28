package mongorepo

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestUpdate(t *testing.T) {
	const name = "NAME"
	const new_name = "NEW NAME"

	me := newMockEntity(name)

	_, err := mockRepo.Create(me)

	if err != nil {
		t.Error(err)
	}

	err = mockRepo.Update(me.ID, bson.M{
		"name": new_name,
	})

	if err != nil {
		t.Error(err)
	}

	foundMe, err := mockRepo.FindOne(me.ID)

	if err != nil {
		t.Error(err)
	}

	if foundMe.ID != me.ID {
		t.Errorf("Wrong ID, Expect: %s, Got: %s", me.ID, foundMe.ID)
	}

	if foundMe.Name != new_name {
		t.Errorf("Wrong Name, Expect: %s, Got: %s", new_name, me.Name)
	}

}
