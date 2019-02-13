package main

import (
    "fmt"
    "github.com/ryanuber/columnize"
)

func main() {
    output := []string{
        "Job | Status | Last Build",
        "build-hello | running | #5",
        "backup-db | stopped | #4",
    }
    result := columnize.SimpleFormat(output)
    fmt.Println(result)
}