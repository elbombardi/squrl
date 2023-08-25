package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	logger := NewLogger(MockConfig())
	assert.NotNil(t, logger, "Logger should not be nil")
}
