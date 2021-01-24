package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func info() {
	app.Name = "Basic Calculation"
	app.Usage = "Arthematic Operations"
}

func commands(args []string) {
	if len(args) < 4 {
		return
	}
	if args[2] == "" || args[3] == "" {
		fmt.Println("Pass arguments to operate on")
		return

	}
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"+"},
				Usage:   "Addition",
				Action: func(c *cli.Context) error {
					result := 0
					for i := range os.Args {
						if i > 1 {
							res, err := strconv.Atoi(os.Args[i])
							if err != nil {
								return err
							}
							result += res
						}
					}
					fmt.Println("Result::", result)
					return nil
				},
			},
			{
				Name:    "sub",
				Aliases: []string{"-"},
				Usage:   "Substraction",
				Action: func(c *cli.Context) error {
					result := 0
					res1, err := strconv.Atoi(os.Args[2])
					res2, err := strconv.Atoi(os.Args[3])
					if err != nil {
						return err
					}
					result = res1 - res2
					fmt.Println("Result::", result)
					return nil
				},
			},
			{
				Name:    "mul",
				Aliases: []string{"m"},
				Usage:   "Multiplication",
				Action: func(c *cli.Context) error {
					result := 1
					for i := range os.Args {
						if i > 1 {
							res, err := strconv.Atoi(os.Args[i])
							if err != nil {
								return err
							}
							result *= res
						}
					}
					fmt.Println("Result::", result)
					return nil
				},
			},
			{
				Name:    "div",
				Aliases: []string{"/"},
				Usage:   "Division",
				Action: func(c *cli.Context) error {
					result := 0
					if os.Args[2] > os.Args[3] {
						res1, err := strconv.Atoi(os.Args[2])
						res2, err := strconv.Atoi(os.Args[3])
						if err != nil {
							return err
						}
						result = res1 / res2
						fmt.Println("Result::", result)
					} else {
						fmt.Println("first element should be greater than first element")
					}

					return nil
				},
			},
			{
				Name:    "mod",
				Aliases: []string{"%"},
				Usage:   "Moderator",
				Action: func(c *cli.Context) error {
					result := 0
					if os.Args[2] > os.Args[3] {
						res1, err := strconv.Atoi(os.Args[2])
						res2, err := strconv.Atoi(os.Args[3])
						if err != nil {
							return err
						}
						result = res1 % res2
						fmt.Println("Result::", result)
					} else {
						fmt.Println("first element should be greater than first element")
					}

					return nil
				},
			},
		},
	}
	sort.Sort(cli.CommandsByName(app.Commands))
}

func main() {
	info()
	commands(os.Args)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
