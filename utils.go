package goutils

import (
	"math/rand"
	"time"
)

// NonceStr Get nonce string
func NonceStr(n int) string {
	s := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
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
