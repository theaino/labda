package eval

import (
	"fmt"
	"strconv"
)

type Expr interface {
	Reduce() Expr
	Apply(Expr) Expr
	String() string
}

type Abstraction struct {
	Variable string
	Term Expr
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

var Identity = Abstraction{"x", Variable{"x"}}

func (a Abstraction) String() string {
	return fmt.Sprintf("($%v.%v)", a.Variable, a.Term)
}

func (a Application) String() string {
	return fmt.Sprintf("%v %v", a.Body, a.Argument)
}

func (v Variable) String() string {
	return v.Name
}

func (s StringLit) String() string {
	return strconv.Quote(s.Value)
}

func (n NumberLit) String() string {
	return strconv.Itoa(n.Value)
}
