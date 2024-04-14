package main

import (
	"encoding/json"
	"io"
	"strings"
	// "github.com/goccy/go-json"
)

type Data struct {
	Value string `json:"value"`
}

func Decode(reader io.Reader) io.Reader {
	var data Data
	if err := json.NewDecoder(reader).Decode(&data); err != nil {
		panic(err)
	}

	return strings.NewReader(data.Value)
}
