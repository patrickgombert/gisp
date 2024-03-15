package runtime

import (
	"testing"

	ds "github.com/patrickgombert/gisp/datastructures"
	"github.com/patrickgombert/gisp/function"
	types "github.com/patrickgombert/gisp/types"
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

func TestFn(t *testing.T) {
	args := list(symbol("a"))
	body := list(symbol("+"), 2, symbol("a"))
	result := fn(args, body).(function.Function).Apply(2)
	assert.Equal(t, 4, result)
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

func TestDef(t *testing.T) {
	def(symbol("foo"), "bar")
	val, exists := DefaultEnvironment().Resolve(symbol("core/foo").(types.Symbol))
	assert.True(t, exists)
	assert.Equal(t, "bar", val)

	def(symbol("foo/bar"), "baz")
	val, exists = DefaultEnvironment().Resolve(symbol("foo/bar").(types.Symbol))
	assert.True(t, exists)
	assert.Equal(t, "baz", val)
}
