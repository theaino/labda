package main

import (
	"context"
	"labda/analysis"
	"labda/eval"
	"labda/pre"
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

func Load(data string) eval.Expr {
	tokens := analysis.Lex(string(data))
	return analysis.Parse(tokens)
}

func Prepare(expr eval.Expr) eval.Expr {
	return DefaultStd().Prepare(expr)
}

func Run(expr eval.Expr) {
	expr = expr.Reduce()
}

func main() {
	cmd := &cli.Command{
		Name: "labda",
		Commands: []*cli.Command{
			{
				Name: "compile",
				Usage: "compile a labda program",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "output",
						Aliases: []string{"o"},
						Value: "",
						Usage: "the output file",
					},
				},
				Action: func(ctx context.Context, c *cli.Command) error {
					path := c.Args().Get(0)
					data, err := os.ReadFile(path)
					if err != nil {
						return err
					}
				
					expr := Prepare(Load(string(data)))
					// expr = pre.CompilePaths(expr)
					outputData, err := pre.Encode(expr)
					if err != nil {
						return err
					}

					outputPath := path + ".lbc"
					if arg := c.String("output"); arg != "" {
						outputPath = arg
					}

					return os.WriteFile(outputPath, outputData, 0644)
				},
			},
			{
				Name: "run",
				Usage: "run a labda program",
				Action: func(ctx context.Context, c *cli.Command) error {
					path := c.Args().Get(0)
					data, err := os.ReadFile(path)
					if err != nil {
						return err
					}

					expr, err := pre.Decode(data)
					if err != nil {
						return err
					}

					Run(Prepare(expr))

					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
