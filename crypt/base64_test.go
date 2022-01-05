package crypt

import (
	"testing"
)

func TestBase64Encode(t *testing.T) {
	rawData := []byte("hello world~")
	data := Base64Encode(rawData)

	if string(data) != "aGVsbG8gd29ybGR+" {
		t.Fatalf("base64 encode failed: %s\n", data)
	}
}

func TestBase64Decode(t *testing.T) {
	data := "aGVsbG8gd29ybGR+"
	rawData, err := Base64Decode(data)
	if err != nil {
		t.Fatalf("base64 decode error: %v\n", err)
	}

	if string(rawData) != "hello world~" {
		t.Fatalf("base64 decode failed: %s\n", rawData)
	}
}
