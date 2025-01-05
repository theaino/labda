package analysis

import (
	"fmt"
	"labda/eval"
)

func Parse(tokens []Token) eval.Expr {
	expr, _ := parseParen(tokens)
	return expr
}

func parseParen(tokens []Token) (eval.Expr, int) {
	var currentExpr eval.Expr
	for idx := 0; idx < len(tokens); idx++ {
		token := tokens[idx]
		var expr eval.Expr
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
			expr = eval.Variable{string(v)}
		case String:
			expr = eval.StringLit{string(v)}
		}
		if currentExpr == nil {
			currentExpr = expr
		} else {
			currentExpr = eval.Application{currentExpr, expr}
		}
	}
	return currentExpr, len(tokens)
}

func parseLambda(tokens []Token) (eval.Expr, int) {
	var variable string
	switch v := tokens[0].(type) {
	case Word:
		variable = string(v)
	default:
		panic(fmt.Sprintf("Expected Word, got %v", v))
	}
	expr, offset := parseParen(tokens[1:])
	return eval.Abstraction{variable, expr}, offset + 1
}
