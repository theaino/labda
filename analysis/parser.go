package analysis

import (
	"labda/eval"
	"strconv"
)

func Parse(tokens []Token) eval.Expr {
	expr, _ := parseBlock(tokens, nil)
	return expr
}

func parseBlock(tokens []Token, stopToken Token) (oldExpr eval.Expr, end int) {
	oldExpr = eval.Identity

	for end = 0; end < len(tokens); end++ {
		var expr eval.Expr
		switch token := tokens[end].(type) {
		case Word:
			expr = eval.Variable{Name: string(token)}
		case String:
			expr = eval.StringLit{Value: string(token)}
		case Number:
			num, err := strconv.Atoi(string(token))
			if err != nil {
				panic(err)
			}
			expr = eval.NumberLit{Value: num}
		case Single:
			switch token {
			case LParen:
				var offset int
				expr, offset = parseBlock(tokens[end+1:], RParen)
				end += offset
			case RParen:
				if stopToken == RParen {
					end++
				}
				return
			case Bar:
				var offset int
				expr, offset = parseBlock(tokens[end+1:], RParen)
				end += offset
				expr = eval.Abstraction{Variable: "", Term: expr}
			case Lambda:
				var offset int
				expr, offset = parseLambda(tokens[end+1:])
				end += offset
			case Dot:
				if stopToken == Dot {
					end++
				}
				return
			default:
				panic("Not implemented!")
			}
		default:
			panic("Not implemented!")
		}
		if oldExpr == eval.Identity {
			oldExpr = expr
		} else {
			oldExpr = eval.Application{Body: oldExpr, Argument: expr}
		}
	}

	return
}

func parseLambda(tokens []Token) (eval.Expr, int) {
	var name string
	switch token := tokens[0].(type) {
	case Word:
		name = string(token)
	default:
		panic("Expected variable name")
	}

	switch tokens[1] {
	case Dot:
		term, offset := parseBlock(tokens[2:], nil)
		return eval.Abstraction{Variable: name, Term: term}, offset + 2
	case Equal:
		value, valueOffset := parseBlock(tokens[2:], Dot)
		offset := valueOffset + 2
		body, bodyOffset := parseBlock(tokens[offset:], nil)
		offset += bodyOffset
		// return eval.Substitute(body, name, value), offset
		return eval.Application{Body: eval.Abstraction{Variable: name, Term: body}, Argument: value}, offset
	default:
		panic("Expected Dot or Equal")
	}
}
