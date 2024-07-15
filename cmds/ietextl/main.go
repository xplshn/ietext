package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/xplshn/ietext/pkg/common"
)

type Line struct {
	Number int
	Text   string
}

func main() {
	var lines []Line
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) >= 2 && line[:2] == "00" { // Special case for line continuation
			lines[len(lines)-1].Text += "\n" + line[2:]
		} else if len(line) >= 2 && strings.HasPrefix(line, "0") {
			number, err := strconv.Atoi(line[:2])
			if err != nil {
				fmt.Println("Error parsing line number:", err)
				continue
			}
			text := strings.TrimSpace(line[2:])
			lines = append(lines, Line{Number: number, Text: text})
		} else {
			fmt.Println("Skipping invalid line:", line)
		}
	}

	sort.Slice(lines, func(i, j int) bool {
		return lines[i].Number < lines[j].Number
	})

	for _, line := range lines {
		mLine := fmt.Sprintf("%02d %s\n", line.Number, line.Text)
		if common.IsValidLine(mLine) {
			fmt.Print(mLine)
		}
	}
}
