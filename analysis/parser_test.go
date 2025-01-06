package analysis

import (
	"labda/eval"
	"testing"
)

func TestParseTrivial(t *testing.T) {
	tokens := []Token{Lambda, Word("x"), Dot, Word("a"), Word("x")}
	want := eval.Abstraction{Variable: "x", Term: eval.Application{Body: eval.Variable{Name: "a"}, Argument: eval.Variable{Name: "x"}}}
	got := Parse(tokens)
	if want != got {
		t.Fatalf("Wanted: %v, got: %v", want, got)
	}
}

func TestParseComplex(t *testing.T) {
	tokens := []Token{LParen, Lambda, Word("x"), Dot, Word("x"), String("Hello, World!"), RParen, LParen, Word("123"), RParen}
	want := eval.Application{Body: eval.Abstraction{Variable: "x", Term: eval.Application{Body: eval.Variable{Name: "x"}, Argument: eval.StringLit{Value: "Hello, World!"}}}, Argument: eval.Variable{Name: "123"}}
	got := Parse(tokens)
	if want != got {
		t.Fatalf("Wanted: %v, got: %v", want, got)
	}
}

func TestParseChain(t *testing.T) {
	tokens := []Token{Word("abs"), Word("1"), Word("2"), Word("3")}
	want := eval.Application{Body: eval.Application{Body: eval.Application{Body: eval.Variable{Name: "abs"}, Argument: eval.Variable{Name: "1"}}, Argument: eval.Variable{Name: "2"}}, Argument: eval.Variable{Name: "3"}}
	got := Parse(tokens)
	if want != got {
		t.Fatalf("Wanted: %v, got: %v", want, got)
	}
}
