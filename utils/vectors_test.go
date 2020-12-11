package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVector2_Add(t *testing.T) {
	v1 := Vector2{X: 1, Y: 2}
	v2 := Vector2{X: 3, Y: 4}

	v3 := v1.Add(v2)

	assert.Equal(t, Vector2{X: 4, Y: 6}, v3)
}

func TestVector2_Reduce(t *testing.T) {
	assert.Equal(t, Vector2{X: 1, Y: 1}, (Vector2{X: 1, Y: 1}).Reduce())
	assert.Equal(t, Vector2{X: 1, Y: 1}, (Vector2{X: 2, Y: 2}).Reduce())
	assert.Equal(t, Vector2{X: 1, Y: 1}, (Vector2{X: 3, Y: 3}).Reduce())
	assert.Equal(t, Vector2{X: 1, Y: 1}, (Vector2{X: 4, Y: 4}).Reduce())

	assert.Equal(t, Vector2{X: 1, Y: 2}, (Vector2{X: 2, Y: 4}).Reduce())

	assert.Equal(t, Vector2{X: 2, Y: 1}, (Vector2{X: 4, Y: 2}).Reduce())

	assert.Equal(t, Vector2{X: 3, Y: 10}, (Vector2{X: 3*2*2*5, Y: 10*2*2*5}).Reduce())
}

func TestVector2_Multiply(t *testing.T) {
	v1 := Vector2{X: 3, Y: 4}

	v3 := v1.Multiply(2)

	assert.Equal(t, Vector2{X: 6, Y: 8}, v3)
}
