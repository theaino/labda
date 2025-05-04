package eval

import (
	"testing"
)

func TestCollapse(t *testing.T) {
	expr := Application{Body: &Abstraction{Variable: "x", Term: &Variable{Name: "x"}}, Argument: &StringLit{Value: "Hello, World!"}}
	want := StringLit{Value: "Hello, World!"}
	got := expr.Reduce()
	if !want.Compare(got) {
		t.Fatalf("Wanted: %v, got: %v", want, got)
	}
}

func TestCollapseHalf(t *testing.T) {
	expr := Application{Body: &Application{Body: &Abstraction{Variable: "x", Term: &Variable{Name: "x"}}, Argument: &StringLit{Value: "1"}}, Argument: &Application{Body: &Abstraction{Variable: "x", Term: &Variable{Name: "x"}}, Argument: &StringLit{Value: "2"}}}
	want := Application{Body: &StringLit{Value: "1"}, Argument: &Application{Body: &Abstraction{Variable: "x", Term: &Variable{Name: "x"}}, Argument: &StringLit{Value: "2"}}}
	got := expr.Reduce()
	if !want.Compare(got) {
		t.Fatalf("Wanted: %v, got: %v", want, got)
	}
}
