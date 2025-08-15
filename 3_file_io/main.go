package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Problem 3 — File I/O (Reading & Writing)

Goal:
  Read a text file line-by-line, uppercase each line, and write to output.txt in this folder.

Why:
  Teaches file handling, buffered I/O, error checking, and working with strings.

Run:
  echo -e "hello\nworld" > input.txt
  go run ./3_file_io < input.txt
  cat 3_file_io/output.txt

TODOs:
  1) Read from stdin line-by-line (bufio.Scanner).
  2) Uppercase each line.
  3) Create/overwrite output.txt and write the transformed lines.
  4) Handle errors gracefully (print to stderr and exit non-zero if needed).
*/
func main() {
	// TODO: implement
	_ = bufio.NewScanner
	_ = os.Create
	_ = strings.ToUpper
	fmt.Println("TODO")
}
