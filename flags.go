// go run flags.go --message TechTalks
package main 

import (
	"flag"
    "fmt"
)

func main() {
    message := flag.String("message", "", "Text to show")
    flag.Parse()
    fmt.Printf("Hello %s!\n", *message)
}