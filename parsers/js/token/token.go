// generated by Textmapper; DO NOT EDIT

package token

import (
	"fmt"
)

// Token is an enum of all terminal symbols of the js language.
type Token int32

// Token values.
const (
	UNAVAILABLE Token = iota - 1
	EOI
	INVALID_TOKEN
	ERROR
	WHITESPACE
	MULTILINECOMMENT
	SINGLELINECOMMENT
	IDENTIFIER
	PRIVATEIDENTIFIER
	AWAIT              // await
	BREAK              // break
	CASE               // case
	CATCH              // catch
	CLASS              // class
	CONST              // const
	CONTINUE           // continue
	DEBUGGER           // debugger
	DEFAULT            // default
	DELETE             // delete
	DO                 // do
	ELSE               // else
	EXPORT             // export
	EXTENDS            // extends
	FINALLY            // finally
	FOR                // for
	FUNCTION           // function
	IF                 // if
	IMPORT             // import
	IN                 // in
	INSTANCEOF         // instanceof
	NEW                // new
	RETURN             // return
	SUPER              // super
	SWITCH             // switch
	THIS               // this
	THROW              // throw
	TRY                // try
	TYPEOF             // typeof
	VAR                // var
	VOID               // void
	WHILE              // while
	WITH               // with
	YIELD              // yield
	ENUM               // enum
	NULL               // null
	TRUE               // true
	FALSE              // false
	AS                 // as
	ASSERT             // assert
	ASSERTS            // asserts
	ASYNC              // async
	FROM               // from
	GET                // get
	LET                // let
	OF                 // of
	SET                // set
	STATIC             // static
	TARGET             // target
	SATISFIES          // satisfies
	IMPLEMENTS         // implements
	INTERFACE          // interface
	PRIVATE            // private
	PROTECTED          // protected
	PUBLIC             // public
	ANY                // any
	UNKNOWN            // unknown
	BOOLEAN            // boolean
	NUMBER             // number
	STRING             // string
	SYMBOL             // symbol
	BIGINT             // bigint
	UNDEFINED          // undefined
	NEVER              // never
	OBJECT             // object
	ABSTRACT           // abstract
	CONSTRUCTOR        // constructor
	DECLARE            // declare
	GLOBAL             // global
	IS                 // is
	MODULE             // module
	NAMESPACE          // namespace
	OVERRIDE           // override
	REQUIRE            // require
	TYPE               // type
	ACCESSOR           // accessor
	READONLY           // readonly
	KEYOF              // keyof
	UNIQUE             // unique
	INFER              // infer
	LBRACE             // {
	RBRACE             // }
	LPAREN             // (
	RPAREN             // )
	LBRACK             // [
	RBRACK             // ]
	DOT                // .
	DOTDOTDOT          // ...
	SEMICOLON          // ;
	COMMA              // ,
	LT                 // <
	GT                 // >
	LTASSIGN           // <=
	GTASSIGN           // >=
	ASSIGNASSIGN       // ==
	EXCLASSIGN         // !=
	ASSIGNASSIGNASSIGN // ===
	EXCLASSIGNASSIGN   // !==
	ATSIGN             // @
	PLUS               // +
	MINUS              // -
	MULT               // *
	DIV                // /
	REM                // %
	PLUSPLUS           // ++
	MINUSMINUS         // --
	LTLT               // <<
	GTGT               // >>
	GTGTGT             // >>>
	AND                // &
	OR                 // |
	XOR                // ^
	EXCL               // !
	TILDE              // ~
	ANDAND             // &&
	OROR               // ||
	QUEST              // ?
	QUESTQUEST         // ??
	QUESTDOT           // ?.
	COLON              // :
	ASSIGN             // =
	PLUSASSIGN         // +=
	MINUSASSIGN        // -=
	MULTASSIGN         // *=
	DIVASSIGN          // /=
	REMASSIGN          // %=
	LTLTASSIGN         // <<=
	GTGTASSIGN         // >>=
	GTGTGTASSIGN       // >>>=
	ANDASSIGN          // &=
	ORASSIGN           // |=
	XORASSIGN          // ^=
	ASSIGNGT           // =>
	MULTMULT           // **
	MULTMULTASSIGN     // **=
	QUESTQUESTASSIGN   // ??=
	ORORASSIGN         // ||=
	ANDANDASSIGN       // &&=
	NUMERICLITERAL
	STRINGLITERAL
	NOSUBSTITUTIONTEMPLATE
	TEMPLATEHEAD
	TEMPLATEMIDDLE
	TEMPLATETAIL
	REGULAREXPRESSIONLITERAL
	JSXSTRINGLITERAL
	JSXIDENTIFIER
	JSXTEXT
	RESOLVESHIFT

	NumTokens
)

var tokenStr = [...]string{
	"EOI",
	"INVALID_TOKEN",
	"ERROR",
	"WHITESPACE",
	"MULTILINECOMMENT",
	"SINGLELINECOMMENT",
	"IDENTIFIER",
	"PRIVATEIDENTIFIER",
	"await",
	"break",
	"case",
	"catch",
	"class",
	"const",
	"continue",
	"debugger",
	"default",
	"delete",
	"do",
	"else",
	"export",
	"extends",
	"finally",
	"for",
	"function",
	"if",
	"import",
	"in",
	"instanceof",
	"new",
	"return",
	"super",
	"switch",
	"this",
	"throw",
	"try",
	"typeof",
	"var",
	"void",
	"while",
	"with",
	"yield",
	"enum",
	"null",
	"true",
	"false",
	"as",
	"assert",
	"asserts",
	"async",
	"from",
	"get",
	"let",
	"of",
	"set",
	"static",
	"target",
	"satisfies",
	"implements",
	"interface",
	"private",
	"protected",
	"public",
	"any",
	"unknown",
	"boolean",
	"number",
	"string",
	"symbol",
	"bigint",
	"undefined",
	"never",
	"object",
	"abstract",
	"constructor",
	"declare",
	"global",
	"is",
	"module",
	"namespace",
	"override",
	"require",
	"type",
	"accessor",
	"readonly",
	"keyof",
	"unique",
	"infer",
	"{",
	"}",
	"(",
	")",
	"[",
	"]",
	".",
	"...",
	";",
	",",
	"<",
	">",
	"<=",
	">=",
	"==",
	"!=",
	"===",
	"!==",
	"@",
	"+",
	"-",
	"*",
	"/",
	"%",
	"++",
	"--",
	"<<",
	">>",
	">>>",
	"&",
	"|",
	"^",
	"!",
	"~",
	"&&",
	"||",
	"?",
	"??",
	"?.",
	":",
	"=",
	"+=",
	"-=",
	"*=",
	"/=",
	"%=",
	"<<=",
	">>=",
	">>>=",
	"&=",
	"|=",
	"^=",
	"=>",
	"**",
	"**=",
	"??=",
	"||=",
	"&&=",
	"NUMERICLITERAL",
	"STRINGLITERAL",
	"NOSUBSTITUTIONTEMPLATE",
	"TEMPLATEHEAD",
	"TEMPLATEMIDDLE",
	"TEMPLATETAIL",
	"REGULAREXPRESSIONLITERAL",
	"JSXSTRINGLITERAL",
	"JSXIDENTIFIER",
	"JSXTEXT",
	"RESOLVESHIFT",
}

func (tok Token) String() string {
	if tok >= 0 && int(tok) < len(tokenStr) {
		return tokenStr[tok]
	}
	return fmt.Sprintf("token(%d)", tok)
}
