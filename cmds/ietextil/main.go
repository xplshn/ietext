package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/xplshn/ietext/pkg/common"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var buffer strings.Builder
	for scanner.Scan() {
		line := scanner.Text()
		_, text, err := common.ParseLine(line)
		if err != nil {
			// If there's an error, just add the original line (without number) to the buffer.
			// This will ensure that lines without numbers are still included.
			buffer.WriteString(line + "\n")
			continue
		}
		buffer.WriteString(text + "\n")
	}
	fmt.Print(buffer.String())
}
