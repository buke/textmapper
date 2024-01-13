// generated by Textmapper; DO NOT EDIT

package json

import (
	"fmt"

	"github.com/inspirer/textmapper/parsers/json/token"
)

var tmNonterminals = [...]string{
	"JSONText",
	"JSONValue",
	"JSONValue_A",
	"EmptyObject",
	"lookahead_EmptyObject",
	"JSONObject",
	"lookahead_notEmptyObject",
	"JSONMember",
	"JSONMemberList",
	"JSONArray",
	"JSONElementList",
	"JSONElementListopt",
}

func symbolName(sym int32) string {
	if sym == noToken {
		return "<no-token>"
	}
	if sym < int32(token.NumTokens) {
		return token.Type(sym).String()
	}
	if i := int(sym) - int(token.NumTokens); i < len(tmNonterminals) {
		return tmNonterminals[i]
	}
	return fmt.Sprintf("nonterminal(%d)", sym)
}

var tmDefGoto = []int32{
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
}

var tmGoto = []int32{
	59, 34, 50, 5, 20, 16, 35, 27, 45, 47, 63, 64,
}

var tmDefAct = []int32{
	-1, 20, -1, 17, 18, 10, 11, 12, 13, 0, 15, -1, 14, -1, 16, 28, -1, -1, -1,
	-1, -1, 27, 19, 22, -1, 25, -1, 29, -1, 21, -1, 8, 9, 1, 2, 3, 4, 24, 6, 5,
	7, 26, -1, -1, -1,
}

const tmActionBase = -19

var tmAction = []int32{
	13, -19, -2, -19, -19, -19, -19, -19, -19, -19, -19, -1, -19, 2, -19, -19,
	46, 14, 7, 51, 13, -19, -19, -19, 25, -19, 53, -19, 28, -19, 24, -19, -19,
	-19, -19, -19, -19, -19, -19, -19, -19, -19, 58, -19, -19,
}

const tmTableLen = 76

var tmTable = []int8{
	32, -20, -4, 31, -21, 10, 43, 10, -5, -6, -24, -7, -8, -9, -10, 32, 12, -4,
	12, -23, 11, 11, 11, -5, -6, 10, -7, -8, -9, -10, 32, -30, -4, 38, -26, 13,
	12, 13, -33, -34, 11, -35, -36, -37, 39, -38, 25, 14, 11, 14, 9, 30, 15, -22,
	-25, 13, -31, 41, -46, 42, -32, -26, 37, 13, 26, 16, 17, 14, 0, 0, 27, 0, 0,
	0, 0, 40,
}

var tmCheck = []int8{
	2, 2, 4, 5, 2, 0, 1, 2, 10, 11, 3, 13, 14, 15, 16, 2, 0, 4, 2, 5, 0, 1, 2,
	10, 11, 20, 13, 14, 15, 16, 2, 6, 4, 28, 10, 0, 20, 2, 10, 11, 20, 13, 14,
	15, 28, 17, 19, 0, 28, 2, 0, 5, 2, 7, 3, 20, 3, 30, 0, 0, 7, 10, 28, 28, 19,
	2, 2, 20, -1, -1, 20, -1, -1, -1, -1, 28,
}

var tmRuleLen = []int8{
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 0, 4, 3, 0, 3, 1,
	3, 3, 1, 3, 1, 0, 0,
}

var tmRuleSymbol = []int32{
	19, 20, 20, 20, 20, 20, 20, 20, 20, 20, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	22, 23, 24, 24, 25, 26, 27, 27, 28, 29, 29, 30, 30, 25,
}

var tmRuleType = [...]NodeType{
	JSONText,    // JSONText : JSONValue_A
	JSONValue,   // JSONValue : 'null'
	JSONValue,   // JSONValue : 'true'
	JSONValue,   // JSONValue : 'false'
	JSONValue,   // JSONValue : 'B'
	JSONValue,   // JSONValue : JSONObject
	JSONValue,   // JSONValue : EmptyObject
	JSONValue,   // JSONValue : JSONArray
	JSONValue,   // JSONValue : JSONString
	JSONValue,   // JSONValue : JSONNumber
	JSONValue,   // JSONValue_A : 'null'
	JSONValue,   // JSONValue_A : 'true'
	JSONValue,   // JSONValue_A : 'false'
	JSONValue,   // JSONValue_A : 'A'
	JSONValue,   // JSONValue_A : JSONObject
	JSONValue,   // JSONValue_A : EmptyObject
	JSONValue,   // JSONValue_A : JSONArray
	JSONValue,   // JSONValue_A : JSONString
	JSONValue,   // JSONValue_A : JSONNumber
	EmptyObject, // EmptyObject : lookahead_EmptyObject '{' '}'
	0,           // lookahead_EmptyObject :
	JSONObject,  // JSONObject : lookahead_notEmptyObject '{' JSONMemberList '}'
	JSONObject,  // JSONObject : lookahead_notEmptyObject '{' '}'
	0,           // lookahead_notEmptyObject :
	JSONMember,  // JSONMember : JSONString ':' JSONValue
	0,           // JSONMemberList : JSONMember
	0,           // JSONMemberList : JSONMemberList ',' JSONMember
	JSONArray,   // JSONArray : '[' JSONElementListopt ']'
	0,           // JSONElementList : JSONValue_A
	0,           // JSONElementList : JSONElementList ',' JSONValue_A
	0,           // JSONElementListopt : JSONElementList
	0,           // JSONElementListopt :
}

// set(first JSONValue_A) = LBRACE, LBRACK, JSONSTRING, JSONNUMBER, NULL, TRUE, FALSE, CHAR_A
var Literals = []token.Type{
	2, 4, 10, 11, 13, 14, 15, 16,
}

// set(follow ERROR) =
var afterErr = []token.Type{}
