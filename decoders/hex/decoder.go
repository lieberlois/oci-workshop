package main

import (
	"encoding/hex"
	"io"
)

func Decode(reader io.Reader) io.Reader {
	return hex.NewDecoder(reader)
}
