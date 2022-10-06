package goutils

import (
	"math/rand"
	"time"
)

const (
	numbers      = "0123456789"
	lowerChars   = "abcdefghijklmnopqrstuvwxyz"
	upperChars   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialChars = "!@#$%^&*"
)

// NonceStr: gen nonce str with given length
func NonceStr(n int) string {
	s := lowerChars + upperChars

	return nonceSeq(s, n)
}

// NoncePassword: gen nonce password with given length
func NoncePassword(n int) string {
	s := numbers + lowerChars + upperChars + specialChars

	return nonceSeq(s, n)
}

// gen nonce with given length
func nonceSeq(s string, n int) string {
	l := len(s)
	b := make([]byte, n)
	for i := range b {
		b[i] = s[rand.Intn(l)]
	}

	return string(b)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
