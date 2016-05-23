package tm_test

import (
	"fmt"
	"testing"

	"github.com/inspirer/textmapper/tm-go/parsers/tm"
	"github.com/inspirer/textmapper/tm-go/parsers/tm/ast"
)

func PanicOnError(line, offset, len int, msg string) {
	panic(fmt.Sprintf("%d, %d: %s", line, offset, msg))
}

func testParser(input []byte, t *testing.T) *ast.Input {
	l := new(tm.Lexer)
	l.Init(input, PanicOnError)

	p := new(tm.Parser)
	p.Init(PanicOnError)
	ok, val := p.ParseInput(l)
	if !ok {
		t.Error("Not parsed.")
		return nil
	}
	if val == nil {
		t.Errorf("Not input: %v", val)
	}
	return val
}

func TestParserExample(t *testing.T) {
	input := testParser([]byte(`
language aaa.bbb.json(java);

positions = "line,offset"
endpositions = "offset"
package = "org.textmapper.json"

:: lexer

'{': /\{/
'}': /\}/
'[': /\[/
']': /\]/
':': /:/
',': /,/

space: /[\t\r\n ]+/ (space)

hex = /[0-9a-fA-F]/

# TODO
JSONString: /"([^"\\]|\\(["\/\\bfnrt]|u{hex}{4}))*"/
#JSONString: /"([^"\\\x00-\x1f]|\\(["\/\\bfnrt]|u{hex}{4}))*"/

fraction = /\.[0-9]+/
exp = /[eE][+-]?[0-9]+/
JSONNumber: /-?(0|[1-9][0-9]*){fraction}?{exp}?/

'null': /null/
'true': /true/
'false': /false/


:: parser

%input JSONText;

JSONText ::=
	  JSONValue ;

JSONValue ::=
	  'null'
	| 'true'
	| 'false'
	| @aa: JSONObject
	| JSONArray
	| JSONString
	| JSONNumber
;

JSONObject ::=
	  ('{' JSONMemberList? '}' separator ',')+ ;

JSONMember ::=
	  JSONString ':' JSONValue ;

JSONMemberList ::=
	  JSONMember
	| JSONMemberList ',' JSONMember
;

JSONArray ::=
	  '[' JSONElementList? ']' ;

JSONElementList ::=
	  JSONValue
	| JSONElementList ',' JSONValue
;

%%

${template java_lexer.lexercode}
${end}


	`), t)

	if input == nil {
		return
	}

	if input.Header.Name.QualifiedId != "aaa.bbb.json" {
		t.Errorf("input.Header.Name.QualifiedId = %v, want aaa.bbb.json", input.Header.Name.QualifiedId)
	}
}