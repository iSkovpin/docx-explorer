package random

import (
	"math/rand"
	"time"
)

func String(n int) string {
	runes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	return StringCustom(n, runes)
}

func StringCustom(n int, runes []rune) string {
	rand.Seed(time.Now().UnixNano())
	var letters = runes

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
