package mongorepo

import (
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	common "mongodb.com/common/application"
)

func filterFunc[T interface{}](params Params, repo *MongoRepository[T]) (primitive.D, string, error) {
	var filter bson.D
	errorsMessage := ""

	if params.Search != "" {
		for _, field := range repo.SearchableFields {

			filter = append(filter, primitive.E{Key: field, Value: bson.D{{Key: "$regex", Value: params.Search}}})
		}
	}

	if len(params.Filter) > 0 {
		orFilter := bson.A{}

		for _, value := range params.Filter {
			if value.Column == "id" || value.Column == "_id" {
				orFilter = append(orFilter, bson.D{{Key: "_id", Value: value.Value}})

			} else {
				switch value.Type {
				case STRING:
					orFilter = append(orFilter, bson.D{{Key: value.Column, Value: string(value.Value)}})
				case BOOLEAN:
					boolStr, err := strconv.ParseBool(value.Value)

					if err != nil {
						errorsMessage += "/ " + err.Error()
						common.ErrorLog.Println("error parsing string to boolean")
					}

					orFilter = append(orFilter, bson.D{{Key: value.Column, Value: boolStr}})
				case NUMBER:
					numberString, err := strconv.ParseFloat(value.Value, 64)

					if err != nil {
						errorsMessage += "/ " + err.Error()
						common.ErrorLog.Println("error parsing string to int")
					}

					orFilter = append(orFilter, bson.D{{Key: value.Column, Value: numberString}})
				case OPERATOR:
					orFilter, errorsMessage = operator(orFilter, value, errorsMessage)
				default:
					errorsMessage += "/ missing or wrong filter type."
					orFilter = append(orFilter, bson.D{})
				}

			}
		}

		filter = append(filter, primitive.E{Key: "$or", Value: orFilter})
	}

	return filter, errorsMessage, nil
}
