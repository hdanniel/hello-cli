//go run  urfave-flags.go --language spanish  message TechTalks
package main

import (
	"os"
	"fmt"
	"github.com/urfave/cli"
)

var language string

func main() {

	app := cli.NewApp()
	app.Version = "1.1"
	app.Name = "hello-cli"
	app.Usage = "A CLI tool for Tech Talks"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "language",
			Value:       "english",
			Usage:       "Select a language",
			Destination: &language,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "message",
			Usage: "write a message",
			Action: func(c *cli.Context) error {
				if language == "spanish" {
					fmt.Println("hola", c.Args().First())
				} else {
					fmt.Println("hello", c.Args().First())
				}
				return nil
			},
		},
	}
	app.Run(os.Args)



}