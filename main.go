package main

import (
	// "fmt"
	"labda/analysis"
	"labda/std"
)

func main() {
	source := `
PutStr "Name: " ($_
GetLine ($s
PutStr "Hello, " ($_
PutStr s ($_
PutStr "!\n" ($x x
	`
	tokens := analysis.Lex(source)
	expr := std.Prepare(analysis.Parse(tokens))
	expr = expr.Reduce()
	// fmt.Printf("%v\n", expr)
}
