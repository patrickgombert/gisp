package runtime

import (
	"fmt"

	ds "github.com/patrickgombert/gisp/datastructures"
	"github.com/patrickgombert/gisp/function"
	p "github.com/patrickgombert/gisp/parser"
	t "github.com/patrickgombert/gisp/types"
)

var DEFAULT_NAMESPACE = t.NameSymbol("core")

const MAXINT = int(^uint(0) >> 1)

type BuiltInFunction struct {
	f func(...any) any
}

func (bif BuiltInFunction) Apply(objs ...any) any {
	return bif.f(objs...)
}

var env *t.Environment

func DefaultEnvironment() *t.Environment {
	if env == nil {
		env = t.NewEnvironment(DEFAULT_NAMESPACE)
		env.Bind(t.NamespaceSymbol("core", "+"), BuiltInFunction{f: add})
		env.Bind(t.NamespaceSymbol("core", "def"), BuiltInFunction{f: def})
		env.Bind(t.NamespaceSymbol("core", "eval"), BuiltInFunction{f: eval})
		env.Bind(t.NamespaceSymbol("core", "fn"), BuiltInFunction{f: fn})
		env.Bind(t.NamespaceSymbol("core", "list"), BuiltInFunction{f: list})
		env.Bind(t.NamespaceSymbol("core", "ns"), BuiltInFunction{f: ns})
		env.Bind(t.NamespaceSymbol("core", "parse"), BuiltInFunction{f: parse})
		env.Bind(t.NamespaceSymbol("core", "print"), BuiltInFunction{f: prnt})
		env.Bind(t.NamespaceSymbol("core", "reduce"), BuiltInFunction{f: reduce})
		env.Bind(t.NamespaceSymbol("core", "symbol"), BuiltInFunction{f: symbol})
	}

	return env
}

func add(objs ...any) any {
	var intSum int64 = 0
	var floatSum float64 = 0.0
	usingFloat := false
	seenInt8 := false
	seenInt16 := false
	seenInt32 := false
	seenInt64 := false
	seenFloat64 := false

	for _, obj := range objs {
		switch obj.(type) {
		case int:
			intSum += int64(obj.(int))
			floatSum += float64(obj.(int))
		case int8:
			seenInt8 = true
			intSum += int64(obj.(int8))
			floatSum += float64(obj.(int8))
		case int16:
			seenInt16 = true
			intSum += int64(obj.(int16))
			floatSum += float64(obj.(int16))
		case int32:
			seenInt32 = true
			intSum += int64(obj.(int32))
			floatSum += float64(obj.(int32))
		case int64:
			seenInt64 = true
			intSum += obj.(int64)
			floatSum += float64(obj.(int64))
		case uint8:
			seenInt8 = true
			intSum += int64(obj.(uint8))
			floatSum += float64(obj.(uint8))
		case uint16:
			seenInt16 = true
			intSum += int64(obj.(uint16))
			floatSum += float64(obj.(uint16))
		case uint32:
			seenInt32 = true
			intSum += int64(obj.(uint32))
			floatSum += float64(obj.(uint32))
		case uint64:
			seenInt64 = true
			intSum += int64(obj.(uint64))
			floatSum += float64(obj.(uint64))
		case float32:
			floatSum += float64(obj.(float32))
			usingFloat = true
		case float64:
			seenFloat64 = true
			floatSum += obj.(float64)
			usingFloat = true
		default:
			panic("cannot invoke + on non-numeric type")
		}
	}

	if usingFloat {
		if seenFloat64 {
			return floatSum
		} else {
			return float32(floatSum)
		}
	} else {
		if !seenInt64 && !seenInt32 && !seenInt16 && !seenInt8 {
			return int(intSum)
		} else if seenInt64 || intSum > 2147483647 || intSum < -2147483648 {
			return intSum
		} else if seenInt32 || intSum > 32767 || intSum < -32768 {
			return int32(intSum)
		} else if seenInt16 || intSum > 127 || intSum < -128 {
			return int16(intSum)
		} else if seenInt8 {
			return int8(intSum)
		} else {
			return int(intSum)
		}
	}
}

func def(objs ...any) any {
	if len(objs) != 2 {
		panic("def has an arity of 2")
	}

	sym, ok := objs[0].(t.Symbol)
	if ok {
		if sym.Namespace() == "" {
			sym = t.NamespaceSymbol(DefaultEnvironment().CurrentNamespace().Name(), sym.Name())
		}
	} else if s, sok := objs[0].(string); sok {
		parsedSym, err := t.ParseSymbol(s)
		if err != nil {
			panic(err)
		}
		sym = parsedSym
	} else {
		panic("the first argument of def must be a string or symbol")
	}

	DefaultEnvironment().Bind(sym, objs[1])
	return objs[1]
}

func eval(objs ...any) any {
	if len(objs) != 1 {
		panic("eval has an arity of 1")
	}
	list := objs[0].(*ds.List)
	return function.Eval(list, DefaultEnvironment())
}

func list(objs ...any) any {
	elements := make([]any, len(objs))
	for i := len(objs) - 1; i >= 0; i-- {
		elements[len(objs)-1-i] = objs[i]
	}
	return ds.NewList(elements...)
}

func fn(objs ...any) any {
	if len(objs) != 2 {
		panic("fn has an arity of 2")
	}
	args := objs[0].(ds.Collection)
	body := objs[1].(*ds.List)

	f, err := function.NewFn(args, body, t.ScopedTo(DefaultEnvironment()))
	if err != nil {
		panic(err)
	}
	return f
}

func ns(objs ...any) any {
	namespace := t.NameSymbol(objs[0].(string))
	currentNamespace := DefaultEnvironment().CurrentNamespace()
	DefaultEnvironment().InNamespace(namespace)

	for _, obj := range objs[1:] {
		if list, ok := obj.(*ds.List); ok {
			function.Eval(list, DefaultEnvironment())
		}
	}

	DefaultEnvironment().InNamespace(currentNamespace)
	return nil
}

func parse(objs ...any) any {
	if len(objs) != 1 {
		panic("parse has an arity of 1")
	}
	list, err := p.FromString(objs[0].(string))
	if err != nil {
		panic(fmt.Sprintf("parse error %s", err))
	}
	return list
}

func prnt(objs ...any) any {
	if len(objs) > 0 {
		fmt.Printf(objs[0].(string), objs[1:]...)
	}
	return nil
}

func reduce(objs ...any) any {
	if len(objs) > 3 {
		panic("reduce has an arity of 2 or 3")
	}

	f := objs[0].(function.Function)
	var seq ds.Seq
	var acc any
	if col, ok := objs[1].(ds.Collection); ok {
		seq = col.Seq()
		acc = seq.First()
		seq = seq.Rest()
	} else {
		seq = objs[2].(ds.Collection).Seq()
		acc = objs[1]
	}

	for seq.First() != nil {
		acc = f.Apply(acc, seq.First())
		seq = seq.Rest()
	}

	return acc
}

func symbol(objs ...any) any {
	if len(objs) != 1 {
		panic("symbol has an arity of 1")
	}

	sym, err := t.ParseSymbol(objs[0].(string))
	if err != nil {
		panic(err)
	}
	return sym
}
