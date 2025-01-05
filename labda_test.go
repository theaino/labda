package main

import (
	ana "labda/analysis"
	"labda/eval"
	"testing"
)

func TestFull(t *testing.T) {
	source := "($x x \"\\x2F\") ($a a)"
	want := ana.StringLit{Value: "\x2F"}
	tokens := ana.Lex(source)
	expr := ana.Parse(tokens)
	got := eval.Collapse(expr)
	if want != got {
		t.Fatalf("Wanted: %#v, got: %#v", want, got)
	}
}

func TestYComb(t *testing.T) {
	source := "$f ($x f(x x))($x f(x x))"
	side := ana.Abstraction{Variable: "x", Term: ana.Application{Body: ana.Variable{Name: "f"}, Argument: ana.Application{Body: ana.Variable{Name: "x"}, Argument: ana.Variable{Name: "x"}}}}
	want := ana.Abstraction{Variable: "f", Term: ana.Application{Body: side, Argument: side}}
	tokens := ana.Lex(source)
	expr := ana.Parse(tokens)
	got := eval.Collapse(expr)
	if want != got {
		t.Fatalf("Wanted: %#v, got: %#v", want, got)
	}
}
