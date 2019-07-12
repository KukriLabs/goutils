package stringutils

import (
	"math/rand"
	"time"
)

const alphaNumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = alphaNumeric[rand.Int63()%int64(len(alphaNumeric))]
	}
	return string(b)
}
