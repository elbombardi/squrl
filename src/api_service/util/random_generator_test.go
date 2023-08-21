package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandom(t *testing.T) {
	randomString1 := GenerateRandomString(20)
	randomString2 := GenerateRandomString(20)
	assert.Equal(t, 20, len(randomString1), "Random string should be 20 characters long")
	assert.NotEqual(t, randomString1, randomString2, "Random strings should not be equal")
}
