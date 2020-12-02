package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMustParseInt(t *testing.T) {
	assert.Equal(t, 3, MustParseInt("3"))
	assert.Equal(t, 42, MustParseInt("42"))
}
