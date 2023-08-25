package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	pass, hashed := GeneratePassword()
	assert.NotEmpty(t, pass, "Password should not be empty")
	assert.True(t, len(pass) == 20, "Password should be 20 characters long")
	assert.NotEmpty(t, hashed, "Hashed password should not be empty")
	assert.NotEqual(t, pass, hashed, "Password and hashed password should not be equal")

	assert.True(t, VerifyPassword(hashed, pass), "Password and hashed password should match")
}
