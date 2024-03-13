package main

import (
	"encoding/base64"
	"io"
)

func Decode(reader io.Reader) io.Reader {
	return base64.NewDecoder(base64.StdEncoding, reader)
}
