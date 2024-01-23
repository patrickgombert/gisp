package datastructures

type Seq interface {
	First() interface{}
	Rest() Seq
}

type Collection interface {
	Count() int
	Cons(item interface{}) Collection
	Seq() Seq
}
