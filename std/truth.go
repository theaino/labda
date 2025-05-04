package std

import (
	"labda/eval"
)

var True = Builtin(func(e eval.Expr) eval.Expr {
	return Builtin(func(_ eval.Expr) eval.Expr {
		return e
	})
})

var False = Builtin(func(_ eval.Expr) eval.Expr {
	return Builtin(func(e eval.Expr) eval.Expr {
		return e
	})
})

var Eq = Builtin(func(a eval.Expr) eval.Expr {
	return Builtin(func(b eval.Expr) eval.Expr {
		return truth(a.Reduce().Compare(b.Reduce()))
	})
})

func truth(b bool) eval.Expr {
	if b {
		return BuiltinMap["True"]
	} else {
		return BuiltinMap["False"]
	}
}

func init() {
	Preparers = append(Preparers, func(options Options) {
		BuiltinMap["True"] = True
		BuiltinMap["False"] = False
		BuiltinMap["Eq"] = Eq
	})
}
