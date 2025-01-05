package analysis

import (
	"testing"
)

func TestParseTrivial(t *testing.T) {
	tokens := []Token{Lambda, Word("x"), Word("a"), Word("x")}
	want := Abstraction{"x", Application{Variable{"a"}, Variable{"x"}}}
	got := Parse(tokens)
	if want != got {
		t.Fatalf("Wanted: %#v, got: %#v", want, got)
	}
}

func TestParseComplex(t *testing.T) {
	tokens := []Token{LParen, Lambda, Word("x"), Word("x"), String("Hello, World!"), RParen, LParen, Word("123"), RParen}
	want := Application{Abstraction{"x", Application{Variable{"x"}, StringLit{"Hello, World!"}}}, Variable{"123"}}
	got := Parse(tokens)
	if want != got {
		t.Fatalf("Wanted: %#v, got: %#v", want, got)
	}
}

func TestParseChain(t *testing.T) {
	tokens := []Token{Word("abs"), Word("1"), Word("2"), Word("3")}
	want := Application{Application{Application{Variable{"abs"}, Variable{"1"}}, Variable{"2"}}, Variable{"3"}}
	got := Parse(tokens)
	if want != got {
		t.Fatalf("Wanted: %#v, got: %#v", want, got)
	}
}
