package std

import (
	"fmt"
	"io"
	"labda/eval"
)

func Printer(s string, w io.Writer) BuiltinExpr {
	expr := Builtin(func(e eval.Expr) eval.Expr {
		fmt.Printf("%s", s)
		return e.Apply(eval.Identity)
	})
	expr.Name = fmt.Sprintf("Printer{%v}", s)
	return expr
}

func Print(w io.Writer) BuiltinExpr {
	return Builtin(func(e eval.Expr) eval.Expr {
		switch v := e.Reduce().(type) {
		case *eval.StringLit:
			return Printer(v.Value, w)
		default:
			return Printer(v.String(), w)
		}
	})
}

func Input(r io.Reader) BuiltinExpr {
	return Builtin(func(e eval.Expr) eval.Expr {
		var line string
		fmt.Fscan(r, &line)
		return e.Apply(&eval.StringLit{Value: line})
	})
}

func init() {
	Preparers = append(Preparers, func(options Options) {
		BuiltinMap["Print"] = Print(options.Writer)
		BuiltinMap["Input"] = Input(options.Reader)
	})
}
