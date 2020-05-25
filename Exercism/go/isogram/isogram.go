package isogram

import (
	"strings"
)

// IsIsogram function checks whether input is isogram or not
func IsIsogram(input string) bool {
	input = strings.ToUpper(input)

	characters := map[rune]bool{}
	for _, v := range input {
		if v == ' ' || v == '-' {
			continue
		}

		if characters[v] {
			return false
		}

		characters[v] = true
	}
	return true
}
