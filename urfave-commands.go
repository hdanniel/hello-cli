//go run urfave-commands.go --help
//go run urfave-commands.go --version
package main

import (
	"os"
	"fmt"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Version = "1.1"
	app.Name = "hello-cli"
	app.Usage = "A CLI tool for Tech Talks"

	app.Commands = []cli.Command{
		{
			Name:  "message",
			Usage: "write a message",
			Action: func(c *cli.Context) error {
				fmt.Println("hello", c.Args().First())
				return nil
			},
		},
	}
	app.Run(os.Args)



}