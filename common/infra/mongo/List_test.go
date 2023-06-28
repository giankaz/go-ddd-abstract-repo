package mongorepo

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestList(t *testing.T) {
	const name = "NAME"
	const name2 = "NAME2"

	mockRepo.Collection.DeleteMany(context.TODO(), bson.D{})

	me := newMockEntity(name)
	me2 := newMockEntity(name2)

	_, err := mockRepo.Create(me)

	if err != nil {
		t.Error(err)
	}

	_, err = mockRepo.Create(me2)

	if err != nil {
		t.Error(err)
	}

	baseResult := mockRepo.List(Params{
		Search: "",
		Filter: []Filter{},
		Limit:  10,
		Page:   1,
	})

	if len(baseResult.Items) != 2 {
		t.Errorf("Expected length of 2, instead got %v", len(baseResult.Items))
	}

	searchResult := mockRepo.List(Params{
		Search: name2,
		Filter: []Filter{},
		Limit:  10,
		Page:   1,
	})

	if searchResult.Items[0] != me2 {
		t.Error("Expected search result to equals mock entity 2")
	}

	filterResult := mockRepo.List(Params{
		Search: "",
		Filter: []Filter{
			{
				Column:      "name",
				Value:    name2,
				Type:     "string",
				Operator: EQUAL,
			},
		},
		Limit: 10,
		Page:  1,
	})

	if filterResult.Items[0] != me2 {
		t.Error("Expected filter result to equals mock entity 2")
	}

	limitAndPageResult := mockRepo.List(Params{
		Search: "",
		Filter: []Filter{},
		Limit:  1,
		Page:   2,
	})

	if limitAndPageResult.Items[0] != me2 {
		t.Error("Expected limit and page result to equals mock entity 2")
	}

}
