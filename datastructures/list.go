package datastructures

type List struct {
	first any
	rest  *List
	count int
}

var EMPTY_LIST = &List{
	first: nil,
	rest:  nil,
	count: 0,
}

func NewList(objs ...any) *List {
	l := EMPTY_LIST
	for _, obj := range objs {
		l = l.Cons(obj).(*List)
	}
	return l
}

func (l *List) First() any {
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

func (l *List) Cons(item any) Collection {
	return &List{
		first: item,
		rest:  l,
		count: l.count + 1,
	}
}

func (l *List) Seq() Seq {
	return l
}
