package crypt

import "testing"

func TestSha1Encode(t *testing.T) {
	rawData := []byte("hello world~")

	encryptedData := Sha1Encode(rawData)

	if encryptedData != "8e5bbd6d6430df419aca10716fc9f1c9fd8199fd" {
		t.Fatalf("sha1 encode failed: %s\n", encryptedData)
	}
}
