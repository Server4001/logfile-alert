package main

import "fmt"

type logFilesFlag []string

func (i *logFilesFlag) String() string {
	return fmt.Sprintf("%v", *i)
}

func (i *logFilesFlag) Set(value string) error {
	*i = append(*i, value)

	return nil
}

// TODO : Figure out what this should do.
func (i *logFilesFlag) Type() string {
	return "log files"
}
