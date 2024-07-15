package common

import (
	"errors"
	"strconv"
	"strings"
)

// ParseLine attempts to parse a line into a line number and text.
func ParseLine(line string) (int, string, error) {
	if !IsValidLine(line) {
		return 0, "", errors.New("invalid line format")
	}

	number, text := extractLineNumberAndText(line)
	return number, text, nil
}

func IsValidLine(line string) bool {
	if len(line) < 3 {
		return false
	}
	lineNumberStr := line[:2]
	if _, err := strconv.Atoi(lineNumberStr); err != nil {
		return false
	}
	return true
}

func extractLineNumberAndText(line string) (int, string) {
	numberStr := line[:2]
	number, _ := strconv.Atoi(numberStr)
	text := strings.TrimSpace(line[2:])
	return number, text
}
