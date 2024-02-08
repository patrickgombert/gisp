package runtime

import (
	t "github.com/patrickgombert/gisp/types"
)

var DEFAULT_NAMESPACE = t.NameSymbol("core")

type Environment struct {
	inNamespace t.Symbol
	functions   map[t.Symbol]t.Function
}

type BuiltInFunction struct {
	f func(...interface{}) interface{}
}

func (bif BuiltInFunction) Apply(objs ...interface{}) interface{} {
	return bif.f(objs...)
}

var env *Environment

func DefaultEnvironment() *Environment {
	if env == nil {
		env = &Environment{
			inNamespace: DEFAULT_NAMESPACE,
			functions: map[t.Symbol]t.Function{
				t.NamespaceSymbol("core", "+"): BuiltInFunction{f: add},
			},
		}
	}

	return env
}

func add(objs ...interface{}) interface{} {
	var intSum int64 = 0
	var floatSum float64 = 0.0
	usingFloat := false
	for _, obj := range objs {
		switch obj.(type) {
		case int:
			intSum += int64(obj.(int))
			floatSum += float64(obj.(int))
		case int8:
			intSum += int64(obj.(int8))
			floatSum += float64(obj.(int8))
		case int16:
			intSum += int64(obj.(int16))
			floatSum += float64(obj.(int16))
		case int32:
			intSum += int64(obj.(int32))
			floatSum += float64(obj.(int32))
		case int64:
			intSum += obj.(int64)
			floatSum += float64(obj.(int64))
		case uint8:
			intSum += int64(obj.(uint8))
			floatSum += float64(obj.(uint8))
		case uint32:
			intSum += int64(obj.(uint32))
			floatSum += float64(obj.(uint32))
		case uint64:
			intSum += int64(obj.(uint64))
			floatSum += float64(obj.(uint64))
		case float32:
			floatSum += float64(obj.(float32))
			usingFloat = true
		case float64:
			floatSum += obj.(float64)
			usingFloat = true
		default:
			panic("cannot invoke + on non-numeric type")
		}
	}

	if usingFloat {
		return floatSum
	} else {
		return intSum
	}
}
