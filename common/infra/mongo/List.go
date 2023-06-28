package mongorepo

import (
	"context"
	"encoding/json"
	"errors"
	"math"
	"net/url"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	common "mongodb.com/common/application"
)

func (repo *MongoRepository[T]) List(params Params) ListResponse[T] {
	filter := bson.D{}
	var opts *options.FindOptions
	messages := ""

	if len(params.Filter) > 0 || params.Search != "" {
		filter, messages, _ = filterFunc(params, repo)

	}

	count, err := repo.Collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		count = 0
	}

	if params.Limit != 0 || params.Page != 0 {
		opts = optionsFunc(params)

	}

	var pages int64 = 0

	if params.Limit > 0 {
		pages = int64(math.Ceil(float64(count) / float64(params.Limit)))
	}

	if params.Page < 1 {
		params.Page = 1
	} else if int64(params.Page) > pages {
		params.Page = pages
	}

	cursor, err := repo.Collection.Find(context.TODO(), filter, opts)

	if err != nil {
		common.ErrorLog.Println(err)
	}

	defer func() {
		err := cursor.Close(context.TODO())

		if err != nil {
			common.ErrorLog.Println(err)
		}
	}()

	if err == mongo.ErrNoDocuments {
		common.ErrorLog.Println("Document not found")
	}

	results := []T{}

	parseError := cursor.All(context.TODO(), &results)

	if parseError != nil {
		common.ErrorLog.Println(parseError)
	}

	if len(results) <= 0 {
		results = []T{}
	}

	return ListResponse[T]{
		Items:    results,
		Total:    count,
		Page:     int64(params.Page),
		LastPage: pages,
		Limit:    int64(params.Limit),
		Error:    messages,
	}

}

func (repo *MongoRepository[T]) QueryParser(query url.Values) (Params, error) {

	filter := query.Get("filter")

	parsedFilter := []Filter{}

	if filter != "" {
		err := json.Unmarshal([]byte(filter), &parsedFilter)
		if err != nil {
			common.ErrorLog.Println(err)
			return Params{}, errors.New("check your filter for wrong information")

		}
	}

	limit, err := strconv.Atoi(query.Get("limit"))

	if err != nil || limit <= 0 {
		limit = DEFAULT_LIMIT
	}

	offset, err := strconv.Atoi(query.Get("offset"))

	if err != nil || offset <= 0 {
		offset = DEFAULT_LIMIT
	}

	page := 1

	queryPage, err := strconv.Atoi(query.Get("page"))

	if err == nil || queryPage <= 0 {
		page = queryPage
	}

	params := Params{
		Search: query.Get("search"),
		Limit:  int64(limit),
		Page:   int64(page),
		Filter: parsedFilter,
	}

	return params, nil

}
