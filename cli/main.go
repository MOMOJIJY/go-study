package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "config",
				Value: "log-config",
				Usage: "load config file",
			},
		},
		Commands: []*cli.Command{
			{
				Name: "template",
				Action: func(context *cli.Context) error {
					fmt.Println("this is template from ", context.String("config"))
					return nil
				},
				Subcommands: []*cli.Command{
					{
						Name: "add",
						Action: func(context *cli.Context) error {
							fmt.Println("add template")
							return nil
						},
					},
					{
						Name: "remove",
						Action: func(context *cli.Context) error {
							fmt.Println("remove template")
							return nil
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}