package std

import (
	"io"
	"labda/eval"
)

type Options struct {
	Writer io.Writer
	Reader io.Reader
}

type Builtin func(eval.Expr) eval.Expr

func (b Builtin) Reduce() eval.Expr {
	return b
}

func (b Builtin) Apply(argument eval.Expr) eval.Expr {
	return b(argument)
}

func (o Options) Prepare(expr eval.Expr) eval.Expr {
	return IOPrepare(expr, o)
}
