package main

import (
	"labda/analysis"
	"labda/eval"
	"testing"
)

func TestFull(t *testing.T) {
	source := "($x.x \"\\x2F\") ($a.a)"
	want := eval.StringLit{Value: "\x2F"}
	tokens := analysis.Lex(source)
	expr := analysis.Parse(tokens)
	got := expr.Reduce()
	if !want.Compare(got) {
		t.Fatalf("Wanted: %#v, got: %#v", want, got)
	}
}

func TestYComb(t *testing.T) {
	source := "$f.($x.f(x x))($x.f(x x))"
	side := eval.Abstraction{Variable: "x", Term: &eval.Application{Body: &eval.Variable{Name: "f"}, Argument: &eval.Application{Body: &eval.Variable{Name: "x"}, Argument: &eval.Variable{Name: "x"}}}}
	want := eval.Abstraction{Variable: "f", Term: &eval.Application{Body: &side, Argument: &side}}
	tokens := analysis.Lex(source)
	expr := analysis.Parse(tokens)
	got := expr.Reduce()
	if !want.Compare(got) {
		t.Fatalf("Wanted: %#v, got: %#v", want, got)
	}
}
