package random

import (
	"math/rand"
	"time"
)

func GenerateRandomString(length int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[int(rand.Int63()%int64(len(letters)))]
	}
	return string(b)
}
