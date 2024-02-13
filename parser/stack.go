package parser

type stack [][]any

func newStack() stack {
	return make(stack, 0)
}

func (s stack) Push(item []any) stack {
	return append(s, item)
}

func (s stack) Pop() (stack, []any) {
	l := len(s)
	if l == 0 {
		return nil, nil
	}
	return s[:l-1], s[l-1]
}
