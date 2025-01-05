package std

import (
	"labda/eval"
)

type Builtin func(eval.Expr) eval.Expr

func (b Builtin) Reduce() eval.Expr {
	return b
}

func (b Builtin) Apply(argument eval.Expr) eval.Expr {
	return b(argument)
}

func Prepare(expr eval.Expr) eval.Expr {
	return IOPrepare(expr)
}
