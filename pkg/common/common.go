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

	number, text := ExtractLineNumberAndText(line)
	return number, text, nil
}

func IsValidLine(line string) bool {
	// Find the position of the first non-digit character
	for i, ch := range line {
		if !IsDigit(ch) {
			// Ensure there's at least one digit at the start of the line
			return i > 0
		}
	}
	return false
}

func ExtractLineNumberAndText(line string) (int, string) {
	// Find the position of the first non-digit character
	for i, ch := range line {
		if !IsDigit(ch) {
			numberStr := line[:i]
			number, _ := strconv.Atoi(numberStr)
			text := strings.TrimSpace(line[i:])
			return number, text
		}
	}
	return 0, line // This case should never be hit if IsValidLine works correctly
}

func IsDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}
