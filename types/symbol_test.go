package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSymbolNoNamespace(t *testing.T) {
	sym, _ := ParseSymbol("one")
	assert.Equal(t, "", sym.Namespace())
	assert.Equal(t, "one", sym.Name())
}

func TestSymbolNamespaced(t *testing.T) {
	sym, _ := ParseSymbol("one/two")
	assert.Equal(t, "one", sym.Namespace())
	assert.Equal(t, "two", sym.Name())
}

func TestSymbolMayOnlyHave2StepPath(t *testing.T) {
	_, err := ParseSymbol("one/two/three")
	assert.Error(t, err)
}
