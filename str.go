package random

import (
	"math/rand"
	"time"
)

const charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func Seed() {
	rand.Seed(time.Now().UTC().UnixNano() + 1337)
}

func String(n int) string {
	var b = make([]byte, n)
	for i := range b {
		b[i] = charSet[rand.Intn(len(charSet))]
	}
	return string(b)
}
