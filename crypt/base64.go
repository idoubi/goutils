package crypt

import "encoding/base64"

// Base64Encode base64加密
func Base64Encode(rawData []byte) string {
	return base64.StdEncoding.EncodeToString(rawData)
}

// Base64Decode base64解密
func Base64Decode(data string) ([]byte, error) {
	rawData, err := base64.StdEncoding.DecodeString(data)

	return rawData, err
}
