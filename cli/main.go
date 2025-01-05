package main

import (
	"context"
	"labda/analysis"
	"labda/eval"
	"labda/std"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)


func DefaultStd() std.Options {
	return std.Options{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}
}

func Pre(data string) eval.Expr {
	tokens := analysis.Lex(string(data))
	return DefaultStd().Prepare(analysis.Parse(tokens))
}

func Run(expr eval.Expr) {
	expr.Reduce()
}

func main() {
	cmd := &cli.Command{
		Name: "labda",
		Usage: "run a labda program",
		Action: func(ctx context.Context, c *cli.Command) error {
			data, err := os.ReadFile(c.Args().Get(0))
			if err != nil {
				return err
			}
			Run(Pre(string(data)))

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
