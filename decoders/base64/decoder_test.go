package main

import (
	"io"
	"strings"
	"testing"
)

func TestDecide(t *testing.T) {
	// Arrange
	reader := strings.NewReader("SGVsbG8gd29ybGQK")
	expected := "Hello world\n"

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
