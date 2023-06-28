package mongorepo

import (
	"go.mongodb.org/mongo-driver/mongo/options"
)

func optionsFunc(params Params) *options.FindOptions {
	opts := options.Find()

	if params.Limit != 0 {
		opts = opts.SetLimit(int64(params.Limit))
	}

	if params.Page != 0 {
		opts = opts.SetSkip(int64((params.Page - 1) * params.Limit))
	}

	return opts
}
