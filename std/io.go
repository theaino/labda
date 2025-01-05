package std

import (
	"fmt"
	"labda/eval"
)

var PutStrer = func(s eval.StringLit) Builtin {
	return Builtin(func(e eval.Expr) eval.Expr {
		fmt.Printf("%s", s.Value)
		return e.Apply(eval.Identity)
	})
}

var PutStr = Builtin(func(e eval.Expr) eval.Expr {
	switch v := e.(type) {
	case eval.StringLit:
		return PutStrer(v)
	default:
		panic("Can only print string")
	}
})


var GetLine = Builtin(func(e eval.Expr) eval.Expr {
	var line string
	fmt.Scan(&line)
	return e.Apply(eval.StringLit{line})
})

func IOPrepare(expr eval.Expr) eval.Expr {
	return eval.Substitute(eval.Substitute(expr, "PutStr", PutStr), "GetLine", GetLine)
}
