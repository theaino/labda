package eval

import (
	"fmt"
	"strconv"
)

type Expr interface {
	Reduce() Expr
	Apply(Expr) Expr
	String() string
	Compare(Expr) bool
}

type Abstraction struct {
	Variable string
	Term Expr
}

type PathedAbstraction struct {
	Variable string
	Term Expr
	Paths [][]int
}

type Application struct {
	Body Expr
	Argument Expr
}

type Variable struct {
	Name string
}

type StringLit struct {
	Value string
}

type NumberLit struct {
	Value int
}

var Identity = &Abstraction{"x", &Variable{"x"}}

func (a *Abstraction) String() string {
	return fmt.Sprintf("($%v.%v)", a.Variable, a.Term)
}

func (a *Abstraction) Compare(expr Expr) bool {
	if v, ok := expr.(*Abstraction); ok {
		return a.Term.Compare(v.Term)
	}
	return false
}

func (p *PathedAbstraction) String() string {
	return fmt.Sprintf("($%v%v.%v)", p.Variable, p.Paths, p.Term)
}

func (p *PathedAbstraction) Compare(expr Expr) bool {
	if v, ok := expr.(*PathedAbstraction); ok {
		return p.Variable == v.Variable && p.Term.Compare(v.Term)
	}
	return false
}

func (a *Application) String() string {
	return fmt.Sprintf("%v %v", a.Body, a.Argument)
}

func (a *Application) Compare(expr Expr) bool {
	if v, ok := expr.(*Application); ok {
		return a.Body.Compare(v.Body) && a.Argument.Compare(v.Argument)
	}
	return false
}

func (v *Variable) String() string {
	return v.Name
}

func (v *Variable) Compare(expr Expr) bool {
	if o, ok := expr.(*Variable); ok {
		return v.Name == o.Name
	}
	return false
}

func (s *StringLit) String() string {
	return strconv.Quote(s.Value)
}

func (s *StringLit) Compare(expr Expr) bool {
	if v, ok := expr.(*StringLit); ok {
		return s.Value == v.Value
	}
	return false
}

func (n *NumberLit) String() string {
	return strconv.Itoa(n.Value)
}

func (n *NumberLit) Compare(expr Expr) bool {
	if v, ok := expr.(*NumberLit); ok {
		return n.Value == v.Value
	}
	return false
}
