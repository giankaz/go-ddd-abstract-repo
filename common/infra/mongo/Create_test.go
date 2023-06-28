package mongorepo

import (
	"testing"
)

func TestCreate(t *testing.T) {

	me := newMockEntity("NAME")

	created, err := mockRepo.Create(me)

	if err != nil {
		t.Error(err)
	}

	if created.ID != me.ID {
		t.Errorf("Wrong ID, Expect: %s, Got: %s", me.ID, created.ID)
	}

	if created.Name != "NAME" {
		t.Errorf("Wrong Name, Expect: NAME, Got: %s", me.Name)
	}
}
