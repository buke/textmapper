package syntax_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/inspirer/textmapper/syntax"
	"github.com/inspirer/textmapper/util/dump"
)

func expand(m *syntax.Model) error {
	return syntax.Expand(m, &syntax.ExpandOptions{})
}

var modelTests = []struct {
	fnName string
	fn     func(m *syntax.Model) error
	input  string
	want   string
}{
	{"PropagateLookaheads", syntax.PropagateLookaheads,
		`%lookahead flag V; Z: A<V=true>; A: B; B: C; C: [V] d;`,
		`%flag V; Z: A<V=true>; A<V>: B<V=V>; B<V>: C<V=V>; C<V>: [V] d;`,
	},
	{"PropagateLookaheads", syntax.PropagateLookaheads,
		`%lookahead flag V; Z: A<V=true>; A: B; B: (C | c); C: [V] d;`,
		`%flag V; Z: A<V=true>; A<V>: B<V=V>; B<V>: C<V=V> | c; C<V>: [V] d;`,
	},
	{"PropagateLookaheads", syntax.PropagateLookaheads,
		`%flag P; %lookahead flag V; Z<P>: a A; A: B<V=true>; B: C; C: d Z<P=V>;`,
		`%flag P; %flag V; Z<P>: a A; A: B<V=true>; B<V>: C<V=V>; C<V>: d Z<P=V>;`,
	},
	{"PropagateLookaheads", syntax.PropagateLookaheads,
		`%lookahead flag V; Z: A<V=true>; A: B; B: C? c; C: [V] d;`,
		`ERR: input:1:40: cannot propagate lookahead flag V through nonterminal B; avoid nullable alternatives and optional clauses`,
	},
	{"PropagateLookaheads", syntax.PropagateLookaheads,
		`%lookahead flag V; Z: A<V=true>; A: B; B: (C | c)?; C: [V] d;`,
		`ERR: input:1:40: cannot propagate lookahead flag V through nonterminal B; avoid nullable alternatives and optional clauses`,
	},
	{"PropagateLookaheads", syntax.PropagateLookaheads,
		`%lookahead flag V; Z: A<V=true>; A: B; B: (C | c)*; C: [V] d;`,
		`ERR: input:1:40: cannot propagate lookahead flag V through nonterminal B; avoid nullable alternatives and optional clauses`,
	},
	{"PropagateLookaheads", syntax.PropagateLookaheads,
		`%lookahead flag V; Z: A<V=true>; A: B; B: (C | c)+; C: [V] d;`,
		`%flag V; Z: A<V=true>; A<V>: B<V=V>; B<V>: (C<V=V> | c)+; C<V>: [V] d;`,
	},
	{"PropagateLookaheads", syntax.PropagateLookaheads,
		`%lookahead flag V; Z: A<V=true>; A: B; B: (C | C<V=false> c)+; C: [V] d;`,
		`%flag V; Z: A<V=true>; A<V>: B<V=V>; B<V>: (C<V=V> | C<V=false> c)+; C<V>: [V] d;`,
	},
	{"PropagateLookaheads", syntax.PropagateLookaheads,
		`%lookahead flag V; Z: A<V=true>; A: B; B: (C | c?); C: [V] d;`,
		`ERR: input:1:40: cannot propagate lookahead flag V through nonterminal B; avoid nullable alternatives and optional clauses`,
	},
	{"PropagateLookaheads", syntax.PropagateLookaheads,
		`%lookahead flag V; Z: A<V=true>; A: B; B: C; C: d;`,
		`ERR: input:1:25: V is not used in A`,
	},
	{"PropagateLookaheads", syntax.PropagateLookaheads,
		`%lookahead flag V; Z: A<V=true>; A: B<V=true>; B: C; C: [V] d;`,
		`ERR: input:1:25: V is not used in A`,
	},
	{"PropagateLookaheads", syntax.PropagateLookaheads,
		`%lookahead flag V;Z2:[V] A; A: B<V=true>; B: C; C: [V] d;`,
		`ERR: input:1:23: lookahead flag V is never provided`,
	},
	// Lookahead flag arguments in token sets.
	{"PropagateLookaheads", syntax.PropagateLookaheads,
		`%lookahead flag V; Z: A set(B<V=true>); A: c; B: C; C: [V] d;`,
		`%flag V; Z: A set(B<V=true>); A: c; B<V>: C<V=V>; C<V>: [V] d; `,
	},

	// Template instantiations.
	{"Instantiate", syntax.Instantiate,
		`%input Z; %flag V; Z: A set(B<V=true>); A: c; B<V>: C<V=V>; C<V>: [V] d; F2: c;`,
		`%input Z; Z: A set(B_V); A: c; B_V: C_V; C_V: d;`,
	},
	{"Instantiate", syntax.Instantiate,
		`%input Z; %flag V; Z: b B<V=true> | c B<V=false>; B<V>: [V] a | b;`,
		`%input Z; Z: b B_V | c B; B: b; B_V: a | b;`,
	},
	{"Instantiate", syntax.Instantiate,
		`%input Z; %flag V; %flag T; Z: b B<V=true> | c B<V=false>; B<V>: [V] a | Q<T=V>; Q<T>: [!T] n | [T] t;`,
		`%input Z; Z: b B_V | c B; B: Q; B_V: a | Q_T; Q: n; Q_T: t;`,
	},
	{"Instantiate", syntax.Instantiate,
		`%input Z; %flag A; %flag B;
       Z: F<A=true,B=true> | F<A=true,B=false> | F<A=false,B=true> | F<A=false, B=false>;
       F<A,B>: [A && B] n a b | [A] a | [B] b | [!A && !B] n;`,
		`%input Z; Z: F_A_B | F_A | F_B | F; F: n; F_A: a; F_A_B: n a b | a | b; F_B: b;`,
	},
	{"Instantiate", syntax.Instantiate,
		`%input Z; %flag T; Z: F<T=false> F<T=true>; F<T>: a ([T] b | a) a;`,
		`%input Z; Z: F F_T; F: a (a) a; F_T: a (b | a) a;`,
	},
	{"Instantiate", syntax.Instantiate,
		`%input Z; %flag T; Z: F<T=false> F<T=true>; F<T>: a ([T] b) a;`,
		`%input Z; Z: F F_T; F: a a; F_T: a (b) a;`,
	},

	// Syntax sugar expansion.
	{"Expand", expand,
		`Z: a?;`,
		`Z: a | ;`,
	},
	{"Expand", expand,
		`Z: a? | b?;`,
		`Z: a | | b ;`,
	},
	{"Expand", expand,
		`Z: (a | b)?;`,
		`Z: a | b | ;`,
	},
	{"Expand", expand,
		`Z: (a b?)?;`,
		`Z: a b | a |  ;`,
	},
	{"Expand", expand,
		`Z: (a b|b) (c|d);`,
		`Z: a b c | a b d | b c | b d ;`,
	},
	{"Expand", expand,
		`Z: a? %prec b ;`,
		`Z: a %prec b | %prec b ;`,
	},
	{"Expand", expand,
		`Z: a? -> A ;`,
		`Z: a -> A | -> A ;`,
	},
	{"Expand", expand,
		`Z: a=a? ;`,
		`Z: a=a | ;`,
	},
	{"Expand", expand,
		`Z: a? {Foo} -> A ;`,
		`Z: a {Foo} -> A | {Foo} -> A ;`,
	},
	{"Expand", expand,
		`Z: a+ | q ;`,
		`A_list: A_list a | a; Z: A_list | q ;`,
	},
	{"Expand", expand,
		`Z: a* | q ;`,
		`A_optlist: A_optlist a | ; Z: A_optlist | q ;`,
	},
	{"Expand", expand,
		`Z: b | (a separator b)+ ;`,
		`A_list_B_separated: A_list_B_separated b a | a; Z: b | A_list_B_separated ;`,
	},
	{"Expand", expand,
		`Z: b | (a separator b)* ;`,
		`A_list_B_separated: A_list_B_separated b a | a; A_list_B_separatedopt: A_list_B_separated | ; Z: b | A_list_B_separatedopt ;`,
	},
	{"Expand", expand,
		`Z: set(a | ~b);`,
		`Z: set(a | ~b);`, // top level sets are not expanded
	},
	{"Expand", expand,
		`Z: a b set(a | ~b) | c ;`,
		`Z: a b setof_a_or_not_b | c ; setof_a_or_not_b: set(a | ~b) ;`,
	},
	{"Expand", expand,
		`Z: (?= A); A:a|b;`,
		`Z: (?= A); A:a|b;`, // top level lookaheads are not expanded
	},
	{"Expand", expand,
		`Z: a (?= A & !B) b | c; A: a|b; B: a|b;`,
		`Z: a lookahead_A_notB b | c; lookahead_A_notB: (?= A & !B); A: a|b; B: a|b;`,
	},
	{"Expand", expand,
		`Z: A+ | C+ | B+; A: a|x; B: b|y; C: c|z;`, // sorting test
		`A_list: A_list A | A; B_list: B_list B | B; C_list: C_list C | C; Z: A_list | C_list | B_list; A: a|x; B: b|y; C: c|z;`,
	},
	{"Expand", expand,
		`%input X; X: B+ | Y+ | A+; A: a|x; B: b|y; Y: c|z;`, // sorting test #2
		`%input X; A_list: A_list A | A; B_list: B_list B | B; X: B_list | Y_list | A_list; Y_list: Y_list Y | Y; A: a|x; B: b|y; Y: c|z;`,
	},
}

func TestModelTransforms(t *testing.T) {
	for _, tc := range modelTests {
		model, err := parse(tc.input)
		if err != nil {
			t.Errorf("cannot parse %q: %v", tc.input, err)
			continue
		}

		if err := tc.fn(model); err != nil {
			const prefix = "ERR: "
			if !strings.HasPrefix(tc.want, prefix) {
				t.Errorf("%v(%v) failed with %v", tc.fnName, tc.input, err)
				continue
			}
			if got := fmt.Sprintf("ERR: %v", err); got != tc.want {
				t.Errorf("%v(%v) failed with %v, want: %v", tc.fnName, tc.input, got, tc.want)
			}
			continue
		}

		want, err := parse(tc.want)
		if err != nil {
			t.Errorf("cannot parse %q: %v", tc.want, err)
			continue
		}

		stripSelfRef(model)
		stripOrigin(model)
		stripSelfRef(want)
		stripOrigin(want)
		if diff := dump.Diff(want, model); diff != "" {
			t.Errorf("%v(%v) produced diff (-want +got):\n%s", tc.fnName, tc.input, diff)
		}
	}
}
