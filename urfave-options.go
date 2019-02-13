//go run urfave-options.go --help
//go run urfave-options.go --version
package main

import (
	"os"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Version = "1.1"
	app.Name = "hello-cli"
	app.Usage = "A CLI tool for Tech Talks"
	app.Run(os.Args)

}