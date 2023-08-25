package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAll(t *testing.T) {
	jwt, err := GenerateJWT("test", "secret")
	require.NoError(t, err, "Error should be nil")
	require.NotEmpty(t, jwt, "JWT should not be empty")

	_, err = ValidateJWT("", "secret")
	require.Error(t, err, "Error should not be nil")

	user, err := ValidateJWT("Bearer "+jwt, "secret")
	require.NoError(t, err, "Error should be nil")
	require.Equal(t, "test", user, "User should be equal")
}
