//go run urfave.go
package main

import (
	"os"
	"fmt"
	"github.com/urfave/cli"
)

func main() {
	err := cli.NewApp().Run(os.Args)
	if err != nil {
    	fmt.Println("Error")
  	}
}