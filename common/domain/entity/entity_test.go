package common

import "testing"

type mockEntity struct {
	Entity `bson:"-" json:"-"`
	ID     string
	name   string
}

func TestEntity(t *testing.T) {
	newEntity := mockEntity{
		name: "entity",
		ID:   "fake-id",
	}

	if newEntity.ID != "fake-id" {
		t.Error("wrong id was generated")
	}
}
