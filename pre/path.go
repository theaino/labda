package pre

import "labda/eval"

func CompilePaths(expr eval.Expr) eval.Expr {
	switch e := expr.(type) {
	case *eval.Application:
		return &eval.Application{Body: CompilePaths(e.Body), Argument: CompilePaths(e.Argument)}
	case *eval.Abstraction:
		return &eval.PathedAbstraction{Variable: e.Variable, Term: CompilePaths(e.Term), Paths: GetPaths(e.Term, e.Variable)}
	case *eval.PathedAbstraction:
		return &eval.PathedAbstraction{Variable: e.Variable, Term: CompilePaths(e.Term), Paths: e.Paths}
	default:
		return e
	}
}

func GetPaths(expr eval.Expr, name string) [][]int {
	paths := make([][]int, 0)
	switch e := expr.(type) {
	case *eval.Application:
		leftPaths := GetPaths(e.Body, name)
		rightPaths := GetPaths(e.Argument, name)
		for _, leftPath := range leftPaths {
			paths = append(paths, append([]int{0}, leftPath...))
		}
		for _, rightPath := range rightPaths {
			paths = append(paths, append([]int{1}, rightPath...))
		}
	case *eval.Abstraction:
		if e.Variable != name {
			paths = append(paths, GetPaths(e.Term, name)...)
		}
	case *eval.PathedAbstraction:
		paths = append(paths, GetPaths(e.Term, name)...)
	case *eval.Variable:
		if e.Name == name {
			paths = append(paths, make([]int, 0))
		}
	}
	return paths
}
