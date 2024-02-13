package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyList(t *testing.T) {
	l := NewList()
	assert.Equal(t, 0, l.Count())
	assert.Equal(t, l, l.Seq())

	newL := l.Cons("obj")
	assert.Equal(t, 1, newL.Count())
	assert.Equal(t, newL, newL.Seq())

	assert.Nil(t, l.First())
	assert.Equal(t, 0, l.Rest().(*List).Count())

	assert.Equal(t, "obj", newL.Seq().First())
	assert.Equal(t, l, newL.Seq().Rest())
}

func TestConsPrepends(t *testing.T) {
	l := NewList(3, 2)
	newL := l.Cons(1)

	assert.Equal(t, 3, newL.Count())
	assert.Equal(t, 1, newL.Seq().First())
	assert.Equal(t, 2, newL.Seq().Rest().First())
	assert.Equal(t, 3, newL.Seq().Rest().Rest().First())
}
