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
			fmt.Println(err)
			continue
		}
		buffer.WriteString(text + "\n")
	}
	fmt.Print(buffer.String())
}
