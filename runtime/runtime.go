package runtime

import (
	ds "github.com/patrickgombert/gisp/datastructures"
)

type Function struct {
	bindings []Symbol
	body     *ds.List
}

type Environment struct {
	inNamespace Symbol
	functions   map[Symbol]Function
}
