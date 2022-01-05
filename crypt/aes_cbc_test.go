package crypt

import (
	"testing"
)

func TestAesCbcEncrypt(t *testing.T) {
	rawData := []byte("hello world~")
	key := []byte("xDf0UIWJc8PjnCmsmTF0dXBLmFh3lpCx")

	encryptedData, err := AesCbcEncrypt(rawData, key, nil)
	if err != nil {
		t.Fatalf("aes cbc encrypt error: %v\n", err)
	}

	encryptedDataB64 := Base64Encode(encryptedData)
	if encryptedDataB64 != "z7n5rS2/qmon0VVOIFCDGA==" {
		t.Fatalf("aes cbc encrypt failed: %s\n", encryptedDataB64)
	}
}

func TestAesCbcDecrypt(t *testing.T) {
	encryptedDataB64 := "z7n5rS2/qmon0VVOIFCDGA=="
	encryptedData, _ := Base64Decode(encryptedDataB64)
	key := []byte("xDf0UIWJc8PjnCmsmTF0dXBLmFh3lpCx")

	decryptedData, err := AesCbcDecrypt(encryptedData, key, nil)
	if err != nil {
		t.Fatalf("aes cbc decrypt error: %v\n", err)
	}

	if string(decryptedData) != "hello world~" {
		t.Fatalf("aes cbc decrypt failed: %s\n", decryptedData)
	}
}
