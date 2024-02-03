package parser

type stack [][]interface{}

func newStack() stack {
	return make(stack, 0)
}

func (s stack) Push(item []interface{}) stack {
	return append(s, item)
}

func (s stack) Pop() (stack, []interface{}) {
	l := len(s)
	if l == 0 {
		return nil, nil
	}
	return s[:l-1], s[l-1]
}
