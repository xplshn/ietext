// ietextil.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/xplshn/ietext/pkg/common"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lines := make(map[int]string)
	var lastNumber int
	var lastLine string
	hasLastLine := false

	for scanner.Scan() {
		line := scanner.Text()
		number, text, err := common.ParseLine(line)
		if err != nil {
			if hasLastLine {
				// Append this line to the last numbered line
				lines[lastNumber] = lastLine + " " + line
				lastLine = lines[lastNumber]
			}
			continue
		}
		// Store the line, overriding any existing line with the same number
		lines[number] = text
		lastNumber = number
		lastLine = text
		hasLastLine = true
	}

	// Collect and sort the line numbers
	keys := make([]int, 0, len(lines))
	for key := range lines {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	// Print the lines in the sorted order of their numbers
	for _, key := range keys {
		fmt.Println(lines[key])
	}
}
