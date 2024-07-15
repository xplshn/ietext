package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lineNumber := 10 // Starting line number

	// Read from standard input
	for scanner.Scan() {
		line := scanner.Text()

		// Prepend the line number to the line
		fmt.Printf("%02d %s\n", lineNumber, line)

		// Increment the line number by 10 for the next line
		lineNumber += 10
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
