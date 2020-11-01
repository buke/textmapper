// generated by Textmapper; DO NOT EDIT

package tm

import (
	"strings"
	"unicode/utf8"
)

// Lexer states.
const (
	StateInitial        = 0
	StateAfterID        = 1
	StateAfterColonOrEq = 2
	StateAfterGT        = 3
)

// Lexer uses a generated DFA to scan through a utf-8 encoded input string. If
// the string starts with a BOM character, it gets skipped.
type Lexer struct {
	source string

	ch          rune // current character, -1 means EOI
	offset      int  // character offset
	tokenOffset int  // last token offset
	line        int  // current line number (1-based)
	tokenLine   int  // last token line
	lineOffset  int  // current line offset
	scanOffset  int  // scanning offset
	value       interface{}

	State int // lexer state, modifiable

	inStatesSelector bool
	prev             Token
}

var bomSeq = "\xef\xbb\xbf"

// Init prepares the lexer l to tokenize source by performing the full reset
// of the internal state.
func (l *Lexer) Init(source string) {
	l.source = source

	l.ch = 0
	l.offset = 0
	l.tokenOffset = 0
	l.line = 1
	l.tokenLine = 1
	l.lineOffset = 0
	l.State = 0
	l.inStatesSelector = false
	l.prev = UNAVAILABLE

	if strings.HasPrefix(source, bomSeq) {
		l.offset += len(bomSeq)
	}

	l.rewind(l.offset)
}

// Next finds and returns the next token in l.source. The source end is
// indicated by Token.EOI.
//
// The token text can be retrieved later by calling the Text() method.
func (l *Lexer) Next() Token {
restart:
	l.tokenLine = l.line
	l.tokenOffset = l.offset

	state := tmStateMap[l.State]
	hash := uint32(0)
	backupRule := -1
	var backupOffset int
	backupHash := hash
	for state >= 0 {
		var ch int
		if uint(l.ch) < tmRuneClassLen {
			ch = int(tmRuneClass[l.ch])
		} else if l.ch < 0 {
			state = int(tmLexerAction[state*tmNumClasses])
			continue
		} else {
			ch = 1
		}
		state = int(tmLexerAction[state*tmNumClasses+ch])
		if state > tmFirstRule {
			if state < 0 {
				state = (-1 - state) * 2
				backupRule = tmBacktracking[state]
				backupOffset = l.offset
				backupHash = hash
				state = tmBacktracking[state+1]
			}
			hash = hash*uint32(31) + uint32(l.ch)

			if l.ch == '\n' {
				l.line++
				l.lineOffset = l.offset
			}

			// Scan the next character.
			// Note: the following code is inlined to avoid performance implications.
			l.offset = l.scanOffset
			if l.offset < len(l.source) {
				r, w := rune(l.source[l.offset]), 1
				if r >= 0x80 {
					// not ASCII
					r, w = utf8.DecodeRuneInString(l.source[l.offset:])
				}
				l.scanOffset += w
				l.ch = r
			} else {
				l.ch = -1 // EOI
			}
		}
	}

	rule := tmFirstRule - state
recovered:
	switch rule {
	case 38:
		hh := hash & 63
		switch hh {
		case 2:
			if hash == 0x43733a82 && "lookahead" == l.source[l.tokenOffset:l.offset] {
				rule = 62
				break
			}
			if hash == 0x6856c82 && "shift" == l.source[l.tokenOffset:l.offset] {
				rule = 72
				break
			}
		case 3:
			if hash == 0x41796943 && "returns" == l.source[l.tokenOffset:l.offset] {
				rule = 69
				break
			}
		case 6:
			if hash == 0xac107346 && "assert" == l.source[l.tokenOffset:l.offset] {
				rule = 46
				break
			}
			if hash == 0x688f106 && "space" == l.source[l.tokenOffset:l.offset] {
				rule = 73
				break
			}
		case 7:
			if hash == 0x32a007 && "left" == l.source[l.tokenOffset:l.offset] {
				rule = 60
				break
			}
		case 10:
			if hash == 0x5fb57ca && "input" == l.source[l.tokenOffset:l.offset] {
				rule = 55
				break
			}
		case 11:
			if hash == 0xfde4e8cb && "brackets" == l.source[l.tokenOffset:l.offset] {
				rule = 47
				break
			}
		case 12:
			if hash == 0x621a30c && "lexer" == l.source[l.tokenOffset:l.offset] {
				rule = 61
				break
			}
		case 13:
			if hash == 0x5c2854d && "empty" == l.source[l.tokenOffset:l.offset] {
				rule = 49
				break
			}
			if hash == 0x658188d && "param" == l.source[l.tokenOffset:l.offset] {
				rule = 66
				break
			}
		case 14:
			if hash == 0x36758e && "true" == l.source[l.tokenOffset:l.offset] {
				rule = 45
				break
			}
		case 20:
			if hash == 0x375194 && "void" == l.source[l.tokenOffset:l.offset] {
				rule = 74
				break
			}
		case 24:
			if hash == 0x9fd29358 && "language" == l.source[l.tokenOffset:l.offset] {
				rule = 58
				break
			}
		case 25:
			if hash == 0xb96da299 && "inline" == l.source[l.tokenOffset:l.offset] {
				rule = 54
				break
			}
		case 28:
			if hash == 0x677c21c && "right" == l.source[l.tokenOffset:l.offset] {
				rule = 70
				break
			}
		case 31:
			if hash == 0xc4ab3c1f && "parser" == l.source[l.tokenOffset:l.offset] {
				rule = 67
				break
			}
		case 32:
			if hash == 0x540c92a0 && "nonempty" == l.source[l.tokenOffset:l.offset] {
				rule = 65
				break
			}
			if hash == 0x34a220 && "prec" == l.source[l.tokenOffset:l.offset] {
				rule = 68
				break
			}
		case 34:
			if hash == 0x1bc62 && "set" == l.source[l.tokenOffset:l.offset] {
				rule = 44
				break
			}
		case 35:
			if hash == 0x5cb1923 && "false" == l.source[l.tokenOffset:l.offset] {
				rule = 40
				break
			}
			if hash == 0xb5e903a3 && "global" == l.source[l.tokenOffset:l.offset] {
				rule = 53
				break
			}
		case 37:
			if hash == 0xb96173a5 && "import" == l.source[l.tokenOffset:l.offset] {
				rule = 42
				break
			}
			if hash == 0x6748e2e5 && "separator" == l.source[l.tokenOffset:l.offset] {
				rule = 43
				break
			}
		case 40:
			if hash == 0x53d6f968 && "nonassoc" == l.source[l.tokenOffset:l.offset] {
				rule = 64
				break
			}
		case 42:
			if hash == 0xbddafb2a && "layout" == l.source[l.tokenOffset:l.offset] {
				rule = 59
				break
			}
		case 44:
			if hash == 0x2fff6c && "flag" == l.source[l.tokenOffset:l.offset] {
				rule = 51
				break
			}
		case 48:
			if hash == 0xc97057b0 && "implements" == l.source[l.tokenOffset:l.offset] {
				rule = 41
				break
			}
		case 50:
			if hash == 0xc32 && "as" == l.source[l.tokenOffset:l.offset] {
				rule = 39
				break
			}
		case 51:
			if hash == 0xc1e742f3 && "no-eoi" == l.source[l.tokenOffset:l.offset] {
				rule = 63
				break
			}
			if hash == 0x73 && "s" == l.source[l.tokenOffset:l.offset] {
				rule = 71
				break
			}
		case 52:
			if hash == 0x8d046634 && "explicit" == l.source[l.tokenOffset:l.offset] {
				rule = 50
				break
			}
		case 53:
			if hash == 0x6be81575 && "generate" == l.source[l.tokenOffset:l.offset] {
				rule = 52
				break
			}
		case 56:
			if hash == 0x5a5a978 && "class" == l.source[l.tokenOffset:l.offset] {
				rule = 48
				break
			}
			if hash == 0x78 && "x" == l.source[l.tokenOffset:l.offset] {
				rule = 75
				break
			}
		case 57:
			if hash == 0x1df56d39 && "interface" == l.source[l.tokenOffset:l.offset] {
				rule = 56
				break
			}
		case 59:
			if hash == 0x3291bb && "lalr" == l.source[l.tokenOffset:l.offset] {
				rule = 57
				break
			}
		}
	}

	token := tmToken[rule]
	var space bool
	switch rule {
	case 0:
		if backupRule >= 0 {
			rule = backupRule
			hash = backupHash
			l.rewind(backupOffset)
		} else if l.offset == l.tokenOffset {
			l.rewind(l.scanOffset)
		}
		if rule != 0 {
			goto recovered
		}
	case 4: // templates: /%%/
		{
			l.rewind(len(l.source))
		}
	case 5: // whitespace: /[\n\r\t ]+/
		space = true
	}
	if space {
		goto restart
	}
	switch token {
	case LT:
		l.inStatesSelector = l.State == StateInitial || l.State == StateAfterColonOrEq
		l.State = StateInitial
	case GT:
		if l.inStatesSelector {
			l.State = StateAfterGT
			l.inStatesSelector = false
		} else {
			l.State = StateInitial
		}
	case ID, LEFT, RIGHT, NONASSOC, GENERATE, ASSERT, EMPTY,
		BRACKETS, INLINE, PREC, SHIFT, RETURNS, INPUT,
		NONEMPTY, GLOBAL, EXPLICIT, LOOKAHEAD, PARAM, FLAG,
		CHAR_S, CHAR_X, CLASS, INTERFACE, VOID, SPACE,
		LAYOUT, LANGUAGE, LALR:
		l.State = StateAfterID
	case LEXER, PARSER:
		if l.prev == COLONCOLON {
			l.State = StateInitial
		} else {
			l.State = StateAfterID
		}
	case ASSIGN, COLON:
		l.State = StateAfterColonOrEq
	case CODE:
		if !l.skipAction() {
			token = INVALID_TOKEN
		}
		fallthrough
	default:
		l.State = StateInitial
	}
	l.prev = token
	return token
}

// Pos returns the start and end positions of the last token returned by Next().
func (l *Lexer) Pos() (start, end int) {
	start = l.tokenOffset
	end = l.offset
	return
}

// Line returns the line number of the last token returned by Next().
func (l *Lexer) Line() int {
	return l.tokenLine
}

// Text returns the substring of the input corresponding to the last token.
func (l *Lexer) Text() string {
	return l.source[l.tokenOffset:l.offset]
}

// Value returns the value associated with the last returned token.
func (l *Lexer) Value() interface{} {
	return l.value
}

// rewind can be used in lexer actions to accept a portion of a scanned token, or to include
// more text into it.
func (l *Lexer) rewind(offset int) {
	if offset < l.offset {
		l.line -= strings.Count(l.source[offset:l.offset], "\n")
	} else {
		if offset > len(l.source) {
			offset = len(l.source)
		}
		l.line += strings.Count(l.source[l.offset:offset], "\n")
	}
	l.lineOffset = 1 + strings.LastIndexByte(l.source[:offset], '\n')

	// Scan the next character.
	l.scanOffset = offset
	l.offset = offset
	if l.offset < len(l.source) {
		r, w := rune(l.source[l.offset]), 1
		if r >= 0x80 {
			// not ASCII
			r, w = utf8.DecodeRuneInString(l.source[l.offset:])
		}
		l.scanOffset += w
		l.ch = r
	} else {
		l.ch = -1 // EOI
	}
}
