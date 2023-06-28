package common

import (
	"encoding/json"

	common "mongodb.com/common/application"
)

type EntityInterface interface {
	ToJSON() []byte
}

type Entity struct{}

func (e *Entity) ToJSON() []byte {
	jsonEntity, err := json.Marshal(e)

	if err != nil {
		common.ErrorLog.Println(err)
	}

	return jsonEntity
}