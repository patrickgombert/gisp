package parser

import (
	"testing"

	ds "github.com/patrickgombert/gisp/datastructures"
	"github.com/patrickgombert/gisp/types"
	"github.com/stretchr/testify/assert"
)

func TestReadStringValue(t *testing.T) {
	val, _ := FromString("\"foo\"")
	assert.Equal(t, val, "foo")
}

func TestReadIntegerValue(t *testing.T) {
	val, _ := FromString("42")
	assert.Equal(t, val, 42)
}

func TestReadFloatValue(t *testing.T) {
	val, _ := FromString("42.24")
	assert.Equal(t, val, 42.24)
}

func TestReadSymbol(t *testing.T) {
	val, _ := FromString("symbol")
	assert.Equal(t, "", val.(types.Symbol).Namespace())
	assert.Equal(t, "symbol", val.(types.Symbol).Name())

	val, _ = FromString("sym/bol")
	assert.Equal(t, "sym", val.(types.Symbol).Namespace())
	assert.Equal(t, "bol", val.(types.Symbol).Name())

	_, err := FromString("sy/m/bol")
	assert.Error(t, err)
}

func TestReadSingleList(t *testing.T) {
	val, _ := FromString("(+ 1 2)")
	l := val.(*ds.List)

	assert.Equal(t, 3, l.Count())
	sym := l.First().(types.Symbol)
	assert.Equal(t, "+", sym.Name())
	arg1 := l.Rest().First()
	assert.Equal(t, 1, arg1)
	arg2 := l.Rest().Rest().First()
	assert.Equal(t, 2, arg2)
}

func TestReadNestedList(t *testing.T) {
	val, _ := FromString("(+ (- 1 (* 2 3)) 3)")
	l := val.(*ds.List)

	assert.Equal(t, 3, l.Count())
	sym := l.First().(types.Symbol)
	assert.Equal(t, "+", sym.Name())

	innerList1 := l.Rest().First().(*ds.List)
	assert.Equal(t, 3, innerList1.Count())
	sym = innerList1.First().(types.Symbol)
	assert.Equal(t, "-", sym.Name())
	arg1 := innerList1.Rest().First()
	assert.Equal(t, 1, arg1)
	innerInnerList1 := innerList1.Rest().Rest().First().(*ds.List)
	assert.Equal(t, 3, innerInnerList1.Count())
	sym = innerInnerList1.First().(types.Symbol)
	assert.Equal(t, "*", sym.Name())
	assert.Equal(t, 2, innerInnerList1.Rest().First())
	assert.Equal(t, 3, innerInnerList1.Rest().Rest().First())

	assert.Equal(t, 3, l.Rest().Rest().First())
}

func TestUnmatchedParanthesis(t *testing.T) {
	_, err := FromString("(+ 1 2")
	assert.Error(t, err)

	_, err = FromString("(+ (- 1 2)")
	assert.Error(t, err)
}
