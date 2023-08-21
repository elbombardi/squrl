package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmailWithInvalidEmail(t *testing.T) {
	err := ValidateEmail("invalidemail")
	assert.Error(t, err, "Email validation should fail")

}

func TestValidateEmailWithValidEmail(t *testing.T) {
	err := ValidateEmail("test@example.com")
	assert.NoError(t, err, "Email validation should pass")
}

func TestValidateUsernameWithInvalidUsername(t *testing.T) {
	err := ValidateUsername("invalid username!")
	assert.Error(t, err, "Username validation should fail")
}

func TestValidateUsernameWithValidUsername(t *testing.T) {
	err := ValidateUsername("validusername")
	assert.NoError(t, err, "Username validation should pass")
}
