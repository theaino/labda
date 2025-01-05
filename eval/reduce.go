package eval

import (
	ana "labda/analysis"
)

func Collapse(expr ana.Expr) ana.Expr {
	switch v := expr.(type) {
	case ana.Application:
		body := Collapse(v.Body)
		argument := Collapse(v.Argument)
		switch b := body.(type) {
		case ana.Abstraction:
			return Substitute(b.Term, b.Variable, argument)
		default:
			return ana.Application{Body: body, Argument: argument}
		}
	default:
		return v
	}
}

func Substitute(expr ana.Expr, variable string, value ana.Expr) ana.Expr {
	switch v := expr.(type) {
	case ana.Application:
		return ana.Application{Body: Substitute(v.Body, variable, value), Argument: v.Argument}
	case ana.Abstraction:
		if v.Variable == variable {
			return v
		} else {
			return ana.Abstraction{Variable: v.Variable, Term: Substitute(v.Term, variable, value)}
		}
	case ana.Variable:
		if v.Name == variable {
			return value
		} else {
			return v
		}
	default:
		return v
	}
}
