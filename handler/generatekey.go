package handler

import (
	"math/rand"
	"sync"
	"time"
)

var (
	urlStore     = make(map[string]string)
	reverseStore = make(map[string]string)
	mutex        = sync.RWMutex{}
	rng          = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func generateShortKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 6
	shortKey := make([]byte, keyLength)

	for i := range shortKey {
		shortKey[i] = charset[rng.Intn(len(charset))]
	}
	return string(shortKey)
}