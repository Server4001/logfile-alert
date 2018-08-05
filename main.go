package main

import (
	"fmt"

	flag "github.com/spf13/pflag"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return fmt.Sprintf("%v", *i)
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)

	return nil
}

func (i *arrayFlags) Type() string {
	return ""
}

var (
	file arrayFlags
)

func init() {
	flag.VarP(&file, "file", "f", "Log files to tail.")
}

func main() {
	flag.Parse()

	fmt.Println("Test")
	fmt.Println(file)
}
