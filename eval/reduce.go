package eval

func apply(body, argument Expr) Expr {
	return Application{Body: body, Argument: argument}
}

func (a Abstraction) Reduce() Expr {
	return a
}

func (a Abstraction) Apply(argument Expr) Expr {
	return Substitute(a.Term, a.Variable, argument)
}


func (a Application) Reduce() Expr {
	body := a.Body.Reduce()
	argument := a.Argument
	expr := body.Apply(argument)
	if expr == nil {
		expr = Application{body, argument}
	} else {
		expr = expr.Reduce()
	}
	return expr
}

func (a Application) Apply(argument Expr) Expr {
	return nil
}


func (v Variable) Reduce() Expr {
	return v
}

func (v Variable) Apply(argument Expr) Expr {
	return nil
}


func (s StringLit) Reduce() Expr {
	return s
}

func (s StringLit) Apply(argument Expr) Expr {
	return nil
}


func (n NumberLit) Reduce() Expr {
	return n
}

func (n NumberLit) Apply(argument Expr) Expr {
	return nil
}


func Substitute(expr Expr, variable string, value Expr) Expr {
	switch v := expr.(type) {
	case Application:
		return Application{Body: Substitute(v.Body, variable, value), Argument: Substitute(v.Argument, variable, value)}
	case Abstraction:
		if v.Variable == variable {
			return v
		} else {
			return Abstraction{Variable: v.Variable, Term: Substitute(v.Term, variable, value)}
		}
	case Variable:
		if v.Name == variable {
			return value
		} else {
			return v
		}
	default:
		return v
	}
}
