package analysis

type Expr interface {
	isExpr()
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

func (a Abstraction) isExpr() {}
func (a Application) isExpr() {}
func (v Variable) isExpr() {}
func (s StringLit) isExpr() {}
