package runtime

import (
	ds "github.com/patrickgombert/gisp/datastructures"
)

type Fn interface {
	Apply(objs ...interface{}) interface{}
}

type BoundFn struct {
	arguments []Symbol
	body      *ds.List
}

func (f *BoundFn) Apply(objs ...interface{}) interface{} {
	return nil
}
