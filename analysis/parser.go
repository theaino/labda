package analysis

import "fmt"

func Parse(tokens []Token) Expr {
	expr, _ := parseParen(tokens)
	return expr
}

func parseParen(tokens []Token) (Expr, int) {
	var currentExpr Expr
	for idx := 0; idx < len(tokens); idx++ {
		token := tokens[idx]
		var expr Expr
		var offset int
		switch v := token.(type) {
		case Single:
			switch v {
			case LParen:
				expr, offset = parseParen(tokens[idx+1:])
			case RParen:
				return currentExpr, idx + 1
			case Lambda:
				expr, offset = parseLambda(tokens[idx+1:])
			}
			idx += offset
		case Word:
			expr = Variable{string(v)}
		case String:
			expr = StringLit{string(v)}
		}
		if currentExpr == nil {
			currentExpr = expr
		} else {
			currentExpr = Application{currentExpr, expr}
		}
	}
	return currentExpr, len(tokens)
}

func parseLambda(tokens []Token) (Expr, int) {
	var variable string
	switch v := tokens[0].(type) {
	case Word:
		variable = string(v)
	default:
		panic(fmt.Sprintf("Expected Word, got %v", v))
	}
	expr, offset := parseParen(tokens[1:])
	return Abstraction{variable, expr}, offset + 1
}
