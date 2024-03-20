package main

import (
	"io"
	"strings"
	"testing"
)

func TestDecodeBase64(t *testing.T) {
	// Arrange
	reader := strings.NewReader("546573742048657820456e636f64696e67")
	expected := "Test Hex Encoding"

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
