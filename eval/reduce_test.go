package eval

import (
	ana "labda/analysis"
	"testing"
)

func TestCollapse(t *testing.T) {
	expr := ana.Application{Body: ana.Abstraction{Variable: "x", Term: ana.Variable{Name: "x"}}, Argument: ana.StringLit{Value: "Hello, World!"}}
	want := ana.StringLit{Value: "Hello, World!"}
	got := Collapse(expr)
	if want != got {
		t.Fatalf("Wanted: %#v, got: %#v", want, got)
	}
}

func TestCollapseHalf(t *testing.T) {
	expr := ana.Application{Body: ana.Application{Body: ana.Abstraction{Variable: "x", Term: ana.Variable{Name: "x"}}, Argument: ana.StringLit{Value: "1"}}, Argument: ana.Application{Body: ana.Abstraction{Variable: "x", Term: ana.Variable{Name: "x"}}, Argument: ana.StringLit{Value: "2"}}}
	want := ana.Application{Body: ana.StringLit{Value: "1"}, Argument: ana.StringLit{Value: "2"}}
	got := Collapse(Collapse(expr))
	if want != got {
		t.Fatalf("Wanted: %#v, got: %#v", want, got)
	}
}
