package main

import (
	"io"
	"strings"
)

func Decode(reader io.Reader) io.Reader {
	content, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	runes := []rune(string(content))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	result := string(runes)
	return strings.NewReader(result)
}
