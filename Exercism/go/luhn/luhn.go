package luhn

import (
	"strconv"
	"strings"
)

//Valid check Luhn algorithm for given input
func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")
	if len(input) < 2 {
		return false
	}

	odd := len(input)%2 == 0 
	sum := 0
	for _, r := range input {
		n, err := strconv.Atoi(string(r))
		if err != nil {
			return false
		}

		if odd {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
		odd = !odd

	}
	return sum%10 == 0
}
