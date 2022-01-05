package crypt

import (
	"testing"

	"github.com/idoubi/goutils"
)

func TestAesGcmEncrypt(t *testing.T) {
	rawData := []byte("hello world~")
	associatedData := []byte("KRzM8vQbKspOHqOH")      // 长度不限，可为空
	nonceStr := []byte("xMsqrwfX8kCk")                // 12位
	key := []byte("x95IImo3VFgqLtP8hdXoniV48B9vgpVd") // 32位：aes-256；16位：aes-128

	encryptedData, err := AesGcmEncrypt(rawData, associatedData, nonceStr, key)
	if err != nil {
		t.Fatalf("aes encrypt error: %v\n", err)
	}

	encryptedDataB64 := goutils.Base64Encode(string(encryptedData))
	if encryptedDataB64 != "apoCY6iEavJMyx0HelJGVcesxxd2UlTlUuqnmQ==" {
		t.Fatalf("aes encrypt failed: %s\n", encryptedDataB64)
	}
}

func TestAesGcmDecrypt(t *testing.T) {
	encryptedDataB64 := "apoCY6iEavJMyx0HelJGVcesxxd2UlTlUuqnmQ=="

	encryptedData := []byte(goutils.Base64Decode(encryptedDataB64))
	associatedData := []byte("KRzM8vQbKspOHqOH")
	nonceStr := []byte("xMsqrwfX8kCk")
	key := []byte("x95IImo3VFgqLtP8hdXoniV48B9vgpVd")

	decryptedData, err := AesGcmDecrypt(encryptedData, associatedData, nonceStr, key)
	if err != nil {
		t.Fatalf("aes encrypt error: %v\n", err)
	}

	if string(decryptedData) != "hello world~" {
		t.Fatalf("aes decrypt failed: %s\n", decryptedData)
	}
}
