package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	ds "github.com/patrickgombert/gisp/datastructures"
	t "github.com/patrickgombert/gisp/types"
)

func FromString(input string) (interface{}, error) {
	iter := NewTokenIterator(input)
	if !iter.hasNext() {
		return nil, fmt.Errorf("Invalid input %q", input)
	}
	token, err := iter.getToken()
	if err != nil {
		return nil, err
	}

	// Special case for a single value
	if token != OPEN_LIST {
		return toValue(token)
	}

	listStack := newStack()
	currentList := make([]interface{}, 0)
	lastSeenToken := "("

	for iter.hasNext() {
		token, err = iter.getToken()
		if err != nil {
			return nil, err
		}

		if token == OPEN_LIST {
			listStack = listStack.Push(currentList)
			currentList = make([]interface{}, 0)
		} else if token == CLOSE_LIST {
			newListStack, popList := listStack.Pop()
			if popList != nil {
				listStack = newListStack
				list := toList(currentList)
				currentList = append(popList, list)
			}
		} else {
			value, valueError := toValue(token)
			if valueError != nil {
				return nil, valueError
			}
			currentList = append(currentList, value)
		}

		lastSeenToken = token
	}

	if len(listStack) != 0 || lastSeenToken != CLOSE_LIST {
		return nil, errors.New("Unclosed list")
	}

	return toList(currentList), nil
}

func toValue(token string) (interface{}, error) {
	if token[0] == '"' {
		return strings.Trim(token, "\""), nil
	} else if integer, integerErr := strconv.Atoi(token); integerErr == nil {
		return integer, nil
	} else if float, floatErr := strconv.ParseFloat(token, 64); floatErr == nil {
		return float, nil
	} else {
		return t.ParseSymbol(token)
	}
}

func toList(items []interface{}) *ds.List {
	list := ds.NewList()
	for i := len(items) - 1; i >= 0; i-- {
		list = list.Cons(items[i])
	}
	return list
}
