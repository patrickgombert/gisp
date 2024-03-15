package function

import (
	"testing"

	ds "github.com/patrickgombert/gisp/datastructures"
	p "github.com/patrickgombert/gisp/parser"
	t "github.com/patrickgombert/gisp/types"
	"github.com/stretchr/testify/assert"
)

func TestEval(t *testing.T) {
	assert.Equal(t, 3, evalStr("(+ 1 2)"))
}

func TestRecursiveEval(t *testing.T) {
	assert.Equal(t, 10, evalStr("(+ (+ 1 2) (+ 3 4))"))
}

func evalStr(s string) any {
	result, _ := p.FromString(s)
	// We have to bootstrap an env because of a circular dependency with runtime
	env := t.NewEnvironment(t.NameSymbol("eval-test"))
	env.Bind(t.NamespaceSymbol("eval-test", "+"), &PlusFn{})
	return Eval(result.(*ds.List), env)
}

type PlusFn struct{}

func (f *PlusFn) Apply(objs ...any) any {
	total := 0
	for _, o := range objs {
		total = total + o.(int)
	}
	return total
}
