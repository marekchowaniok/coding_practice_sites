package scrabble

import (
	"strings"
)

var data = map[rune]string{
	1:  "AEIOULNRST",
	2:  "DG",
	3:  "BCMP",
	4:  "FHVWY",
	5:  "K",
	8:  "JX",
	10: "QZ",
}

// Score function for scramble
func Score(word string) int {
	result := 0

	for _, i := range strings.ToUpper(word) {
		// fmt.Printf("Size: %d\nType: %s\nCharacter: %c\n", unsafe.Sizeof(aa), reflect.TypeOf(aa), aa)

		for key, value := range data {
			if strings.Contains(value, string(rune(i))) {
				result += int(key)
			}
		}
	}

	return result
}
