package mongorepo

import (
	"testing"
)

func TestFindOne(t *testing.T) {
	const name = "NAME"

	me := newMockEntity(name)

	_, err := mockRepo.Create(me)

	if err != nil {
		t.Error(err)
	}

	foundMe, err := mockRepo.FindOne(me.ID)

	if foundMe.ID != me.ID {
		t.Errorf("Wrong ID, Expect: %s, Got: %s", me.ID, foundMe.ID)
	}

	if foundMe.Name != "NAME" {
		t.Errorf("Wrong Name, Expect: NAME, Got: %s", me.Name)
	}

}
