package runtime

import (
	"testing"

	ds "github.com/patrickgombert/gisp/datastructures"
	p "github.com/patrickgombert/gisp/parser"
	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {
	// TODO: always returning int64 is not great
	assert.Equal(t, int64(3), evalStr("(+ 1 2)"))
}

func TestRecursiveEval(t *testing.T) {
	assert.Equal(t, int64(10), evalStr("(+ (+ 1 2) (+ 3 4))"))
}

func evalStr(s string) interface{} {
	result, _ := p.FromString(s)
	return Eval(result.(*ds.List), DefaultEnvironment())
}
