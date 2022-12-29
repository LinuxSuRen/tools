package pkg

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	result := GeneratePassword(10, true, true)
	fmt.Println(result)
	assert.Equal(t, 10, len(result))
}
