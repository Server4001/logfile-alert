package main

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

var (
	file  logFilesFlag
	regex string
)

func init() {
	flag.VarP(&file, "file", "f", "Log files to tail.")
	flag.StringVarP(&regex, "regex", "r", "", "Regex to look for in log files.")
}

func main() {
	flag.Parse()

	fmt.Println("Test")
	fmt.Println(file)
	fmt.Println(regex)
}
