package main

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestDecodeBase64(t *testing.T) {
	// Arrange
	expected := "Hello world"
	reader := strings.NewReader(fmt.Sprintf("%s%s", PREFIX, expected))

	// Act
	decoded := Decode(reader)
	result, err := io.ReadAll(decoded)

	// Assert
	if err != nil {
		t.Fatalf("decoding failed %v", err)
	}

	if string(result) != expected {
		t.Fatalf("expected %s, got %s", string(expected), string(result))
	}
}
