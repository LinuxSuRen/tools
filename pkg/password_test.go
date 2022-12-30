package pkg

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	tests := []struct {
		name         string
		length       int
		expectLength int
		kind         int
	}{{
		name:         "only lower case",
		length:       10,
		expectLength: 10,
		kind:         lower,
	}, {
		name:         "have capital case",
		length:       8,
		expectLength: 8,
		kind:         lower + capital,
	}, {
		name:         "have number",
		length:       9,
		expectLength: 9,
		kind:         lower + number,
	}, {
		name:         "have special",
		length:       12,
		expectLength: 12,
		kind:         lower + special,
	}, {
		name:         "have capital and number",
		length:       13,
		expectLength: 13,
		kind:         lower + capital + number,
	}, {
		name:         "have special and number",
		length:       1,
		expectLength: 3,
		kind:         lower + special + number,
	}, {
		name:         "have capital and special",
		length:       1,
		expectLength: 3,
		kind:         lower + capital + special,
	}, {
		name:         "have capital, number and special",
		length:       1,
		expectLength: 4,
		kind:         lower + capital + number + special,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			check(t, tt.kind, tt.length, tt.expectLength)
		})
	}
}

func check(t *testing.T, kind int, expectLen, expectLength int) {
	result := GeneratePassword(expectLen, kind&capital == capital, kind&special == special, kind&number == number)
	assert.Equal(t, expectLength, len(result))
	assert.Equal(t, kind&lower == lower, strings.ContainsAny(result, lowerLetters), "lower case check error")
	assert.Equal(t, kind&capital == capital, strings.ContainsAny(result, capitalLetters), "capital case check error")
	assert.Equal(t, kind&special == special, strings.ContainsAny(result, specialLetters), "special case check error")
	assert.Equal(t, kind&number == number, strings.ContainsAny(result, numberLetters), "number check error")
}

var (
	lower   = 01
	capital = lower << 1
	special = capital << 1
	number  = special << 1
)
