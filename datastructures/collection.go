package datastructures

type Seq interface {
	First() any
	Rest() Seq
}

type Collection interface {
	Count() int
	Cons(item any) Collection
	Seq() Seq
}
