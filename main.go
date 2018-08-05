package main

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

var (
	file string
)

func init() {
	flag.StringVarP(&file, "file", "f", "", "Log files to tail")
}

func main() {
	fmt.Println("Test")
}
