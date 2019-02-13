//go run repl.go
package main

import (
	"os"
	"fmt"
	"github.com/urfave/cli"
	"os/signal"
	"syscall"
	"bufio"
	"strings"
)

func Repl(c *cli.Context) {

	input := ""
	for input != "." {
		in := bufio.NewReader(os.Stdin)
		fmt.Print("Hello> ")
		input, _ := in.ReadString('\n')
		inputArgs := strings.Fields(input)
		if len(inputArgs) == 0 {
			//fmt.Println(color("Command not recognized. Have you tried 'help'?", "response"))
			//fmt.Println(color("", "response"))
		} else {
			command := inputArgs[0]
			switch command {
			//case "clear", "cls":
			case "message":
				fmt.Println("hello my friend!")
			case "exit":
				fmt.Println("Bye!")
				os.Exit(3)
			default:
				fmt.Println("Command not recognized. Have you tried 'help'?")
			}
		}
	}
}


func main() {

	app := cli.NewApp()
	app.Version = "1.1"
	app.Name = "hello-cli"
	app.Usage = "A CLI tool for Tech Talks"

	app.Action = func(c *cli.Context) error {
		//Start the REPL if not argument given

		cs := make(chan os.Signal, 2)
		signal.Notify(cs, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-cs
			fmt.Println("Bye!")
			os.Exit(0)
		}()
		if c.NArg() == 0 {
			Repl(c)
		}
		return nil
	}

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