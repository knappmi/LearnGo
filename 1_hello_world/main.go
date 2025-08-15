package main

import (
	"fmt"
	"os"
)

/*
Problem 1 — Hello, World (Syntax & Basics)

Goal:

	Print "Hello, <name>!" to stdout. If a name is provided as the first CLI arg, use it;
	otherwise default to "World".

Why:

	Teaches package layout, main entrypoint, stdlib fmt/os, and basic branching.

Run:

	go run ./1_hello_world Mike
	# -> Hello, Mike!

TODOs:
 1. Read optional name from os.Args.
 2. Default to "World" if no name provided.
 3. Print the greeting.
*/
func main() {
	var nameArg string
	if len(os.Args) == 1 {
		nameArg = "world"
	} else {
		nameArg = os.Args[1]
	}
	fmt.Printf("Hello, %s!", nameArg)
}
