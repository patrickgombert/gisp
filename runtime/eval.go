package runtime

import (
	ds "github.com/patrickgombert/gisp/datastructures"
	t "github.com/patrickgombert/gisp/types"
)

func apply(sym t.Symbol, objs ...any) any {
	lookupSym := sym
	if sym.Namespace() == "" {
		lookupSym = t.NamespaceSymbol(DefaultEnvironment().inNamespace.Name(), sym.Name())
	}

	f := DefaultEnvironment().functions[lookupSym]
	return f.(t.Function).Apply(objs...)
}

func Eval(program *ds.List, env *Environment) any {
	f := program.First()
	arg := program.Rest()
	args := make([]any, 0)
	for arg.First() != nil {
		currentArg := arg.First()
		if l, ok := currentArg.(*ds.List); ok {
			args = append(args, Eval(l, env))
		} else {
			args = append(args, currentArg)
		}
		arg = arg.Rest()
	}

	if sym, ok := f.(t.Symbol); ok {
		return apply(sym, args...)
	} else {
		return f.(t.Function).Apply(args...)
	}
}
