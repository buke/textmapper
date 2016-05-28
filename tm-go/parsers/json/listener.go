package json

import "fmt"

type NodeType int

type Listener interface {
	Node(t NodeType, offset, endoffset int)
}

const (
	JSONText NodeType = iota + 1
	JSONValue
	JSONObject
	JSONMember
	JSONArray
	NodeTypeMax
)

var ruleNodeType = [...]NodeType{
	JSONText, // JSONText ::= JSONValue_A
	JSONValue, // JSONValue ::= 'null'
	JSONValue, // JSONValue ::= 'true'
	JSONValue, // JSONValue ::= 'false'
	JSONValue, // JSONValue ::= 'B'
	JSONValue, // JSONValue ::= JSONObject
	JSONValue, // JSONValue ::= JSONArray
	JSONValue, // JSONValue ::= JSONString
	JSONValue, // JSONValue ::= JSONNumber
	JSONValue, // JSONValue_A ::= 'null'
	JSONValue, // JSONValue_A ::= 'true'
	JSONValue, // JSONValue_A ::= 'false'
	JSONValue, // JSONValue_A ::= 'A'
	JSONValue, // JSONValue_A ::= JSONObject
	JSONValue, // JSONValue_A ::= JSONArray
	JSONValue, // JSONValue_A ::= JSONString
	JSONValue, // JSONValue_A ::= JSONNumber
	JSONObject, // JSONObject ::= '{' JSONMemberList '}'
	JSONObject, // JSONObject ::= '{' '}'
	JSONMember, // JSONMember ::= JSONString ':' JSONValue
	0, // JSONMemberList ::= JSONMember
	0, // JSONMemberList ::= JSONMemberList ',' JSONMember
	JSONArray, // JSONArray ::= '[' JSONElementListopt ']'
	0, // JSONElementList ::= JSONValue_A
	0, // JSONElementList ::= JSONElementList ',' JSONValue_A
	0, // JSONElementListopt ::= JSONElementList
	0, // JSONElementListopt ::=
}

var nodeTypeStr = [...]string{
	"NONE",
	"JSONText",
	"JSONValue",
	"JSONObject",
	"JSONMember",
	"JSONArray",
}

func (t NodeType) String() string {
	if t >= 0 && int(t) < len(nodeTypeStr) {
		return nodeTypeStr[t]
	}
	return fmt.Sprintf("node(%d)", t)
}
