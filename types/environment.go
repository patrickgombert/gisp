package types

type Environment struct {
	inNamespace Symbol
	definitions map[Symbol]any
	scope       *Environment
}

func NewEnvironment(namespace Symbol) *Environment {
	return &Environment{
		inNamespace: namespace,
		definitions: make(map[Symbol]any),
		scope:       nil,
	}
}

func (e *Environment) CurrentNamespace() Symbol {
	return e.inNamespace
}

func (e *Environment) InNamespace(ns Symbol) {
	e.inNamespace = ns
}

func (e *Environment) Bind(s Symbol, boundTo any) {
	e.definitions[s] = boundTo
}

func (e *Environment) Resolve(s Symbol) (any, bool) {
	for e != nil {
		resolved, exists := e.definitions[s]
		if exists {
			return resolved, true
		}
		e = e.ParentScope()
	}
	return nil, false
}

func (e *Environment) ParentScope() *Environment {
	return e.scope
}

func ScopedTo(other *Environment) *Environment {
	return &Environment{
		inNamespace: other.inNamespace,
		definitions: make(map[Symbol]any),
		scope:       other,
	}
}
