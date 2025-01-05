package analysis

import (
	"labda/eval"
	"testing"
)

func TestParseTrivial(t *testing.T) {
	tokens := []Token{Lambda, Word("x"), Word("a"), Word("x")}
	want := eval.Abstraction{"x", eval.Application{eval.Variable{"a"}, eval.Variable{"x"}}}
	got := Parse(tokens)
	if want != got {
		t.Fatalf("Wanted: %#v, got: %#v", want, got)
	}
}

func TestParseComplex(t *testing.T) {
	tokens := []Token{LParen, Lambda, Word("x"), Word("x"), String("Hello, World!"), RParen, LParen, Word("123"), RParen}
	want := eval.Application{eval.Abstraction{"x", eval.Application{eval.Variable{"x"}, eval.StringLit{"Hello, World!"}}}, eval.Variable{"123"}}
	got := Parse(tokens)
	if want != got {
		t.Fatalf("Wanted: %#v, got: %#v", want, got)
	}
}

func TestParseChain(t *testing.T) {
	tokens := []Token{Word("abs"), Word("1"), Word("2"), Word("3")}
	want := eval.Application{eval.Application{eval.Application{eval.Variable{"abs"}, eval.Variable{"1"}}, eval.Variable{"2"}}, eval.Variable{"3"}}
	got := Parse(tokens)
	if want != got {
		t.Fatalf("Wanted: %#v, got: %#v", want, got)
	}
}
