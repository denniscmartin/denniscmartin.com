package main

import "fmt"

func panicInvalidLine(line string, filename string) {
	panic(fmt.Sprintf("invalid line: %s in file %s\n", line, filename))
}