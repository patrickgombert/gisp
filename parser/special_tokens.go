package parser

const OPEN_LIST = "("
const OPEN_LIST_RUNE = '('
const CLOSE_LIST = ")"
const CLOSE_LIST_RUNE = ')'

var CollectionOpeningLiterals = map[rune]rune{OPEN_LIST_RUNE: OPEN_LIST_RUNE}
var CollectionClosingLiterals = map[rune]rune{CLOSE_LIST_RUNE: CLOSE_LIST_RUNE}
