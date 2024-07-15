package common

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseLine extracts the line number and text from a given line.
func ParseLine(line string) (int, string, error) {
	parts := strings.SplitN(line, " ", 2)
	if len(parts) < 2 {
		return 0, "", fmt.Errorf("invalid line format")
	}

	number, err := strconv.Atoi(parts[0])
	if err != nil || number < 0 {
		return 0, "", fmt.Errorf("invalid line number")
	}

	return number, strings.TrimSpace(parts[1]), nil
}
