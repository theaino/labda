package std

import (
	"io"
	"labda/eval"
)

type Options struct {
	Writer io.Writer
	Reader io.Reader
}

type BuiltinExpr struct {
	Name string
	Handler func(eval.Expr) eval.Expr
}

func Builtin(handler func(eval.Expr) eval.Expr) BuiltinExpr {
	return BuiltinExpr{
		Name: "Builtin",
		Handler: handler,
	}
}

var BuiltinMap = make(map[string]BuiltinExpr)
var Preparers = make([]func(Options), 0)

func (b BuiltinExpr) Reduce() eval.Expr {
	return b
}

func (b BuiltinExpr) Apply(argument eval.Expr) eval.Expr {
	handler := b.Handler
	if e, ok := BuiltinMap[b.Name]; ok {
		handler = e.Handler
	}
	return handler(argument)
}

func (o Options) Prepare(expr eval.Expr) eval.Expr {
	for _, preparer := range Preparers {
		preparer(o)
	}
	return Insert(expr)
}

func Insert(expr eval.Expr) eval.Expr {
	for key, value := range BuiltinMap {
		value.Name = key
		BuiltinMap[key] = value
		expr = eval.Substitute(expr, key, value)
	}
	return expr
}

func (b BuiltinExpr) String() string {
	return b.Name
}

func (b BuiltinExpr) Compare(expr eval.Expr) bool {
	if v, ok := expr.(*BuiltinExpr); ok {
		return b.Name == v.Name
	}
	return false
}
