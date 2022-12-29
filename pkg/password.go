package pkg

import (
	"math/rand"
	"time"
)

// GeneratePassword generates a password with specific length
func GeneratePassword(length int, special, number bool) (result string) {
	rand.Seed(time.Now().UnixNano())

	charset := "abcdefghijklmnopqrstuvwxyz"
	if special {
		charset += "!@#$%^&*()"
	}
	if number {
		charset += "0123456789"
	}

	for i := 0; i < length; i++ {
		c := charset[rand.Intn(len(charset))]
		result = result + string(c)
	}
	return
}
