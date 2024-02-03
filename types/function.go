package types

type Function interface {
	Apply(objs ...interface{}) interface{}
}
