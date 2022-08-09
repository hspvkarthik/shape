package shape

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSquare(t *testing.T) {
	t.Run("check initialization of square", func(t *testing.T) {
		assert.IsType(t, Square{}, NewSquare(2.5))
	})

	t.Run("check initialization of square with negative length", func(t *testing.T) {
		assert.Panics(t, func() {
			NewSquare(-5.0)
		})
	})
}

func TestSquareFindPerimeter(t *testing.T) {
	t.Run("check perimeter with positive length", func(t *testing.T) {
		assert.Equal(t, 6.4, NewSquare(1.6).FindPerimeter())
	})

	t.Run("check perimeter with negative length", func(t *testing.T) {
		assert.Panics(t, func() {
			NewSquare(-1.6).FindPerimeter()
		})
	})
}

func TestSquareArea(t *testing.T) {
	t.Run("check area with positive length", func(t *testing.T) {
		assert.Equal(t, 6.25, NewSquare(2.5).FindArea())
	})

	t.Run("check area with negative length", func(t *testing.T) {
		assert.Panics(t, func() {
			NewSquare(-2.5).FindArea()
		})
	})
}
