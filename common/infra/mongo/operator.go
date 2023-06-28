package mongorepo

import (
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	CONTAINS        = "CONTAINS"
	NOT_CONTAINS    = "NOT_CONTAINS"
	EQUAL           = "EQUAL"
	NOT_EQUAL       = "NOT_EQUAL"
	IS_FILLED       = "IS_FILLED"
	IS_NOT_FILLED   = "IS_NOT_FILLED"
	HAD             = "HAD"
	NOT_HAD         = "NOT_HAD"
	IS_GREATER_THAN = "IS_GREATER_THAN"
	IS_LESS_THAN    = "IS_LESS_THAN"
)

func operator(orFilter bson.A, params Filter, errorMessages string) (bson.A, string) {
	switch params.Operator {
	case CONTAINS:
		orFilter = append(orFilter, bson.D{{Key: params.Column, Value: bson.D{{Key: "$regex", Value: params.Value}}}})
	case NOT_CONTAINS:
		orFilter = append(orFilter, bson.D{{Key: params.Column, Value: bson.D{{Key: "$not", Value: bson.D{{Key: "$regex", Value: params.Value}}}}}})
	case EQUAL:
		orFilter = append(orFilter, bson.D{{Key: params.Column, Value: params.Value}})
	case NOT_EQUAL:
		orFilter = append(orFilter, bson.D{{Key: params.Column, Value: bson.D{{Key: "$ne", Value: params.Value}}}})
	case IS_FILLED:
		orFilter = append(orFilter, bson.D{{Key: params.Column, Value: bson.D{{Key: "$ne", Value: nil}}}})
	case IS_NOT_FILLED:
		orFilter = append(orFilter, bson.D{{Key: params.Column, Value: nil}})
	case HAD:
		orFilter = append(orFilter, bson.D{{Key: params.Column, Value: bson.D{{Key: "$gt", Value: 0}}}})
	case NOT_HAD:
		orFilter = append(orFilter, bson.D{{Key: params.Column, Value: 0}})
	case IS_GREATER_THAN:
		parsedToFloat, _ := strconv.ParseFloat(params.Value, 64)
		orFilter = append(orFilter, bson.D{{Key: params.Column, Value: bson.D{{Key: "$gt", Value: parsedToFloat}}}})
	case IS_LESS_THAN:
		parsedToFloat, _ := strconv.ParseFloat(params.Value, 64)
		orFilter = append(orFilter, bson.D{{Key: params.Column, Value: bson.D{{Key: "$lt", Value: parsedToFloat}}}})
	}

	return orFilter, errorMessages
}
