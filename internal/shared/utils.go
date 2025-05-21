package shared

import (
	"math/rand"
	"strings"
	"time"
)

func GenerateRandomID(prefix string, length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return prefix + string(b)
}

func SanitizeString(input string) string {
	return strings.TrimSpace(strings.ToLower(input))
}
