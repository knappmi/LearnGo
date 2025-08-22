package main

import (
	"flag"
	"fmt"
)

/*
Problem 6 — CLI Flags

Goal:

	Parse a --name flag and print "Hello, <name>!". Default to "World".

Why:

	Teaches the standard `flag` package and CLI ergonomics.

Run:

	go run ./6_cli_flags --name Mike
	go run ./6_cli_flags
*/
func main() {
	name := flag.String("name", "World", "What's your name?")
	flag.Parse()
	fmt.Printf("Hello, %s!", *name)
}
