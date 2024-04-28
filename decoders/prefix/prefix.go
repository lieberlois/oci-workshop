package main

import (
	"io"
	"strings"
)

const PREFIX = "PREFIX___"

func Decode(reader io.Reader) io.Reader {
	data, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	truncated := strings.Replace(string(data), PREFIX, "", -1)
	return strings.NewReader(truncated)
}
