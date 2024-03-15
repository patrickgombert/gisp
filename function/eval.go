package function

import (
	"fmt"

	ds "github.com/patrickgombert/gisp/datastructures"
	t "github.com/patrickgombert/gisp/types"
)

func apply(sym t.Symbol, env *t.Environment, objs ...any) any {
	lookupSym := sym
	if sym.Namespace() == "" {
		lookupSym = t.NamespaceSymbol(env.CurrentNamespace().Name(), sym.Name())
	}

	f, exists := env.Resolve(lookupSym)
	if !exists {
		panic(fmt.Sprintf("%s is not bound", lookupSym.Format()))
	}
	return f.(Function).Apply(objs...)
}

func Eval(program *ds.List, env *t.Environment) any {
	f := program.First()
	arg := program.Rest()
	args := make([]any, 0)
	for arg.First() != nil {
		currentArg := arg.First()
		if l, ok := currentArg.(*ds.List); ok {
			args = append(args, Eval(l, env))
		} else if sym, ok := currentArg.(t.Symbol); ok {
			if val, exists := env.Resolve(sym); exists {
				args = append(args, val)
			} else {
				sym := t.NamespaceSymbol(env.CurrentNamespace().Name(), sym.Name())
				if val, exists := env.Resolve(sym); exists {
					args = append(args, val)
				} else {
					panic(fmt.Sprintf("symbol %s not bound", sym.Format()))
				}
			}
		} else {
			args = append(args, currentArg)
		}
		arg = arg.Rest()
	}

	if sym, ok := f.(t.Symbol); ok {
		return apply(sym, env, args...)
	} else {
		return f.(Function).Apply(args...)
	}
}
