package main

import (
	"io"
	"strings"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	// Arrange
	reader := strings.NewReader("4202 ranimeS ynapmoC CD")
	expected := "DC Company Seminar 2024"

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
