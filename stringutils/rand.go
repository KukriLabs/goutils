package stringutils

import (
	"math/rand"
	"time"
)

func RandString(size int, inputChars string) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, size)
	for i := range b {
		b[i] = inputChars[rand.Int63()%int64(len(inputChars))]
	}
	return string(b)

}

func RandAlphaNumericString(size int) string {
	const alphaNumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	return RandString(size, alphaNumeric)
}
