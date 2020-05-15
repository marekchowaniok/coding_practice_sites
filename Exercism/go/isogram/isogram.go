package isogram

import (
	"strings"
)

func IsIsogram(input string) bool {
	input = strings.ReplaceAll(input, "-", "")
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ToUpper(input)

	for _, v := range input {
		if strings.Count(input, string(v)) > 1 {
			return false
		}
	}
	return true
}
