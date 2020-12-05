package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax(t *testing.T) {
	assert.Equal(t, 42, Max(3, 42))
	assert.Equal(t, 42, Max(42, 3))
}
