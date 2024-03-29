package parser

import (
	"errors"
	"strings"
)

type tokenIterator struct {
	init      bool
	token     string
	remaining string
}

func NewTokenIterator(input string) *tokenIterator {
	return &tokenIterator{
		init:      false,
		token:     "",
		remaining: strings.TrimSpace(input),
	}
}

func (iter *tokenIterator) hasNext() bool {
	if !iter.init {
		iter.init = true
	}
	if iter.remaining == "" {
		return false
	}

	var sb strings.Builder
	remainingSliceStart := 0
	openingQuote := false
	for _, char := range iter.remaining {
		if char == '\n' && !openingQuote {
			continue
		}

		if !openingQuote {
			if char == ' ' {
				break
			}

			_, isOpeningLiteral := CollectionOpeningLiterals[char]
			if isOpeningLiteral {
				sb.WriteRune(char)
				remainingSliceStart++
				break
			}

			_, isClosingLiteral := CollectionClosingLiterals[char]
			if isClosingLiteral {
				sb.WriteRune(char)
				remainingSliceStart++
				break
			}
		}

		if char == '"' {
			if openingQuote {
				sb.WriteRune(char)
				remainingSliceStart++
				break
			} else {
				openingQuote = true
			}
		}

		sb.WriteRune(char)
		remainingSliceStart++
	}

	iter.token = sb.String()
	if len(iter.token) > 1 {
		lastCharacter := iter.token[len(iter.token)-1:]
		if lastCharacter == CLOSE_LIST {
			iter.token = iter.token[:len(iter.token)-1]
			remainingSliceStart--
		}
	}
	iter.remaining = strings.TrimSpace(iter.remaining[remainingSliceStart:])

	return true
}

func (iter *tokenIterator) getToken() (string, error) {
	if !iter.init {
		return "", errors.New("getToken called before hasNext")
	}
	return iter.token, nil
}
