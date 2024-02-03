package runtime

import (
	t "github.com/patrickgombert/gisp/types"
)

type Environment struct {
	inNamespace t.Symbol
	functions   map[t.Symbol]t.Function
}
