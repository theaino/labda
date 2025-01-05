package std

import (
	"fmt"
	"io"
	"labda/eval"
)

func PutStrer(s eval.StringLit, w io.Writer) Builtin {
	return Builtin(func(e eval.Expr) eval.Expr {
		fmt.Printf("%s", s.Value)
		return e.Apply(eval.Identity)
	})
}

func PutStr(w io.Writer) Builtin {
	return Builtin(func(e eval.Expr) eval.Expr {
		switch v := e.(type) {
		case eval.StringLit:
			return PutStrer(v, w)
		default:
			panic("Can only print string")
		}
	})
}

func GetLine(r io.Reader) Builtin {
	return Builtin(func(e eval.Expr) eval.Expr {
		var line string
		fmt.Fscan(r, &line)
		return e.Apply(eval.StringLit{line})
	})
}

func Insert(expr eval.Expr, funcs map[string]eval.Expr) eval.Expr {
	for key, value := range funcs {
		expr = eval.Substitute(expr, key, value)
	}
	return expr
}

func IOPrepare(expr eval.Expr, options Options) eval.Expr {
	return Insert(expr, map[string]eval.Expr{
		"PutStr": PutStr(options.Writer),
		"GetLine": GetLine(options.Reader),
	})
}
