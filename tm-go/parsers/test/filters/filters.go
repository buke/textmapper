// generated by Textmapper; DO NOT EDIT

package filters

import (
	"github.com/inspirer/textmapper/tm-go/parsers/test"
)

type NodeFilter func(nt test.NodeType) bool

var (
	Block         = func(t test.NodeType) bool { return t == test.Block }
	Decl1         = func(t test.NodeType) bool { return t == test.Decl1 }
	Decl2         = func(t test.NodeType) bool { return t == test.Decl2 }
	QualifiedName = func(t test.NodeType) bool { return t == test.QualifiedName }
	Test          = func(t test.NodeType) bool { return t == test.Test }
	Declaration   = OneOf(test.Declaration...)
)

func OneOf(types ...test.NodeType) NodeFilter {
	if len(types) == 0 {
		return func(test.NodeType) bool { return false }
	}
	const bits = 32
	size := (int(types[len(types)-1]) + bits - 1) / bits
	bitarr := make([]int32, size)
	for _, t := range types {
		bitarr[uint(t)/bits] |= 1 << (uint(t) % bits)
	}
	return func(t test.NodeType) bool {
		return bitarr[uint(t)/bits]&(1<<(uint(t)%bits)) != 0
	}
}
