package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTokenCalledBeforeHasNext(t *testing.T) {
	iter := NewTokenIterator("")
	_, err := iter.getToken()
	assert.Error(t, err)
}

func TestEmptyString(t *testing.T) {
	AssertIteration(t, "")
}

func TestSingleToken(t *testing.T) {
	AssertIteration(t, "42", "42")
}

func TestMultipleToken(t *testing.T) {
	AssertIteration(t, "42 24", "42", "24")
}

func TestListOpeningLiteral(t *testing.T) {
	AssertIteration(t, "(", "(")
	AssertIteration(t, "(+ 42", "(", "+", "42")
	AssertIteration(t, "(+ (", "(", "+", "(")
}

func TestListClosingLiteral(t *testing.T) {
	AssertIteration(t, ")", ")")
	AssertIteration(t, "42 24)", "42", "24", ")")
	AssertIteration(t, "+ 42) 24)", "+", "42", ")", "24", ")")
}

func TestRepeatedListLiteral(t *testing.T) {
	AssertIteration(t, "(())", "(", "(", ")", ")")
	AssertIteration(t, "(+ (- 1 (* 2 3)) 3)", "(", "+", "(", "-", "1", "(", "*", "2", "3", ")", ")", "3", ")")
}

func TestExtraEmptySpace(t *testing.T) {
	AssertIteration(t, "  (  1  )  ", "(", "1", ")")
}

func TestStringLiteral(t *testing.T) {
	AssertIteration(t, "\"", "\"")
	AssertIteration(t, "\"\"", "\"\"")
	AssertIteration(t, "\"foo\"", "\"foo\"")
	AssertIteration(t, "\"foo bar\"", "\"foo bar\"")
	AssertIteration(t, "(\"foo ) ( bar\")", "(", "\"foo ) ( bar\"", ")")
}

func AssertIteration(t *testing.T, input string, tokens ...string) {
	iter := NewTokenIterator(input)
	for _, expectedToken := range tokens {
		assert.True(t, iter.hasNext())
		parsedToken, _ := iter.getToken()
		assert.Equal(t, expectedToken, parsedToken)
	}
	assert.False(t, iter.hasNext())
}
