package datastructures

type List struct {
	first interface{}
	rest  *List
	count int
}

var EMPTY_LIST = &List{
	first: nil,
	rest:  nil,
	count: 0,
}

func NewList(objs ...interface{}) *List {
	l := EMPTY_LIST
	for _, obj := range objs {
		l = l.Cons(obj)
	}
	return l
}

func (l *List) First() interface{} {
	return l.first
}

func (l *List) Rest() Seq {
	if l.rest == nil {
		return EMPTY_LIST
	}
	return l.rest
}

func (l *List) Count() int {
	return l.count
}

func (l *List) Cons(item interface{}) *List {
	return &List{
		first: item,
		rest:  l,
		count: l.count + 1,
	}
}

func (l *List) Seq() Seq {
	return l
}
