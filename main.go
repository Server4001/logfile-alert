package main

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

var (
	file logFilesFlag
)

func init() {
	flag.VarP(&file, "file", "f", "Log files to tail.")
}

func main() {
	flag.Parse()

	fmt.Println("Test")
	fmt.Println(file)
}
