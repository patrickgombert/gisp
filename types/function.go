package types

type Function interface {
	Apply(objs ...any) any
}
