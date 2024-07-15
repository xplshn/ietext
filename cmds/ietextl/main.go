package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"../pkg/common"
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
		number, text, err := lib.ParseLine(line)
		if err != nil {
			fmt.Println(err)
			continue
		}
		lines = append(lines, Line{Number: number, Text: text})
	}

	sort.Slice(lines, func(i, j int) bool {
		return lines[i].Number < lines[j].Number
	})

	for _, line := range lines {
		fmt.Printf("%02d %s\n", line.Number, line.Text)
	}
}
