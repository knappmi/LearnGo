package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 1) Create/overwrite output.txt
	out, err := os.Create("output.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "create output.txt:", err)
		os.Exit(1)
	}
	// Ensure the file is closed even if we return early
	defer func() {
		if cerr := out.Close(); cerr != nil {
			fmt.Fprintln(os.Stderr, "close output.txt:", cerr)
			os.Exit(1)
		}
	}()

	// 2) Buffer writes for performance
	writer := bufio.NewWriter(out)
	defer func() {
		if ferr := writer.Flush(); ferr != nil {
			fmt.Fprintln(os.Stderr, "flush output.txt:", ferr)
			os.Exit(1)
		}
	}()

	// 3) Read stdin line-by-line
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		upper := strings.ToUpper(line)

		// 4) Write the transformed line + newline
		if _, err := writer.WriteString(upper + "\n"); err != nil {
			fmt.Fprintln(os.Stderr, "write to output.txt:", err)
			os.Exit(1)
		}
	}

	// 5) Detect scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "read stdin:", err)
		os.Exit(1)
	}
}
