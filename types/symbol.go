package types

import (
	"fmt"
	"strings"
)

type Symbol struct {
	ns   string
	name string
}

func NameSymbol(n string) Symbol {
	return Symbol{
		name: n,
	}
}

func NamespaceSymbol(n1 string, n2 string) Symbol {
	return Symbol{
		ns:   n1,
		name: n2,
	}
}

func ParseSymbol(raw string) (Symbol, error) {
	parts := strings.Split(raw, "/")
	switch len(parts) {
	case 1:
		return Symbol{
			ns:   "",
			name: parts[0],
		}, nil
	case 2:
		return Symbol{
			ns:   parts[0],
			name: parts[1],
		}, nil
	default:
		return Symbol{}, fmt.Errorf("Invalid path for symbol %s", raw)
	}
}

func (s Symbol) Namespace() string {
	return s.ns
}

func (s Symbol) Name() string {
	return s.name
}

func (s Symbol) Format() string {
	if s.ns != "" {
		return fmt.Sprintf("%s/%s", s.ns, s.name)
	} else {
		return fmt.Sprintf("%s", s.name)
	}
}
