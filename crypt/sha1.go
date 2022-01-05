package crypt

import (
	"crypto/sha1"
	"fmt"
)

// Sha1Encode sha1加密
func Sha1Encode(rawData []byte) string {
	s := sha1.New()

	s.Write(rawData)

	return fmt.Sprintf("%x", s.Sum(nil))
}
