package main

import (
	"encoding/json"
	"io"
)

func parseJson(jsonToParse io.Reader) (interface{}, error) {
	decoder := json.NewDecoder(jsonToParse)
	var json interface{}
	err := decoder.Decode(&json)
	return json, err
}
