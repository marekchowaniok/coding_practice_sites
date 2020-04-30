package scrabble

import "strings"

// Score function for scramble
func Score(word string) int {

	result := 0
	for i := 0; i < len(word); i++ {
		result += getScore(strings.ToUpper(string(word[i])))
	}

	return result
}

func getScore(character string) int {
	data := map[int]string{
		1:  "AEIOULNRST",
		2:  "DG",
		3:  "BCMP",
		4:  "FHVWY",
		5:  "K",
		8:  "JX",
		10: "QZ",
	}

	result := 0
	for key, value := range data {
		if strings.Contains(value, character) {
			result = key
		}
	}
	return result
}
