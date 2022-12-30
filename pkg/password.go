package pkg

import (
	"math/rand"
	"time"
)

const (
	lowerLetters   = "abcdefghijklmnopqrstuvwxyz"
	capitalLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberLetters  = "0123456789"
	specialLetters = "!@#$%^&*()"
)

// GeneratePassword generates a password with specific length
func GeneratePassword(length int, capital, special, number bool) (result string) {
	rand.Seed(time.Now().UnixNano())

	if special {
		length--
		c := specialLetters[rand.Intn(len(specialLetters))]
		result = result + string(c)
	}
	if number {
		length--
		c := numberLetters[rand.Intn(len(numberLetters))]
		result = result + string(c)
	}
	if capital {
		length--
		c := capitalLetters[rand.Intn(len(capitalLetters))]
		result = result + string(c)
	}

	if length <= 0 {
		length = 1
	}
	for i := 0; i < length; i++ {
		c := lowerLetters[rand.Intn(len(lowerLetters))]
		result = result + string(c)
	}
	return
}
