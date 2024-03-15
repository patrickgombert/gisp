package function

import (
	"errors"
	"fmt"

	ds "github.com/patrickgombert/gisp/datastructures"
	t "github.com/patrickgombert/gisp/types"
)

type Function interface {
	Apply(objs ...any) any
}

type fn struct {
	args []t.Symbol
	body *ds.List
	env  *t.Environment
}

func NewFn(a ds.Collection, b *ds.List, e *t.Environment) (*fn, error) {
	seq := a.Seq()
	args := []t.Symbol{}
	for seq.First() != nil {
		s, ok := seq.First().(t.Symbol)
		if !ok {
			return nil, errors.New("Function arguments must be symbols")
		}
		if s.Namespace() != "" {
			return nil, fmt.Errorf("Function argument %s must not be namespace scoped", s.Format())
		}
		args = append(args, s)
		seq = seq.Rest()
	}

	return &fn{
		args: args,
		body: b,
		env:  e,
	}, nil
}

func (f *fn) Apply(objs ...any) any {
	if len(f.args) != len(objs) {
		panic(fmt.Sprintf("function has arity of %d", len(f.args)))
	}

	for i := 0; i < len(f.args); i++ {
		f.env.Bind(f.args[i], objs[i])
	}

	return Eval(f.body, f.env)
}
