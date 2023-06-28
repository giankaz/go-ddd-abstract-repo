package mongorepo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	STRING         = "string"
	NUMBER         = "number"
	OPERATOR       = "operator"
	BOOLEAN        = "boolean"
	DEFAULT_LIMIT  = 10
	DEFAULT_OFFSET = 10
)

type Filter struct {
	Column   string `json:"column" bson:"column"`
	Value    string `json:"value" bson:"value"`
	Type     string `json:"type" bson:"type"`
	Operator string `json:"operator" bson:"operator"`
}

type Params struct {
	Search string   `json:"search" bson:"search"`
	Filter []Filter `json:"filter" bson:"filter"`
	Limit  int64    `json:"limit" bson:"limit"`
	Page   int64    `json:"page" bson:"page"`
}

type ListResponse[T interface{}] struct {
	Items    []T    `json:"items"`
	Total    int64  `json:"total"`
	Page     int64  `json:"page"`
	LastPage int64  `json:"last_page"`
	Limit    int64  `json:"limit"`
	Error    string `json:"error"`
}

type MongoRepository[T interface{}] struct {
	Collection       *mongo.Collection
	SearchableFields []string
}
