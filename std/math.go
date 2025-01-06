package std

import "labda/eval"

func BinaryOp(op func(int, int) int) BuiltinExpr {
	return Builtin(func(aExpr eval.Expr) eval.Expr {
		switch a := aExpr.Reduce().(type) {
		case eval.NumberLit:
			return Builtin(func(bExpr eval.Expr) eval.Expr {
				switch b := bExpr.Reduce().(type) {
				case eval.NumberLit:
					return eval.NumberLit{Value: op(a.Value, b.Value)}
				default:
					panic("Operation only supports numbers")
				}
			})
		default:
			panic("Operation only supports numbers")
		}
	})
}

func init() {
	Preparers = append(Preparers, func(options Options) {
		BuiltinMap["+"] = BinaryOp(func(a, b int) int { return a + b })
		BuiltinMap["-"] = BinaryOp(func(a, b int) int { return a - b })
		BuiltinMap["*"] = BinaryOp(func(a, b int) int { return a * b })
		BuiltinMap["/"] = BinaryOp(func(a, b int) int { return a / b })
	})
}
