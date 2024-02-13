package runtime

import (
	"testing"

	ds "github.com/patrickgombert/gisp/datastructures"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	result := add(1, 1)
	assert.Equal(t, 2, result)

	result = add(1, 1.2)
	assert.Equal(t, 2.2, result)

	result = add(int8(1), int16(2), int8(3))
	assert.Equal(t, int16(6), result)
}

func TestReduce(t *testing.T) {
	result := reduce(BuiltInFunction{f: add}, list(1, 2, 3))
	assert.Equal(t, 6, result)

	result = reduce(BuiltInFunction{f: add}, 1, list(1, 2, 3))
	assert.Equal(t, 7, result)
}

func TestList(t *testing.T) {
	result := list().(*ds.List)
	assert.Equal(t, nil, result.First())

	result = list("first", "second", "third").(*ds.List)
	assert.Equal(t, "first", result.First())
	assert.Equal(t, "second", result.Rest().First())
	assert.Equal(t, "third", result.Rest().Rest().First())
}
