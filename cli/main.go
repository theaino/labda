package main

import (
	"context"
	"labda/analysis"
	"labda/std"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)


func main() {
	cmd := &cli.Command{
		Name: "labda",
		Usage: "run a labda program",
		Action: func(ctx context.Context, c *cli.Command) error {
			data, err := os.ReadFile(c.Args().Get(0))
			if err != nil {
				return err
			}
			tokens := analysis.Lex(string(data))
			expr := std.Prepare(analysis.Parse(tokens))
			expr = expr.Reduce()
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
