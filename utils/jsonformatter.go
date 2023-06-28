package utils

import "encoding/json"

func JsonFormatter(something interface{}) []byte {
	jsonData, err := json.MarshalIndent(something, "", "   ")

	if err != nil {
		panic(err)
	}

	return jsonData
}
