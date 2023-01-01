package crypt

import (
	"testing"
)

func TestAesEcbEncrypt(t *testing.T) {
	rawData := []byte("hello world~")
	key := []byte("xDf0UIWJc8PjnCmsmTF0dXBLmFh3lpCx")

	encryptedData, err := AesEcbEncrypt(rawData, key)
	if err != nil {
		t.Fatalf("aes Ecb encrypt error: %v\n", err)
	}

	encryptedDataB64 := Base64Encode(encryptedData)
	if encryptedDataB64 != "OTvgPBWp/PrzWtKhDkkqcw==" {
		t.Fatalf("aes Ecb encrypt failed: %s\n", encryptedDataB64)
	}
}

func TestAesEcbDecrypt(t *testing.T) {
	encryptedDataB64 := "OTvgPBWp/PrzWtKhDkkqcw=="
	encryptedData, _ := Base64Decode(encryptedDataB64)
	key := []byte("xDf0UIWJc8PjnCmsmTF0dXBLmFh3lpCx")

	decryptedData, err := AesEcbDecrypt(encryptedData, key)
	if err != nil {
		t.Fatalf("aes Ecb decrypt error: %v\n", err)
	}

	if string(decryptedData) != "hello world~" {
		t.Fatalf("aes Ecb decrypt failed: %s\n", decryptedData)
	}
}
